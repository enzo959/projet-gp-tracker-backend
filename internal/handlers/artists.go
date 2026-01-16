package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
	"github.com/go-chi/chi/v5"
)

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistHandler struct {
	DB DBQuerier
}

/*
DBQuerier est une petite interface pour éviter
de dépendre directement de pgxpool partout
*/
type DBQuerier interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close()
}

type CreateArtistInput struct {
	Name string `json:"name"`
}

func NewArtistHandler(db DBQuerier) *ArtistHandler {
	return &ArtistHandler{DB: db}
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		context.Background(),
		"SELECT id, name FROM artists ORDER BY id",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var artists []Artist

	for rows.Next() {
		var a Artist
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		artists = append(artists, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	var input CreateArtistInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if input.Name == "" {
		http.Error(w, "artist name is required", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO artists (name) VALUES ($1) ON CONFLICT DO NOTHING`,
		input.Name,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "artist created successfully",
	})
}

func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID de l'artiste depuis l'URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Décode le JSON reçu dans le body
	var input struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Vérifie que le nom n'est pas vide
	if input.Name == "" {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}

	// Mise à jour dans la base de données
	_, err = database.DB.Exec(
		context.Background(),
		"UPDATE artists SET name=$1 WHERE id=$2",
		input.Name,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Artist updated successfully",
	})
}

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID de l'artiste depuis l'URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Supprime l'artiste dans la DB
	_, err = database.DB.Exec(
		context.Background(),
		"DELETE FROM artists WHERE id=$1",
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Artist deleted successfully",
	})
}
