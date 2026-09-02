package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CheckoutRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if request.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	items, err := getCart(request.UserID)
	if err != nil {
		http.Error(w, "Unable to retrieve cart", http.StatusBadGateway)
		return
	}

	products, err := getProducts()
	if err != nil {
		http.Error(w, "Unable to retrieve products", http.StatusBadGateway)
		return
	}

	total := calculateTotal(items, products)
	payment, err := processPayment(PaymentRequest{
		Amount:     total,
		CardNumber: request.CardNumber,
		CardExpiry: request.CardExpiry,
		CardCVV:    request.CardCVV,
	})

	if err != nil {
		http.Error(w, "Payment failed", http.StatusBadGateway)
		return
	}

	if payment.Status != "approved" {
		http.Error(w, "Payment declined", http.StatusPaymentRequired)
		return
	}
	order := Order{
		OrderID: uuid.NewString(),
		UserID:  request.UserID,
		Items:   items,
		Total:   total,
		Status:  "created",
	}

	if err := publishOrder(order); err != nil {
		http.Error(w, "Unable to publish order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func calculateTotal(items []CartItem, products []Product) float64 {
	var total float64

	for _, item := range items {
		for _, product := range products {
			if item.ProductID == product.ID {
				total += product.Price * float64(item.Quantity)
				break
			}
		}
	}

	return total
}
