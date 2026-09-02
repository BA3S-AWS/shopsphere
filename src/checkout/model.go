package main

type CheckoutRequest struct {
	UserID     string `json:"user_id"`
	CardNumber string `json:"card_number"`
	CardExpiry string `json:"card_expiry"`
	CardCVV    string `json:"card_cvv"`
}

type Order struct {
	OrderID       string     `json:"order_id"`
	UserID        string     `json:"user_id"`
	Items         []CartItem `json:"items"`
	ShippingCost  float64    `json:"shipping_cost"`
	EstimatedDays int        `json:"estimated_days"`
	Total         float64    `json:"total"`
	Status        string     `json:"status"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type FraudRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type FraudResponse struct {
	Approved  bool   `json:"approved"`
	RiskScore int    `json:"risk_score"`
	Reason    string `json:"reason"`
}
type ShippingRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type ShippingResponse struct {
	ShippingCost  float64 `json:"shipping_cost"`
	EstimatedDays int     `json:"estimated_days"`
}
