package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	ID      int     `json:"id"`
	ItemIDs []int   `json:"item_ids"`
	Total   float64 `json:"total"`
	Status  string  `json:"status"`
}

type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

var payments = []Payment{}
var paymentID = 1

func validateOrder(orderID int) (float64, bool) {
	orderServiceURL := os.Getenv("ORDER_SERVICE_URL")
	if orderServiceURL == "" {
		orderServiceURL = "http://order-service:8082" // Fallback
	}
	resp, err := http.Get(orderServiceURL + "/orders")
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, false
	}
	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return 0, false
	}
	for _, order := range orders {
		if order.ID == orderID {
			return order.Total, true
		}
	}
	return 0, false
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrderID int `json:"order_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	total, valid := validateOrder(req.OrderID)
	if !valid {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	payment := Payment{ID: paymentID, OrderID: req.OrderID, Amount: total, Status: "completed"}
	payments = append(payments, payment)
	paymentID++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func getPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

func getPayment(w http.ResponseWriter, r *http.Request) {
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
	for _, payment := range payments {
		if payment.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(payment)
			return
		}
	}
	http.Error(w, "Payment not found", http.StatusNotFound)
}

func updatePayment(w http.ResponseWriter, r *http.Request) {
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
	var updatedPayment Payment
	if err := json.NewDecoder(r.Body).Decode(&updatedPayment); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i, payment := range payments {
		if payment.ID == id {
			updatedPayment.ID = id
			total, valid := validateOrder(updatedPayment.OrderID)
			if !valid {
				http.Error(w, "Invalid order ID", http.StatusBadRequest)
				return
			}
			updatedPayment.Amount = total
			payments[i] = updatedPayment
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedPayment)
			return
		}
	}
	http.Error(w, "Payment not found", http.StatusNotFound)
}

func deletePayment(w http.ResponseWriter, r *http.Request) {
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
	for i, payment := range payments {
		if payment.ID == id {
			payments = append(payments[:i], payments[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Payment not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createPayment(w, r)
		} else if r.Method == http.MethodGet {
			getPayments(w, r)
		}
	})
	http.HandleFunc("/payments/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getPayment(w, r)
		} else if r.Method == http.MethodPut {
			updatePayment(w, r)
		} else if r.Method == http.MethodDelete {
			deletePayment(w, r)
		}
	})
	http.ListenAndServe(":8084", nil)
}
