package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type MenuItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var menu = []MenuItem{
	{ID: 1, Name: "Burger", Price: 10.99},
	{ID: 2, Name: "Pizza", Price: 12.99},
	{ID: 3, Name: "Soda", Price: 2.99},
}
var nextID = 4

func createMenuItem(w http.ResponseWriter, r *http.Request) {
	var item MenuItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	item.ID = nextID
	nextID++
	menu = append(menu, item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func getMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

func getMenuItem(w http.ResponseWriter, r *http.Request) {
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
	for _, item := range menu {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func updateMenuItem(w http.ResponseWriter, r *http.Request) {
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
	var updatedItem MenuItem
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i, item := range menu {
		if item.ID == id {
			updatedItem.ID = id
			menu[i] = updatedItem
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteMenuItem(w http.ResponseWriter, r *http.Request) {
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
	for i, item := range menu {
		if item.ID == id {
			menu = append(menu[:i], menu[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createMenuItem(w, r)
		} else if r.Method == http.MethodGet {
			getMenu(w, r)
		}
	})
	http.HandleFunc("/menu/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getMenuItem(w, r)
		} else if r.Method == http.MethodPut {
			updateMenuItem(w, r)
		} else if r.Method == http.MethodDelete {
			deleteMenuItem(w, r)
		}
	})
	http.ListenAndServe(":8081", nil)
}
