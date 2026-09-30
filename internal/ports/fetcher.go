package ports

import "context"

type Result struct {
	Body string
}

type Fetcher interface {
	Fetch(ctx context.Context, url string) (Result, error)
}
