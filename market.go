package main

import "errors"

type InstrumentKind string

const (
	KindStock  InstrumentKind = "stock"
	KindIndex  InstrumentKind = "index"
	KindCrypto InstrumentKind = "crypto"
)

type Instrument struct {
	Code string
	Name string
	Kind InstrumentKind
}

type Quote struct {
	Instrument Instrument
	Price      float64
	PrevClose  float64
}

func tickerToQuote(t BithumbTicker) Quote {
	return Quote{
		Instrument: Instrument{
			Code: t.Market,
			Name: t.Market,
			Kind: KindCrypto,
		},
		Price:     t.TradePrice,
		PrevClose: t.PrevClosingPrice,
	}
}

func changePct(q Quote) (float64, error) {
	if q.PrevClose == 0 {
		return 0, errors.New("previous close is zero")
	}

	return (q.Price - q.PrevClose) / q.PrevClose * 100, nil
}
