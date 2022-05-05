package storage

import (
	"context"
	"fmt"

	"github.com/branthony-92/log-server/config"
	"github.com/branthony-92/log-server/log"
	"github.com/branthony-92/log-server/storage/mongo"
)

const (
	MongoStorage = "mongo"
)

type LogStorage interface {
	UploadLogs(ctx context.Context, log []log.LogMessage) error
	FindLog(ctx context.Context, logID string) (*log.LogMessage, error)
	DeleteLog(ctx context.Context, logID string) error
}

func InitStorage(cfg config.Config) (LogStorage, error) {
	switch cfg.Storage.Type {
	case MongoStorage:
		storage, err := mongo.NewMongoStorage(cfg)
		if err != nil {
			return nil, fmt.Errorf("could not initialze storage, %v", err)
		}
		return storage, nil
	default:

	}
	return nil, fmt.Errorf("")
}
