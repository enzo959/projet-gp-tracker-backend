package main

import (
	"log"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

func main() {
	// Connexion à la DB
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	if err := seedArtists(); err != nil {
		log.Fatal(err)
	}

	if err := seedConcerts(); err != nil {
		log.Fatal(err)
	}

	if err := seedUsers(); err != nil {
		log.Fatal(err)
	}

}
