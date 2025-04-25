package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Reservation struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	Time         string `json:"time"`
	TableNumber  int    `json:"table_number"`
}

var reservations = []Reservation{}
var reservationID = 1

func createReservation(w http.ResponseWriter, r *http.Request) {
	var req Reservation
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	req.ID = reservationID
	reservationID++
	reservations = append(reservations, req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func getReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

func getReservation(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, reservation := range reservations {
		if reservation.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(reservation)
			return
		}
	}
	http.Error(w, "Reservation not found", http.StatusNotFound)
}

func updateReservation(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedReservation Reservation
	if err := json.NewDecoder(r.Body).Decode(&updatedReservation); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i, reservation := range reservations {
		if reservation.ID == id {
			updatedReservation.ID = id
			reservations[i] = updatedReservation
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedReservation)
			return
		}
	}
	http.Error(w, "Reservation not found", http.StatusNotFound)
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, reservation := range reservations {
		if reservation.ID == id {
			reservations = append(reservations[:i], reservations[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Reservation not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/reservations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createReservation(w, r)
		} else if r.Method == http.MethodGet {
			getReservations(w, r)
		}
	})
	http.HandleFunc("/reservations/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getReservation(w, r)
		} else if r.Method == http.MethodPut {
			updateReservation(w, r)
		} else if r.Method == http.MethodDelete {
			deleteReservation(w, r)
		}
	})
	http.ListenAndServe(":8085", nil)
}
