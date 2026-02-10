package main

import (
	"context"
	"log"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

func seedArtists() error {

	artists := []struct {
		Name     string
		ImageURL string
	}{
		{
			Name:     "Daft Punk",
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/8/83/Daft_Punk_in_2013_2.jpg",
		},
		{
			Name:     "Coldplay",
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/c/cc/ColdplayWembley120925_%28cropped%29.jpg",
		},
		{
			Name:     "Radiohead",
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/a/a1/RadioheadO2211125_composite.jpg",
		},
		{
			Name:     "Adele",
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/6/68/Adele.jpg",
		},
		{
			Name:     "Imagine Dragons",
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/a/a2/Imagine_Dragons%2C_Roundhouse%2C_London_%2835390234536%29_%28cropped%29.jpg",
		},
		{
			Name:     "Gims",
			ImageURL: "https://upload.wikimedia.org/wikipedia/commons/6/67/NDLE2025Gims_1.jpg",
		},
	}

	ctx := context.Background()

	for _, artist := range artists {
		_, err := database.DB.Exec(
			ctx,
			`INSERT INTO artists (name, image_url)
			 VALUES ($1, $2)
			 ON CONFLICT DO NOTHING`,
			artist.Name,
			artist.ImageURL,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed artists completed ")
	return nil
}
