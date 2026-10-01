package application

import (
	"context"
	"fmt"

	"github.com/snufkin23/web-crawler.git/internal/ports"
)

type Crawler struct {
	fetcher ports.Fetcher
}

func NewCrawler(fetcher ports.Fetcher) *Crawler {
	return &Crawler{
		fetcher: fetcher,
	}
}

func (c *Crawler) Crawl(ctx context.Context, targetURl string) error {

	_, err := c.fetcher.Fetch(ctx, targetURl)
	if err != nil {
		return fmt.Errorf("crawler failed to fetch target: %w", err)
	}

	return nil

}
