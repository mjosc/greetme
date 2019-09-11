package main

import (
	"context"
	"flag"

	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/client"
	"github.com/mjosc/greetme/pkg/mocks/mockserver/mockapi"
	"github.com/mjosc/greetme/pkg/restapi"
	"github.com/mjosc/greetme/pkg/server"
	"github.com/mjosc/greetme/pkg/service"
	"go.uber.org/fx"
)

func main() {

	devMode := flag.Bool("dev", false, "uses the built-in mockserver for all third-party api calls")
	flag.Parse()

	config := app.Config{
		DevMode: *devMode,
	}

	app := fx.New(
		restapi.FXOptions(),
		mockapi.FXOptions(),
		client.FXOptions(),
		service.FXOptions(&config),
		// This must be called last in order to properly configure the api (see the invoke methods within each package)
		server.FXOptions(&config),
	)

	app.Start(context.Background())
	<-app.Done()
	app.Stop(context.Background())
}
