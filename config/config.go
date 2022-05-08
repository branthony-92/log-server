package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

type StorageType string

type StorageConfig struct {
	Type string
	URL  string

	AccessKey string
	SecretKey string
}

type MessageBrokerConfig struct {
	URL string
}

type Config struct {
	Storage   StorageConfig
	Messaging MessageBrokerConfig
}

func InitConfig() Config {
	// load flags
	storageTypeFlag := flag.String("storagetype", "mongo", "Database used to store logs, default mongo")
	flag.Parse()

	// load environment variables

	godotenv.Load()

	// init storage config
	storageAccessKey := os.Getenv("STORAGE_ACCESS_KEY")
	storageSecretKey := os.Getenv("STORAGE_SECRET_KEY")
	storageURL := os.Getenv("STORAGE_URL")

	// init message broker config
	brokerURL := os.Getenv("AMQP_URL")

	c := Config{
		Storage: StorageConfig{
			Type:      *storageTypeFlag,
			URL:       storageURL,
			AccessKey: storageAccessKey,
			SecretKey: storageSecretKey,
		},
		Messaging: MessageBrokerConfig{
			URL: brokerURL,
		},
	}
	return c
}
