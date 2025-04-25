package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type ProcessedOrder struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

var processedOrders = []ProcessedOrder{}
var nextKitchenID = 1

func createProcessedOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	processed := ProcessedOrder{ID: nextKitchenID, Status: req.Status}
	processedOrders = append(processedOrders, processed)
	nextKitchenID++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processed)
}

func getProcessedOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processedOrders)
}

func getProcessedOrder(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, order := range processedOrders {
		if order.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

func updateProcessedOrder(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedOrder ProcessedOrder
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i, order := range processedOrders {
		if order.ID == id {
			updatedOrder.ID = id
			processedOrders[i] = updatedOrder
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

func deleteProcessedOrder(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, order := range processedOrders {
		if order.ID == id {
			processedOrders = append(processedOrders[:i], processedOrders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/kitchen/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createProcessedOrder(w, r)
		} else if r.Method == http.MethodGet {
			getProcessedOrders(w, r)
		}
	})
	http.HandleFunc("/kitchen/orders/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getProcessedOrder(w, r)
		} else if r.Method == http.MethodPut {
			updateProcessedOrder(w, r)
		} else if r.Method == http.MethodDelete {
			deleteProcessedOrder(w, r)
		}
	})
	http.ListenAndServe(":8083", nil)
}
