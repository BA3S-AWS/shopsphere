package main

import (
	"encoding/json"
	"net/http"
)

func fraudCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req FraudRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	response := FraudResponse{
		Approved:  true,
		RiskScore: 10,
		Reason:    "low risk",
	}

	// Simple fraud rule for the ShopShore demo.
	if req.Amount > 5000 {
		response.Approved = false
		response.RiskScore = 90
		response.Reason = "transaction amount exceeds fraud threshold"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
