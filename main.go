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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tickers, err := fetchTickers(ctx, client, markets)
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

	for _, entry := range snapshot.Entries {
		fmt.Printf("%s: %.2f, %+.2f%%\n",
			entry.Quote.Instrument.Name,
			entry.Quote.Price,
			entry.ChangePct,
		)
	}

	for _, failure := range snapshot.Failures {
		fmt.Printf("calculate %s failed: %v\n",
			failure.Market,
			failure.Err,
		)
	}
}
