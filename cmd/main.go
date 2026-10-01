package main

import (
	"context"
	"time"

	httpadapter "github.com/snufkin23/web-crawler.git/internal/adapters/http"
	"github.com/snufkin23/web-crawler.git/internal/adapters/storage/mongodb"
	"github.com/snufkin23/web-crawler.git/internal/application"
	"github.com/snufkin23/web-crawler.git/internal/config"
	"github.com/snufkin23/web-crawler.git/pkg/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	logger.Info("Initializing web crawler system...")

	// 1. Load configuration
	cfg := config.NewConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// 2. Connect to MongoDB
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// 3. Select database and collection
	collection := client.Database(mongodb.DatabaseName).Collection(mongodb.CollectionName)

	// 4. Instantiate adapters (The outside world)
	fetcher := httpadapter.NewHTTPFetcher()
	archiver := mongodb.NewMongoArchiver(collection)

	// 5. Inject adapters into application core
	crawler := application.NewCrawler(fetcher, archiver)

	// 6. Execute the crawler workflow
	targetURL := "https://google.com"
	logger.Info("Starting crawl for target: %s", targetURL)

	if err := crawler.Crawl(ctx, targetURL); err != nil {
		logger.Error("Crawl workflow failed: %v", err)
		return
	}

	logger.Success("Crawl successful! Page fetched and archived to MongoDB.")
}
