package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request PaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if request.Amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	if !validCard(request.CardNumber, request.CardExpiry, request.CardCVV) {
		http.Error(w, "Payment declined", http.StatusPaymentRequired)
		return
	}

	response := PaymentResponse{
		TransactionID: uuid.NewString(),
		Status:        "approved",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func validCard(number, expiry, cvv string) bool {
	number = strings.ReplaceAll(number, " ", "")

	return len(number) >= 12 &&
		expiry != "" &&
		len(cvv) >= 3
}
