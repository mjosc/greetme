package main

import (
	"flag"

	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/client"
	"github.com/mjosc/greetme/pkg/mocks/mockserver/mockapi"
	"github.com/mjosc/greetme/pkg/server"
)

func main() {

	devMode := flag.Bool("dev", false, "uses the built-in mockserver for all third-party api calls")
	flag.Parse()

	config := &app.Config{
		DevMode: *devMode,
		Port:    8000,
	}

	app := app.New(
		config,
		mockapi.Register,
		client.Register,
		server.Register,
	)

	app.Run()
}
