package main

import (
	"context"
	"log"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

func seedArtists() error {

	artists := []string{
		"Daft Punk",
		"Coldplay",
		"Radiohead",
		"Adele",
		"Imagine Dragons",
	}

	ctx := context.Background()

	for _, name := range artists {
		_, err := database.DB.Exec(
			ctx,
			`INSERT INTO artists (name) VALUES ($1) ON CONFLICT DO NOTHING`,
			name,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed artists completed ")
	return nil
}
