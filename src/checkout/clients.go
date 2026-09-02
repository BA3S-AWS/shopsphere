package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func getCart(userID string) ([]CartItem, error) {
	cartURL := os.Getenv("CART_URL")

	resp, err := http.Get(fmt.Sprintf("%s/cart/%s", cartURL, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cart service returned status %d", resp.StatusCode)
	}

	var cart struct {
		Items []CartItem `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&cart); err != nil {
		return nil, err
	}

	return cart.Items, nil
}

func getProducts() ([]Product, error) {
	productCatalogURL := os.Getenv("PRODUCT_CATALOG_URL")

	resp, err := http.Get(productCatalogURL + "/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"product catalog returned status %d",
			resp.StatusCode,
		)
	}

	var products []Product

	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

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

func processPayment(request PaymentRequest) (PaymentResponse, error) {
	paymentURL := os.Getenv("PAYMENT_URL")

	body, err := json.Marshal(request)
	if err != nil {
		return PaymentResponse{}, err
	}

	resp, err := http.Post(
		paymentURL+"/payment",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return PaymentResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return PaymentResponse{}, fmt.Errorf(
			"payment service returned status %d",
			resp.StatusCode,
		)
	}

	var payment PaymentResponse

	if err := json.NewDecoder(resp.Body).Decode(&payment); err != nil {
		return PaymentResponse{}, err
	}

	return payment, nil
}
