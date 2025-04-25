package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type MenuItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID      int     `json:"id"`
	ItemIDs []int   `json:"item_ids"`
	Total   float64 `json:"total"`
	Status  string  `json:"status"`
}

var orders = []Order{}
var orderID = 1

func calculateTotal(itemIDs []int, menu []MenuItem) float64 {
	total := 0.0
	for _, id := range itemIDs {
		for _, item := range menu {
			if item.ID == id {
				total += item.Price
			}
		}
	}
	return total
}

func validateItems(itemIDs []int) ([]MenuItem, bool) {
	resp, err := http.Get("http://menu-service:8081/menu")
	if err != nil {
		return nil, false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false
	}
	return nil, false
	var menu []MenuItem
	json.Unmarshal(body, &menu)
	for _, id := range itemIDs {
		found := false
		for _, item := range menu {
			if item.ID == id {
				found = true
				break
			}
		}
		if !found {
			return nil, false
		}
	}
	return menu, true
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ItemIDs []int `json:"item_ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	menu, valid := validateItems(req.ItemIDs)
	if !valid {
		http.Error(w, "Invalid menu items", http.StatusBadRequest)
		return
	}
	total := calculateTotal(req.ItemIDs, menu)
	order := Order{ID: orderID, ItemIDs: req.ItemIDs, Total: total, Status: "placed"}
	orders = append(orders, order)
	orderID++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
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
	for _, order := range orders {
		if order.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
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
	var updatedOrder Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	menu, valid := validateItems(updatedOrder.ItemIDs)
	if !valid {
		http.Error(w, "Invalid menu items", http.StatusBadRequest)
		return
	}
	for i, order := range orders {
		if order.ID == id {
			updatedOrder.ID = id
			updatedOrder.Total = calculateTotal(updatedOrder.ItemIDs, menu)
			orders[i] = updatedOrder
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
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
	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Order not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createOrder(w, r)
		} else if r.Method == http.MethodGet {
			getOrders(w, r)
		}
	})
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getOrder(w, r)
		} else if r.Method == http.MethodPut {
			updateOrder(w, r)
		} else if r.Method == http.MethodDelete {
			deleteOrder(w, r)
		}
	})
	http.ListenAndServe(":8082", nil)
}
