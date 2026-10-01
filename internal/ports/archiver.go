package ports

import (
	"context"

	"github.com/snufkin23/web-crawler.git/internal/domain"
)

type Archiver interface {
	Save(ctx context.Context, page *domain.Page) (bool, error)
}
