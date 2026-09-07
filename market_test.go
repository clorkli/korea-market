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
