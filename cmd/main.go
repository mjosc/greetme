package main

import (
	"flag"

	"github.com/mjosc/greetme/pkg/app"
)

func main() {

	devMode := flag.Bool("dev", false, "uses the built-in mockserver for all third-party api calls")
	flag.Parse()

	config := &app.Config{
		DevMode: *devMode,
	}

	app := app.New(
		config,
	)

	app.Run()
}
