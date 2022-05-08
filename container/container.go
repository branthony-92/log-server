package container

import (
	"fmt"

	broker "github.com/branthony-92/amqp-client"
	"github.com/branthony-92/log-server/config"
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

func InitContainer(cfg config.Config) (Container, error) {

	broker, err := broker.NewMessageBroker(cfg.Messaging.URL)
	if err != nil {
		return nil, fmt.Errorf("could not init container, %v", err)
	}

	storage, err := storage.InitStorage(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not init container, %v", err)
	}

	c := container{
		broker: broker, storage: storage,
	}

	return &c, nil
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
