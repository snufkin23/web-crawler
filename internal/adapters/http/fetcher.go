package httpadapter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/snufkin23/web-crawler.git/internal/domain"
	"github.com/snufkin23/web-crawler.git/internal/ports"
)

type HTTPFetcher struct {
	client *http.Client
}

func NewHTTPFetcher() *HTTPFetcher {
	return &HTTPFetcher{
		client: &http.Client{},
	}
}

func (f *HTTPFetcher) Fetch(ctx context.Context, rawURL string) (ports.Result, error) {

	parsedURL, err := url.Parse(rawURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return ports.Result{}, domain.ErrUnsupportedScheme
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return ports.Result{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return ports.Result{}, fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") && !strings.Contains(contentType, "application/xhtml+xml") {
		return ports.Result{}, domain.ErrNonHTML
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ports.Result{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return ports.Result{
		Body: string(bodyBytes),
	}, nil

}
