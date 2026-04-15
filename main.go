package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for the evcc application.
// evcc is an EV Charge Controller that manages electric vehicle charging
// based on solar surplus, grid pricing, and other configurable parameters.
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
