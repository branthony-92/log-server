package main

import (
	"context"
	"os"
	"time"

	broker "github.com/branthony-92/amqp-client"
	"github.com/branthony-92/log-server/container"
	logging "github.com/branthony-92/log-server/log"
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

	cont, err := container.InitContainer(b)

	b.RegisterRoute("logs", "logging.logs.#", "", func(msg broker.Message) {
		var log logging.LogMessage

		if err := msg.GetBody(&log); err != nil {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := cont.GetStorage().UploadLogs(ctx, []logging.LogMessage{log}); err != nil {
			return
		}
	})

	if err != nil {
		panic(err.Error())
	}
	defer b.Shutdown()
}
