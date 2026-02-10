package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
	"github.com/go-chi/chi/v5"
)

type CreateConcertInput struct {
	ArtistID     int       `json:"artist_id"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	PriceCents   int       `json:"price_cents"`
	TotalTickets int       `json:"total_tickets"`
	Detail       string    `json:"detail"`
	ImageURL     string    `json:"image_url"`
}

type Concert struct {
	ID           int       `json:"id"`
	ArtistID     int       `json:"artist_id"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	PriceCents   int       `json:"price_cents"`
	TotalTickets int       `json:"total_tickets"`
	Detail       string    `json:"detail"`
	ImageURL     string    `json:"image_url"`
}

type ArtistResponse struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Concerts []Concert `json:"concerts"`
}

func fetchConcerts(query string, args ...any) ([]Concert, error) {
	rows, err := database.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	concerts := []Concert{}

	for rows.Next() {
		var c Concert
		if err := rows.Scan(
			&c.ID,
			&c.ArtistID,
			&c.Date,
			&c.Location,
			&c.PriceCents,
			&c.TotalTickets,
			&c.Detail,
			&c.ImageURL,
		); err != nil {
			return nil, err
		}
		concerts = append(concerts, c)
	}

	return concerts, nil
}

func GetConcerts(w http.ResponseWriter, r *http.Request) {
	concerts, err := fetchConcerts(`
		SELECT id, artist_id, date, location, price_cents, total_tickets, detail, image_url
		FROM concerts
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(concerts)
}

func GetConcertsByArtist(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "id")

	concerts, err := fetchConcerts(`
		SELECT id, artist_id, date, location, price_cents, total_tickets, detail, image_url
		FROM concerts
		WHERE artist_id = $1
	`, artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(concerts)
}

func GetArtistByID(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "id")
	ctx := context.Background()

	var artist ArtistResponse
	err := database.DB.QueryRow(
		ctx,
		`SELECT id, name FROM artists WHERE id = $1`,
		artistID,
	).Scan(&artist.ID, &artist.Name)

	if err != nil {
		http.Error(w, "artist not found", http.StatusNotFound)
		return
	}

	concerts, err := fetchConcerts(`
		SELECT id, artist_id, date, location, price_cents, total_tickets, detail, image_url
		FROM concerts
		WHERE artist_id = $1
	`, artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	artist.Concerts = concerts

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}

func CreateConcert(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	role := r.Context().Value("role").(string)

	fmt.Println("User ID:", userID, "Role:", role)

	var input CreateConcertInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.ArtistID == 0 || input.Location == "" || input.TotalTickets <= 0 {
		http.Error(w, "missing or invalid fields", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO concerts (artist_id, date, location, price_cents, total_tickets, detail, image_url)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		input.ArtistID,
		input.Date,
		input.Location,
		input.PriceCents,
		input.TotalTickets,
		input.Detail,
		input.ImageURL,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "concert créé avec succès.",
	})
}

func UpdateConcert(w http.ResponseWriter, r *http.Request) {
	var input CreateConcertInput
	id := chi.URLParam(r, "id") // récupère l'ID du concert depuis l'URL

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(
		context.Background(),
		`UPDATE concerts
		 SET artist_id=$1, date=$2, location=$3, price_cents=$4, total_tickets=$5, detail=$6, image_url=$7
		 WHERE id=$6`,
		input.ArtistID,
		input.Date,
		input.Location,
		input.PriceCents,
		input.TotalTickets,
		input.Detail,
		input.ImageURL,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "concert mis à jour avec succès",
	})
}

func DeleteConcert(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // récupère l'ID du concert depuis l'URL

	_, err := database.DB.Exec(
		context.Background(),
		`DELETE FROM concerts WHERE id=$1`,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "concert supprimé avec succès",
	})
}
