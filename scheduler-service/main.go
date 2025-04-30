package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TriggerResponse struct {
	Service string `json:"service"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Scheduler Service is healthy")
	})

	// Trigger endpoint to invoke APIs
	http.HandleFunc("/trigger", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get environment variables
		menuURL := os.Getenv("MENU_SERVICE_URL")
		orderURL := os.Getenv("ORDER_SERVICE_URL")
		kitchenURL := os.Getenv("KITCHEN_SERVICE_URL")
		paymentURL := os.Getenv("PAYMENT_SERVICE_URL")
		reservationURL := os.Getenv("RESERVATION_SERVICE_URL")

		// Validate URLs
		urls := map[string]string{
			"Menu":        menuURL + "/menu",
			"Order":       orderURL + "/orders",
			"Kitchen":     kitchenURL + "/kitchen/orders",
			"Payment":     paymentURL + "/payments",
			"Reservation": reservationURL + "/reservations",
		}
		for service, url := range urls {
			if url == "" {
				http.Error(w, fmt.Sprintf("%s service URL is not set", service), http.StatusInternalServerError)
				return
			}
		}

		// Call APIs
		client := &http.Client{Timeout: 10 * time.Second}
		results := []TriggerResponse{}
		for service, url := range urls {
			log.Printf("Calling %s service: %s", service, url)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				results = append(results, TriggerResponse{
					Service: service,
					URL:     url,
					Status:  "Failed",
					Error:   err.Error(),
				})
				log.Printf("Error creating request for %s: %v", service, err)
				continue
			}
			resp, err := client.Do(req)
			if err != nil {
				results = append(results, TriggerResponse{
					Service: service,
					URL:     url,
					Status:  "Failed",
					Error:   err.Error(),
				})
				log.Printf("Error calling %s: %v", service, err)
				continue
			}
			results = append(results, TriggerResponse{
				Service: service,
				URL:     url,
				Status:  resp.Status,
			})
			log.Printf("%s service responded with status: %s", service, resp.Status)
			resp.Body.Close()
		}

		// Return results
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	log.Println("Scheduler Service running on port 8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
