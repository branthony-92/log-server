package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	client *mongo.Client
}

func NewMongoStorage(cfg config.Config) (*mongoStorage, error) {
	return nil, nil
}

func (s *mongoStorage) UploadLogs(ctx context.Context, log []LogMessage) error {
	return nil
}

func (s *mongoStorage) FindLog(ctx context.Context, logID string) (*LogMessage, error) {
	return nil, nil
}

func (s *mongoStorage) DeleteLog(ctx context.Context, logID string) error {
	return nil
}
