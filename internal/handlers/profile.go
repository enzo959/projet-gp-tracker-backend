package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

type ProfileResponse struct {
	ID        int             `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Surname   string          `json:"surname"`
	Email     string          `json:"email"`
	Image     string          `json:"image"`
	Bio       string          `json:"bio"`
	Tickets   []TicketProfile `json:"tickets"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Surname   string `json:"surname"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
}

type TicketProfile struct {
	ConcertID int       `json:"concert_id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Location  string    `json:"location"`
	BoughtAt  time.Time `json:"bought_at"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	var profile ProfileResponse
	profile.Tickets = []TicketProfile{}

	// Récupérer les infos de l'utilisateur
	row := database.DB.QueryRow(context.Background(), `
        SELECT id,
			COALESCE(first_name, ''),
			COALESCE(last_name, ''),
		   	COALESCE(surname, ''),
		    email,
			COALESCE(image, ''),
			COALESCE(bio, '')
        FROM users WHERE id = $1
    `, userID)

	if err := row.Scan(
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Surname,
		&profile.Email,
		&profile.Image,
		&profile.Bio,
	); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// Récupérer les tickets
	rows, err := database.DB.Query(context.Background(), `
        SELECT c.id, COALESCE(c.location, 'Concert'), c.date, c.location, t.created_at
        FROM tickets t
        JOIN concerts c ON c.id = t.concert_id
        WHERE t.user_id = $1
    `, userID)
	if err != nil {
		http.Error(w, "error fetching tickets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t TicketProfile
		if err := rows.Scan(&t.ConcertID, &t.Name, &t.Date, &t.Location, &t.BoughtAt); err != nil {
			http.Error(w, "error scanning ticket", http.StatusInternalServerError)
			return
		}
		profile.Tickets = append(profile.Tickets, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	val := r.Context().Value("user_id")
	if val == nil {
		http.Error(w, "Utilisateur non authentifié", http.StatusUnauthorized)
		return
	}

	userID, ok := val.(int)
	if !ok {
		http.Error(w, "ID utilisateur invalide", http.StatusUnauthorized)
		return
	}

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(context.Background(), `
    UPDATE users
    SET first_name = $1,
        last_name = $2,
        surname = $3,
        bio = $4,
        image = $5,
        updated_at = NOW()
    WHERE id = $6
`, req.FirstName, req.LastName, req.Surname, req.Bio, req.Image, userID)

	if err != nil {
		fmt.Println("Erreur SQL lors de l'Update:", err)
		http.Error(w, "Erreur base de données", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
