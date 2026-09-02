package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/checkout", checkoutHandler)

	log.Println("Checkout service listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
