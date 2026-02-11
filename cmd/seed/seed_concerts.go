package main

import (
	"context"
	"log"
	"time"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

func seedConcerts() error {

	concerts := []struct {
		ArtistID     int
		ArtistName   string
		Date         string
		Location     string
		PriceCents   int
		TotalTickets int
		Detail       string
		ImageURL     string
	}{
		{1, "DaftPunk", "2025-06-10T20:00:00Z", "Paris", 3500, 500, "les concert est cool", "https://upload.wikimedia.org/wikipedia/commons/c/cd/Palais_Omnisports_de_Paris-Bercy_2007.jpg"},
		{1, "DaftPunk", "2025-07-02T19:30:00Z", "Lyon", 4200, 300, "les concert est cool", "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1d/Gojira_LDLC_Arena_2025_-_Flying_Whales.jpg/1920px-Gojira_LDLC_Arena_2025_-_Flying_Whales.jpg"},
		{2, "ColdPlayd", "2025-08-01T21:00:00Z", "Marseille", 3900, 800, "les concert est cool", "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3d/Views_of_Marseille_from_the_Cit%C3%A9_radieuse_4.jpg/1920px-Views_of_Marseille_from_the_Cit%C3%A9_radieuse_4.jpg"},
		{3, "Radiohead", "2025-09-15T20:30:00Z", "Bordeaux", 4500, 600, "les concert est cool", "https://upload.wikimedia.org/wikipedia/commons/3/31/Vue_du_Pont_Jacques_Chaban-Delmas.jpg"},
		{4, "Adele", "2025-10-05T20:00:00Z", "Nice", 3800, 400, "les concert est cool", "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d1/H%C3%B4pital_lenval_Nice.jpg/1920px-H%C3%B4pital_lenval_Nice.jpg"},
	}
	ctx := context.Background()

	for _, c := range concerts {
		// convertit la date string en time.Time
		date, err := time.Parse(time.RFC3339, c.Date)
		if err != nil {
			return err
		}

		_, err = database.DB.Exec(
			ctx,
			`INSERT INTO concerts (artist_id, artist_name, date, location, price_cents, total_tickets, detail, image_url)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT DO NOTHING`,
			c.ArtistID, c.ArtistName, date, c.Location, c.PriceCents, c.TotalTickets, c.Detail, c.ImageURL,
		)
		if err != nil {
			return err
		}
	}

	log.Println("Seed concerts completed")
	return nil
}
