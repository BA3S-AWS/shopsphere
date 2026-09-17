package main

import (
	"encoding/json"
	"net/http"
        "strconv"
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
                case http.MethodPut:
                        var updatedCart Cart

                        if err := json.NewDecoder(r.Body).Decode(&updatedCart); err != nil {
                                http.Error(w, "Invalid request", http.StatusBadRequest)
                                return
                        }

                        updatedCart.UserID = userID

                        if err := store.saveCart(updatedCart); err != nil {
                                http.Error(w, "Unable to update cart", http.StatusInternalServerError)
                                return
                        }

                        w.Header().Set("Content-Type", "application/json")
                        json.NewEncoder(w).Encode(updatedCart)

                case http.MethodDelete:
                        productIDValue := r.URL.Query().Get("product_id")

                        // Sans product_id : vider entièrement le panier.
                        if productIDValue == "" {
                                if err := store.deleteCart(userID); err != nil {
                                        http.Error(w, "Unable to delete cart", http.StatusInternalServerError)
                                        return
                                }

                                w.WriteHeader(http.StatusNoContent)
                                return
                        }

                        // Avec product_id : supprimer uniquement cet article.
                        productID, err := strconv.Atoi(productIDValue)
                        if err != nil {
                                http.Error(w, "Invalid product_id", http.StatusBadRequest)
                                return
                        }

                        cart, err := store.getCart(userID)
                        if err != nil {
                                http.Error(w, "Cart not found", http.StatusNotFound)
                                return
                        }

                        remainingItems := make([]CartItem, 0, len(cart.Items))

                        for _, item := range cart.Items {
                                if item.ProductID != productID {
                                        remainingItems = append(remainingItems, item)
                                }
                        }

                        cart.Items = remainingItems

                        if len(cart.Items) == 0 {
                                if err := store.deleteCart(userID); err != nil {
                                        http.Error(w, "Unable to delete cart", http.StatusInternalServerError)
                                        return
                                }
                        } else if err := store.saveCart(cart); err != nil {
                                http.Error(w, "Unable to update cart", http.StatusInternalServerError)
                                return
                        }

                        w.Header().Set("Content-Type", "application/json")
                        json.NewEncoder(w).Encode(cart)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
