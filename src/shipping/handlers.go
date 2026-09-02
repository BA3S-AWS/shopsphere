package main

import (
	"encoding/json"
	"net/http"
)

func shippingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request ShippingRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	shippingCost := 9.90

	if request.Amount >= 100 {
		shippingCost = 0
	}

	response := ShippingResponse{
		ShippingCost:  shippingCost,
		EstimatedDays: 3,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
