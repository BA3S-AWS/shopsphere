package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/segmentio/kafka-go"
)

func publishOrder(order Order) error {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	defer writer.Close()

	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(order.OrderID),
			Value: data,
		},
	)
}
