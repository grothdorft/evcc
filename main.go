package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for the evcc application.
// evcc is an EV Charge Controller that manages electric vehicle charging
// based on solar surplus, grid pricing, and other configurable parameters.
//
// Personal fork: using this to manage charging for my VW ID.4 and home solar setup.
// Upstream repo: https://github.com/evcc-io/evcc
//
// Notes to self:
//   - config file lives at ~/.config/evcc/evcc.yaml
//   - run with `evcc --log debug` to get verbose output for troubleshooting
//   - ID.4 sometimes needs a manual wake before charging sessions start
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
