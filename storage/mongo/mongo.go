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
	dbName = "Log-Database"
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

func (s *mongoStorage) UploadLog(ctx context.Context, log models.LogUploadRequest) error {

	collection := s.client.Database(dbName).Collection(log.Source)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, &log.Message); err != nil {
		return fmt.Errorf("Failed to upload logs, %v", err)
	}

	return nil
}

func (s *mongoStorage) FindLogsByDate(ctx context.Context, req models.LogRetrievalRequestDate) ([]models.LogMessage, error) {
	collection := s.client.Database(dbName).Collection(req.Source)

	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{{Key: "$timestamp", Value: bson.D{{Key: "$gte", Value: req.RangeStart}}}},
				bson.D{{Key: "$timestamp", Value: bson.D{{Key: "$lte", Value: req.RangeEnd}}}},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []models.LogMessage{}, fmt.Errorf("could not retrieve logs, %v", err)
	}
	defer cursor.Close(context.TODO())

	logs := make([]models.LogMessage, 0)

	for cursor.Next(context.TODO()) {
		var log models.LogMessage
		if err := cursor.Decode(&log); err != nil {
			return []models.LogMessage{}, fmt.Errorf("could not decode log, %v", err)
		}
		logs = append(logs, log)
	}

	return logs, nil
}

func (s *mongoStorage) DeleteLogsByDate(ctx context.Context, req models.LogRetrievalRequestDate) error {
	collection := s.client.Database(dbName).Collection(req.Source)

	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{{Key: "$timestamp", Value: bson.D{{Key: "$gte", Value: req.RangeStart}}}},
				bson.D{{Key: "$timestamp", Value: bson.D{{Key: "$lte", Value: req.RangeEnd}}}},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.DeleteMany(ctx, filter); err != nil {
		return fmt.Errorf("Failed to delete logs, %v", err)
	}

	return nil
}

func (s *mongoStorage) DeleteLogsByID(ctx context.Context, req models.LogRetrievalRequestID) error {
	collection := s.client.Database(dbName).Collection(req.Source)

	for _, id := range req.Source {
		filter := bson.D{
			{Key: "id", Value: id},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if _, err := collection.DeleteOne(ctx, filter); err != nil {
			return fmt.Errorf("Failed to delete log, %v", err)
		}
	}
	return nil
}
