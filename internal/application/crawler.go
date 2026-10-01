package application

import (
	"context"
	"fmt"
	"time"

	"github.com/snufkin23/web-crawler.git/internal/domain"
	"github.com/snufkin23/web-crawler.git/internal/ports"
)

type Crawler struct {
	fetcher  ports.Fetcher
	archiver ports.Archiver
}

func NewCrawler(fetcher ports.Fetcher, archiver ports.Archiver) *Crawler {
	return &Crawler{
		fetcher:  fetcher,
		archiver: archiver,
	}
}

func (c *Crawler) Crawl(ctx context.Context, targetURl string) error {

	result, err := c.fetcher.Fetch(ctx, targetURl)
	if err != nil {
		return fmt.Errorf("crawler failed to fetch target: %w", err)
	}

	page := &domain.Page{
		URL:  targetURl,
		Time: time.Now(),
		Body: []byte(result.Body),
	}

	if _, err := c.archiver.Save(ctx, page); err != nil {
		return fmt.Errorf("crawler failed to archive page: %w", err)
	}

	return nil
}
