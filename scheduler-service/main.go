package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
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
			log.Fatalf("%s service URL is not set", service)
		}
	}

	// Call APIs
	client := &http.Client{Timeout: 10 * time.Second}
	for service, url := range urls {
		log.Printf("Calling %s service: %s", service, url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Error creating request for %s: %v", service, err)
			continue
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error calling %s: %v", service, err)
			continue
		}
		log.Printf("%s service responded with status: %s", service, resp.Status)
		resp.Body.Close()
	}

	fmt.Println("Scheduler Service execution completed")
}
