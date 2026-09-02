package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type CartStore struct {
	client *redis.Client
}

func newCartStore() *CartStore {
	host := os.Getenv("VALKEY_HOST")
	port := os.Getenv("VALKEY_PORT")

	address := fmt.Sprintf("%s:%s", host, port)

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	return &CartStore{
		client: client,
	}
}

func (s *CartStore) saveCart(cart Cart) error {
	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	key := "cart:" + cart.UserID

	return s.client.Set(ctx, key, data, 0).Err()
}

func (s *CartStore) getCart(userID string) (Cart, error) {
	key := "cart:" + userID

	data, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return Cart{}, err
	}

	var cart Cart

	if err := json.Unmarshal([]byte(data), &cart); err != nil {
		return Cart{}, err
	}

	return cart, nil
}
