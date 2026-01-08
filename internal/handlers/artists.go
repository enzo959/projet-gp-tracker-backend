package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/enzo959/projet-gp-tracker-backend/internal/database"
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
