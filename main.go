package main

import (
	"context"
	"log"

	"github.com/branthony-92/log-server/config"
	"github.com/branthony-92/log-server/container"
	"github.com/branthony-92/log-server/routing"
	"github.com/branthony-92/log-server/services"
	"github.com/branthony-92/log-server/services/handlers"
)

const (
	QueueName = "test-queue"
)

func main() {
	config := config.InitConfig()
	var err error
	cont, err := container.InitContainer(config)

	if err != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	services := []services.Service{
		handlers.NewLogUploadService(cont),
	}

	g := routing.RouteServiceHandlers(ctx, cont.GetBroker(), services)
	log.Println("service running...")
	err = g.Wait()
	log.Println(err.Error())
}
