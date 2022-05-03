package container

import (
	broker "github.com/branthony-92/amqp-client"
)

type Container interface {
	GetBroker() broker.MessageBroker
	Shutdown()
}

type container struct {
	broker broker.MessageBroker
}

func InitContainer(broker broker.MessageBroker) (Container, error) {
	return &container{broker: broker}, nil
}

func (c *container) GetBroker() broker.MessageBroker {
	return c.broker
}

func (c *container) Shutdown() {
	c.broker.Shutdown()
}

var ServiceContainer Container
