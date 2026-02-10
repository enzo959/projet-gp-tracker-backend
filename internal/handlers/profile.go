package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

type ProfileResponse struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Image   string          `json:"image"`
	Tickets []TicketProfile `json:"tickets"`
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

	// Récupérer les infos de l'utilisateur
	row := database.DB.QueryRow(context.Background(), `
        SELECT id, name, email, image
        FROM users
        WHERE id = $1
    `, userID)

	if err := row.Scan(&profile.ID, &profile.Name, &profile.Email, &profile.Image); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// Récupérer les tickets
	rows, err := database.DB.Query(context.Background(), `
        SELECT c.id, c.name, c.date, c.location, t.created_at
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
