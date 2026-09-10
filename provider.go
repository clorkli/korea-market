package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const bithumbTickerBaseURL = "https://api.bithumb.com/v1/ticker?markets="

type BithumbProvider struct {
	client  *http.Client
	baseURL string
}

func (p *BithumbProvider) fetchTickers(ctx context.Context, markets []string) ([]BithumbTicker, error) {
	marketParam := strings.Join(markets, ",")
	url := p.baseURL + marketParam

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create ticker request: %w", err)
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request ticker: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	var tickers []BithumbTicker
	if err = json.Unmarshal(body, &tickers); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}
	if len(tickers) == 0 {
		return nil, fmt.Errorf("empty ticker response")
	}

	return tickers, nil
}
