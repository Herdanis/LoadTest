package main

import (
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	requestRate := vegeta.Rate{Freq: 4000, Per: time.Second}
	duration := 1 * time.Minute

	target := target()
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(target, requestRate, duration, "Loadtest") {
		metrics.Add(res)
	}
	metrics.Close()

	reporter := vegeta.NewTextReporter(&metrics)
	reporter(os.Stdout)
}

func target() vegeta.Targeter {
	return func(tgt *vegeta.Target) error {
		if tgt == nil {
			return vegeta.ErrNilTarget
		}
		tgt.URL = "https://"
		tgt.Method = "GET"
		tgt.Body = nil

		return nil
	}
}
