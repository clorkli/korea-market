package main

type BithumbTicker struct {
	Market           string  `json:"market"`
	OpeningPrice     float64 `json:"opening_price"`
	HighPrice        float64 `json:"high_price"`
	LowPrice         float64 `json:"low_price"`
	TradePrice       float64 `json:"trade_price"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	Change           string  `json:"change"`
	ChangeRate       float64 `json:"change_rate"`
	Timestamp        int64   `json:"timestamp"`
}
