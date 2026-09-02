package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func consumeOrders(db *sql.DB) {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "accounting-service",
	})

	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		var order Order

		if err := json.Unmarshal(message.Value, &order); err != nil {
			log.Println("Invalid order message:", err)
			continue
		}

		if err := saveOrder(db, order); err != nil {
			log.Println("Unable to save order:", err)
			continue
		}

		log.Printf("Order %s saved successfully", order.OrderID)
	}
}

func saveOrder(db *sql.DB, order Order) error {
	_, err := db.Exec(`
		INSERT INTO orders (order_id, user_id, total, status)
		VALUES (?, ?, ?, ?)
	`,
		order.OrderID,
		order.UserID,
		order.Total,
		order.Status,
	)

	return err
}
