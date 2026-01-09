package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/enzo959/projet-gp-tracker-backend/internal/database"
)

type Concert struct {
	ID           int       `json:"id"`
	ArtistID     int       `json:"artist_id"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	PriceCents   int       `json:"price_cents"`
	TotalTickets int       `json:"total_tickets"`
}

func GetConcerts(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		context.Background(), `
		SELECT id, artist_id, date, location, price_cents, total_tickets
		FROM concerts
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		concerts = append(concerts, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(concerts)
}
