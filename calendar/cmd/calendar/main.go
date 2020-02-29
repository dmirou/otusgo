package main

import (
	"log"

	"github.com/dmirou/otusgo/calendar/pkg/app"
	"github.com/spf13/pflag"
)

func main() {
	config := pflag.String("config", "", "config file path")
	pflag.Parse()

	if *config == "" {
		log.Fatalf("Config file is missing. Please specify it with --config option.")
	}

	a, err := app.New(*config)
	if err != nil {
		log.Fatalf("unexpected error in app.New %v", err)
	}

	a.Run()
}
