package main

import (
	"flag"

	"github.com/mjosc/greetme/pkg/mocks/services"

	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/client"
	"github.com/mjosc/greetme/pkg/restapi"
	"github.com/mjosc/greetme/pkg/server"
)

func main() {

	devMode := flag.Bool("dev", false, "uses the built-in mockserver for all third-party api calls")
	flag.Parse()

	config := &app.Config{
		DevMode:             *devMode,
		Port:                8000,
		MockUserServicePort: 8100,
	}

	registrations := []app.FXRegistrationFunc{
		restapi.Register,
		client.Register,
		server.Register,
	}

	if config.DevMode {
		registrations = append(registrations, services.Register)
	}

	app := app.New(
		config,
		registrations...,
	)

	app.Run()
}
