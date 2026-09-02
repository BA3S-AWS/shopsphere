package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func productsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(`
			SELECT id, name, description, price, stock
			FROM products
		`)

		if err != nil {
			http.Error(w, "Unable to retrieve products", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		products := []Product{}

		for rows.Next() {
			var product Product

			err := rows.Scan(
				&product.ID,
				&product.Name,
				&product.Description,
				&product.Price,
				&product.Stock,
			)

			if err != nil {
				http.Error(w, "Unable to read product", http.StatusInternalServerError)
				return
			}

			products = append(products, product)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
