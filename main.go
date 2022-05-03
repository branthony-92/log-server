package main

import (
	"context"
	"os"

	broker "github.com/branthony-92/amqp-client"
)

const (
	QueueName = "test-queue"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	brokerURL := os.Getenv("AMQP-URL")

	b, err := broker.NewMessageBroker(ctx, brokerURL,
		broker.WithQueueName(QueueName),
	)

	if err != nil {
		panic(err.Error())
	}
	defer b.Shutdown()
}
