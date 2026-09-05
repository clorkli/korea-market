package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func fetchTickers(client *http.Client, markets []string) ([]BithumbTicker, error) {
	marketParam := strings.Join(markets, ",")
	const tickerBaseURL = "https://api.bithumb.com/v1/ticker?markets="
	resp, err := client.Get(tickerBaseURL + marketParam)
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
