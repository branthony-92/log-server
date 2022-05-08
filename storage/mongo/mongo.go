package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/branthony-92/log-server/config"
	"github.com/branthony-92/log-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName         = "Log-Database"
	collectionName = "logs"
)

type mongoStorage struct {
	client *mongo.Client
}

func NewMongoStorage(cfg config.StorageConfig) (*mongoStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URL))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	storage := mongoStorage{
		client: client,
	}
	return &storage, nil
}

func (s *mongoStorage) UploadLog(ctx context.Context, log models.LogMessage) error {

	collection := s.client.Database(dbName).Collection(log.Source)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, &log); err != nil {
		return fmt.Errorf("Failed to upload logs, %v", err)
	}

	return nil
}

func (s *mongoStorage) FindLog(ctx context.Context, logID string) (*models.LogMessage, error) {
	collection := s.client.Database(dbName).Collection(collectionName)

	filter := bson.D{
		{Key: "id", Value: logID},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, filter)

	var log models.LogMessage

	if err := result.Decode(&log); err != nil {
		return nil, fmt.Errorf("Failed to retrieve log, %v", err)
	}

	return &log, nil
}

func (s *mongoStorage) DeleteLog(ctx context.Context, logID string) error {
	collection := s.client.Database(dbName).Collection(collectionName)

	filter := bson.D{
		{Key: "id", Value: logID},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("Failed to delete log, %v", err)
	}

	return nil
}
