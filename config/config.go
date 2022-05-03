package config

type StorageType string

type StorageConfig struct {
	StorageType StorageType
	StorageHost string
	StoragePort int

	StorageAccessKey string
	StorageSecretKey string
}

type Config struct {
	Storage StorageConfig
}
