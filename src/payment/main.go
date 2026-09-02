package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/payment", paymentHandler)

	log.Println("Payment service listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
