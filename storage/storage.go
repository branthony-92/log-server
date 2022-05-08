package storage

import (
	"context"
	"fmt"

	"github.com/branthony-92/log-server/config"
	"github.com/branthony-92/log-server/models"
	"github.com/branthony-92/log-server/storage/mongo"
)

const (
	MongoStorage = "mongo"
)

type LogStorage interface {
	UploadLog(ctx context.Context, log models.LogUploadRequest) error
	FindLogsByDate(ctx context.Context, req models.LogRetrievalRequestDate) ([]models.LogMessage, error)
	DeleteLogsByDate(ctx context.Context, req models.LogRetrievalRequestDate) error
	DeleteLogsByID(ctx context.Context, req models.LogRetrievalRequestID) error
}

func InitStorage(cfg config.Config) (LogStorage, error) {
	switch cfg.Storage.Type {
	case MongoStorage:
		storage, err := mongo.NewMongoStorage(cfg.Storage)
		if err != nil {
			return nil, fmt.Errorf("could not initialze storage, %v", err)
		}
		return storage, nil
	default:

	}
	return nil, fmt.Errorf("")
}
