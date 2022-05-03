package storage

const (
	MongoStorage config.StorageType = "mongo"
)

type LogStorage interface {
	UploadLogs(ctx context.Context, log []LogMessage) error
	FindLog(ctx context.Context, logID string) (*LogMessage, error)
	DeleteLog(ctx context.Context, logID string) error
}

func InitStorage(cfg config.Config) (LogStorage, error) {
	switch cfg.Storage.StorageType {
	case MongoStorage:
		storage := NewMongoStorage(cfg)
	default:

	}
	return nil, fmt.Errorf("")
}
