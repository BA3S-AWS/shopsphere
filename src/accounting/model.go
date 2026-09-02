package main

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Order struct {
	OrderID string     `json:"order_id"`
	UserID  string     `json:"user_id"`
	Items   []CartItem `json:"items"`
	Total   float64    `json:"total"`
	Status  string     `json:"status"`
}
