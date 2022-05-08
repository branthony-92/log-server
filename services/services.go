package services

import (
	broker "github.com/branthony-92/amqp-client"
)

type Service interface {
	ServiceKey() string
	ServiceName() string
	ServiceHandler(msg broker.Message)
}
