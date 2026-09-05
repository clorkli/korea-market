package main

import (
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

	tickers, err := fetchTickers(client, markets)
	if err != nil {
		fmt.Printf("fetch failed: %v\n", err)
		return
	}

	for _, ticker := range tickers {
		quote := tickerToQuote(ticker)
		pct, err := changePct(quote)
		if err != nil {
			fmt.Printf("calculate %s failed: %v\n", ticker.Market, err)
			continue
		}

		fmt.Printf("%s: %.2f, %+.2f%%\n",
			quote.Instrument.Name,
			quote.Price,
			pct,
		)
	}

}
