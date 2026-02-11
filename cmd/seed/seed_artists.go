package main

import (
	"context"
	"log"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

func seedArtists() error {

	artists := []struct {
		Name       string
		Bio        string
		ImageURL   string
		MusiqueURL string
	}{
		{
			Name:       "Daft Punk",
			Bio:        "voici ma bio",
			ImageURL:   "https://upload.wikimedia.org/wikipedia/commons/8/83/Daft_Punk_in_2013_2.jpg",
			MusiqueURL: "",
		},
		{
			Name:       "Coldplay",
			Bio:        "ma bio",
			ImageURL:   "https://upload.wikimedia.org/wikipedia/commons/c/cc/ColdplayWembley120925_%28cropped%29.jpg",
			MusiqueURL: "",
		},
		{
			Name:       "Radiohead",
			Bio:        "ma bio",
			ImageURL:   "https://upload.wikimedia.org/wikipedia/commons/a/a1/RadioheadO2211125_composite.jpg",
			MusiqueURL: "",
		},
		{
			Name:       "Adele",
			Bio:        "ma bio",
			ImageURL:   "https://upload.wikimedia.org/wikipedia/commons/6/68/Adele.jpg",
			MusiqueURL: "",
		},
		{
			Name:       "Imagine Dragons",
			Bio:        "ma bio",
			ImageURL:   "https://upload.wikimedia.org/wikipedia/commons/a/a2/Imagine_Dragons%2C_Roundhouse%2C_London_%2835390234536%29_%28cropped%29.jpg",
			MusiqueURL: "",
		},
		{
			Name:       "Gims",
			Bio:        "voici ma bio",
			ImageURL:   "https://upload.wikimedia.org/wikipedia/commons/6/67/NDLE2025Gims_1.jpg",
			MusiqueURL: "",
		},
	}

	ctx := context.Background()

	for _, artist := range artists {
		_, err := database.DB.Exec(
			ctx,
			`INSERT INTO artists (name, bio, image_url, musique_url)
			 VALUES ($1, $2, $3, $4)
			 ON CONFLICT DO NOTHING`,
			artist.Name,
			artist.Bio,
			artist.ImageURL,
			artist.MusiqueURL,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed artists completed ")
	return nil
}
