package container

import (
	broker "github.com/branthony-92/amqp-client"
	"github.com/branthony-92/log-server/storage"
)

type Container interface {
	GetBroker() broker.MessageBroker
	GetStorage() storage.LogStorage
	Shutdown()
}

type container struct {
	broker  broker.MessageBroker
	storage storage.LogStorage
}

func InitContainer(broker broker.MessageBroker) (Container, error) {
	return &container{broker: broker}, nil
}

func (c *container) GetBroker() broker.MessageBroker {
	return c.broker
}

func (c *container) GetStorage() storage.LogStorage {
	return c.storage
}

func (c *container) Shutdown() {
	c.broker.Shutdown()
}

var ServiceContainer Container
