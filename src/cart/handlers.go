package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func cartHandler(store *CartStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := strings.TrimPrefix(r.URL.Path, "/cart/")

		if userID == "" {
			http.Error(w, "user_id is required", http.StatusBadRequest)
			return
		}

		switch r.Method {

		case http.MethodGet:
			cart, err := store.getCart(userID)
			if err != nil {
				http.Error(w, "Cart not found", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cart)

		case http.MethodPost:
			var incomingCart Cart

			if err := json.NewDecoder(r.Body).Decode(&incomingCart); err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			cart, err := store.getCart(userID)
			if err != nil {
				cart = Cart{
					UserID: userID,
					Items:  []CartItem{},
				}
			}

			for _, incomingItem := range incomingCart.Items {
				found := false

				for i := range cart.Items {
					if cart.Items[i].ProductID == incomingItem.ProductID {
						cart.Items[i].Quantity += incomingItem.Quantity
						found = true
						break
					}
				}

				if !found {
					cart.Items = append(cart.Items, incomingItem)
				}
			}

			if err := store.saveCart(cart); err != nil {
				http.Error(w, "Unable to save cart", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(cart)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
