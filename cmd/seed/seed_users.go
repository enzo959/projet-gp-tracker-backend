package main

import (
	"context"
	"log"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func seedUsers() error {
	users := []struct {
		Email    string
		Password string
	}{
		{"test@mail.com", "password123"},
	}

	ctx := context.Background()

	for _, u := range users {
		// Hash du mot de passe
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		_, err = database.DB.Exec(
			ctx,
			`INSERT INTO users (email, password_hash) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING`,
			u.Email, passwordHash,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed users completed")
	return nil
}
