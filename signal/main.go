package main

import (
	"github.com/FlintyLemming/netbird/signal/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
