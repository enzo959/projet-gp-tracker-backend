package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
	"github.com/go-chi/chi/v5"
)

func BuyTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	concertID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var totalTickets int
	err := database.DB.QueryRow(
		context.Background(),
		"SELECT total_tickets FROM concerts WHERE id=$1",
		concertID,
	).Scan(&totalTickets)

	if err != nil {
		http.Error(w, "concert not found", http.StatusNotFound)
		return
	}

	var sold int
	_ = database.DB.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM tickets WHERE concert_id=$1",
		concertID,
	).Scan(&sold)

	if sold >= totalTickets {
		http.Error(w, "sold out", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		context.Background(),
		"INSERT INTO tickets (user_id, concert_id) VALUES ($1, $2)",
		userID,
		concertID,
	)

	if err != nil {
		http.Error(w, "cannot create ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
