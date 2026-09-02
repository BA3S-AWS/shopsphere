package main

type FraudRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type FraudResponse struct {
	Approved  bool   `json:"approved"`
	RiskScore int    `json:"risk_score"`
	Reason    string `json:"reason"`
}
