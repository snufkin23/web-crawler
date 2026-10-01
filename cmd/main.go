package main

import (
	"context"
	"time"

	httpadapter "github.com/snufkin23/web-crawler.git/internal/adapters/http"
	"github.com/snufkin23/web-crawler.git/internal/application"
	"github.com/snufkin23/web-crawler.git/pkg/logger"
)

func main() {

	logger.Info("Initializing web crawler...")

	// 1. context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// 2. init adapter

	fetcher := httpadapter.NewHTTPFetcher()

	// 3. inject fetcher into application

	crawler := application.NewCrawler(fetcher)

	targetURL := "https://ebpearls.com.au/"
	logger.Info("Crawling target: %s", targetURL)

	// 4. execture crawl to use

	err := crawler.Crawl(ctx, targetURL)
	if err != nil {
		logger.Error("crawl failed: %v", err)
		return
	}

	logger.Success("Crawl successful! Target fetched and validated.")

}
