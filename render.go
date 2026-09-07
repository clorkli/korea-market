package main

import "fmt"

func printSnapshot(snapshot Snapshot) {
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
