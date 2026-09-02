package main

type CheckoutRequest struct {
	UserID      string `json:"user_id"`
	CardNumber  string `json:"card_number"`
	CardExpiry  string `json:"card_expiry"`
	CardCVV     string `json:"card_cvv"`
}

type Order struct {
	OrderID string     `json:"order_id"`
	UserID  string     `json:"user_id"`
	Items   []CartItem `json:"items"`
	Total   float64    `json:"total"`
	Status  string     `json:"status"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
