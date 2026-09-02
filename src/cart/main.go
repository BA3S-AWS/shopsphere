package main

import (
	"log"
	"net/http"
)

func main() {
	store := newCartStore()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/cart/", cartHandler(store))

	log.Println("Cart service listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
