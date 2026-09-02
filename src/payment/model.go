package main

type PaymentRequest struct {
	Amount     float64 `json:"amount"`
	CardNumber string  `json:"card_number"`
	CardExpiry string  `json:"card_expiry"`
	CardCVV    string  `json:"card_cvv"`
}

type PaymentResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}
