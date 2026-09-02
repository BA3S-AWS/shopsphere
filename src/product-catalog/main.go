package main

import (
	"log"
	"net/http"
)

func main() {
	db := connectDB()
	defer db.Close()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/products", productsHandler(db))

	log.Println("Product Catalog listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
