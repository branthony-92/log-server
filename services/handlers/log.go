package handlers

import (
	"context"
	"fmt"
	"time"

	broker "github.com/branthony-92/amqp-client"
	"github.com/branthony-92/log-server/container"
	"github.com/branthony-92/log-server/models"
)

const (
	logUploadServiceKey = "keys.logging.log.received"
	logServiceName      = "Log Handler"
)

const (
	logUploadTimeout = 5 * time.Second
)

type logUploadService struct {
	cont container.Container
}

func NewLogUploadService(cont container.Container) *logUploadService {
	return &logUploadService{cont: cont}
}

func (s *logUploadService) ServiceName() string {
	return logServiceName
}
func (s *logUploadService) ServiceKey() string {
	return logUploadServiceKey
}

func (s *logUploadService) ServiceHandler(msg broker.Message) {
	storage := s.cont.GetStorage()
	var logMsg models.LogMessage

	if err := msg.GetBody(&logMsg); err != nil {
		return
	}

	fmt.Printf("log received: %v\n", logMsg)

	ctx, cancel := context.WithTimeout(context.Background(), logUploadTimeout)
	defer cancel()

	storage.UploadLog(ctx, logMsg)
}
