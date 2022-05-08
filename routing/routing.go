package routing

import (
	"context"
	"fmt"

	broker "github.com/branthony-92/amqp-client"
	"github.com/branthony-92/log-server/services"
	"golang.org/x/sync/errgroup"
)

const (
	consumerName = "log-server"
	queueName    = "log-server-queue"
	exchangeName = "log-server-exchange"
)

func RouteServiceHandlers(ctx context.Context, broker broker.MessageBroker, services []services.Service) *errgroup.Group {
	group, ctx := errgroup.WithContext(ctx)

	broker.CreateQueue(queueName)
	broker.CreateExchange(exchangeName)

	group.Go(func() error {
		consumer := broker.CreateConsumer(consumerName)
		for _, s := range services {
			if err := consumer.RegisterRoute(queueName, s.ServiceKey(), exchangeName, s.ServiceHandler); err != nil {
				return fmt.Errorf("could not register route [%s], %v", s.ServiceKey(), err)
			}
		}
		return consumer.StartConsuming(ctx, queueName)
	})
	return group
}
