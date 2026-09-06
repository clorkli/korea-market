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

type SnapshotEntry struct {
	Quote     Quote
	ChangePct float64
}

type Snapshot struct {
	Entries  []SnapshotEntry
	Failures []SnapshotFailure
}

type SnapshotFailure struct {
	Market string
	Err    error
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

func tickerToSnapshotEntry(t BithumbTicker) (SnapshotEntry, error) {
	q := tickerToQuote(t)
	pct, err := changePct(q)
	if err != nil {
		return SnapshotEntry{}, err
	}

	return SnapshotEntry{
		Quote:     q,
		ChangePct: pct,
	}, nil
}

func buildSnapshot(tickers []BithumbTicker) Snapshot {
	var snapshot Snapshot
	for _, ticker := range tickers {
		entry, err := tickerToSnapshotEntry(ticker)
		if err != nil {
			snapshot.Failures = append(
				snapshot.Failures,
				SnapshotFailure{
					Market: ticker.Market,
					Err:    err,
				},
			)
			continue
		}
		snapshot.Entries = append(snapshot.Entries, entry)
	}
	return snapshot
}

func changePct(q Quote) (float64, error) {
	if q.PrevClose == 0 {
		return 0, errors.New("previous close is zero")
	}

	return (q.Price - q.PrevClose) / q.PrevClose * 100, nil
}
