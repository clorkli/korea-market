package main

import (
	"math"
	"testing"
)

func TestChangePct(t *testing.T) {
	tests := []struct {
		name      string
		price     float64
		prevClose float64
		want      float64
		wantErr   bool
	}{
		{
			name:      "price rises",
			price:     110,
			prevClose: 100,
			want:      10,
		},
		{
			name:      "price falls",
			price:     90,
			prevClose: 100,
			want:      -10,
		},
		{
			name:      "unchanged",
			price:     100,
			prevClose: 100,
			want:      0,
		},
		{
			name:      "zero previous close",
			price:     100,
			prevClose: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Quote{
				Price:     tt.price,
				PrevClose: tt.prevClose,
			}

			got, err := changePct(q)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}

func TestBuildSnapshotPartialSuccess(t *testing.T) {
	tickers := []BithumbTicker{
		{
			Market:           "KRW-BTC",
			TradePrice:       110,
			PrevClosingPrice: 100,
		},
		{
			Market:           "KRW-ETH",
			TradePrice:       100,
			PrevClosingPrice: 0,
		},
		{
			Market:           "KRW-XRP",
			TradePrice:       90,
			PrevClosingPrice: 100,
		},
	}

	snapshot := buildSnapshot(tickers)

	if len(snapshot.Entries) != 2 {
		t.Fatalf("got %d entries, want 2", len(snapshot.Entries))
	}
	if len(snapshot.Failures) != 1 {
		t.Fatalf("got %d failures, want 1", len(snapshot.Failures))
	}

	if snapshot.Entries[0].Quote.Instrument.Code != "KRW-BTC" {
		t.Errorf("got %s, want KRW-BTC", snapshot.Entries[0].Quote.Instrument.Code)
	}
	if snapshot.Entries[1].Quote.Instrument.Code != "KRW-XRP" {
		t.Errorf("got %s, want KRW-XRP", snapshot.Entries[1].Quote.Instrument.Code)
	}

	if snapshot.Failures[0].Market != "KRW-ETH" {
		t.Errorf("got %s, want KRW-ETH", snapshot.Failures[0].Market)
	}
	if snapshot.Failures[0].Err == nil {
		t.Error("expected error for KRW-ETH")
	}
}
