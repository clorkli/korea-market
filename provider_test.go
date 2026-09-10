package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchTickersSuccess(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `[{"market": "KRW-BTC", "trade_price": 110, "prev_closing_price": 100}]`)
			gotMarkets := r.URL.Query().Get("markets")
			if gotMarkets != "KRW-BTC" {
				t.Errorf("got markets query %q, want %q", gotMarkets, "KRW-BTC")
			}
		}),
	)
	defer server.Close()

	provider := BithumbProvider{
		client:  server.Client(),
		baseURL: server.URL + "?markets=",
	}
	ctx := context.Background()

	tickers, err := provider.fetchTickers(
		ctx,
		[]string{"KRW-BTC"},
	)

	if err != nil {
		t.Fatalf("fetchTickers returned error: %v", err)
	}
	if len(tickers) != 1 {
		t.Fatalf("got %d tickers, want 1", len(tickers))
	}
	if tickers[0].Market != "KRW-BTC" {
		t.Errorf("got market %q, want %q",
			tickers[0].Market,
			"KRW-BTC",
		)
	}
	if tickers[0].TradePrice != 110 {
		t.Errorf("got trade price %.2f, want 110", tickers[0].TradePrice)
	}
	if tickers[0].PrevClosingPrice != 100 {
		t.Errorf(
			"got previous close %.2f, want 100",
			tickers[0].PrevClosingPrice,
		)
	}
}

func TestFetchTickersErrors(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantErr    string
	}{
		{
			name:       "server error",
			statusCode: http.StatusInternalServerError,
			body:       "internal server error",
			wantErr:    "unexpected status",
		},
		{
			name:       "invalid json",
			statusCode: http.StatusOK,
			body:       "{",
			wantErr:    "decode response body",
		},
		{
			name:       "empty response",
			statusCode: http.StatusOK,
			body:       "[]",
			wantErr:    "empty ticker response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				fmt.Fprint(w, tt.body)
			}))
			defer server.Close()

			provider := BithumbProvider{
				client:  server.Client(),
				baseURL: server.URL + "?markets=",
			}

			_, err := provider.fetchTickers(context.Background(), []string{"KRW-BTC"})
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("got error %q, want containing %q", err, tt.wantErr)
			}
		})
	}
}
