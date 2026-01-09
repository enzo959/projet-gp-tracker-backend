package main

import (
	"context"
	"log"
	"time"

	"github.com/enzo959/projet-gp-tracker-backend/internal/database"
)

func seedConcerts() error {

	concerts := []struct {
		ArtistID     int
		Date         string
		Location     string
		PriceCents   int
		TotalTickets int
	}{
		{1, "2025-06-10T20:00:00Z", "Paris", 3500, 500},
		{1, "2025-07-02T19:30:00Z", "Lyon", 4200, 300},
		{2, "2025-08-01T21:00:00Z", "Marseille", 3900, 800},
		{3, "2025-09-15T20:30:00Z", "Bordeaux", 4500, 600},
		{4, "2025-10-05T20:00:00Z", "Nice", 3800, 400},
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
			`INSERT INTO concerts (artist_id, date, location, price_cents, total_tickets)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT DO NOTHING`,
			c.ArtistID, date, c.Location, c.PriceCents, c.TotalTickets,
		)
		if err != nil {
			return err
		}
	}

	log.Println("Seed concerts completed")
	return nil
}
