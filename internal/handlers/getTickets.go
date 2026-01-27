package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
)

type TicketResponse struct {
	ID       int       `json:"id"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

func GetMyTickets(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	rows, err := database.DB.Query(
		context.Background(),
		`
        SELECT t.id, c.location, c.date
        FROM tickets t
        JOIN concerts c ON t.concert_id = c.id
        WHERE t.user_id=$1
        `,
		userID,
	)
	if err != nil {
		http.Error(w, "error fetching tickets", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tickets []TicketResponse
	for rows.Next() {
		var t TicketResponse
		rows.Scan(&t.ID, &t.Location, &t.Date)
		tickets = append(tickets, t)
	}

	json.NewEncoder(w).Encode(tickets)
}
