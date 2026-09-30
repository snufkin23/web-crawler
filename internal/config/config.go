package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/snufkin23/web-crawler.git/pkg/logger"
)

type Config struct {
	MongoURI string
}

func NewConfig() *Config {

	if err := godotenv.Load(); err != nil {
		logger.Error("failed to load .env file")
	}

	mongodbURL := os.Getenv("MONGODB_URI")

	if mongodbURL == "" {
		logger.Error("MONGODB_URI is empty")
	}

	return &Config{
		MongoURI: mongodbURL,
	}
}
