package config

import (
	"flag"
	"os"
)

type StorageType string

type StorageConfig struct {
	Type string
	Host string
	Port int

	AccessKey string
	SecretKey string
}

type Config struct {
	Storage StorageConfig
}

func InitConfig() Config {
	// load flags
	storageTypeFlag := flag.String("storagetype", "mongo", "Database used to store logs, default mongo")
	storageHost := flag.String("storagehost", "localhost", "Database hostname to connect to storage")
	storagePort := flag.Int("storagehport", 8080, "Database port to connect to storage")

	// load environment variables
	storageAccessKey := os.Getenv("STORAGE-ACCESS-KEY")
	storageSecretKey := os.Getenv("STORAGE-SECRET-KEY")

	flag.Parse()

	return Config{
		Storage: StorageConfig{
			Type:      *storageTypeFlag,
			Host:      *storageHost,
			Port:      *storagePort,
			AccessKey: storageAccessKey,
			SecretKey: storageSecretKey,
		},
	}
}
