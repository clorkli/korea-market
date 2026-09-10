package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	markets := []string{
		"KRW-BTC",
		"KRW-ETH",
		"KRW-XRP",
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	provider := BithumbProvider{
		client:  client,
		baseURL: bithumbTickerBaseURL,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tickers, err := provider.fetchTickers(ctx, markets)
	if err != nil {
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Println("fetch timed out")
		case errors.Is(err, context.Canceled):
			fmt.Println("fetch canceled")
		default:
			fmt.Printf("fetch failed: %v\n", err)
		}
		return
	}

	snapshot := buildSnapshot(tickers)
	printSnapshot(snapshot)
}
