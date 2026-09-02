package main

type ShippingRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type ShippingResponse struct {
	ShippingCost  float64 `json:"shipping_cost"`
	EstimatedDays int     `json:"estimated_days"`
}
