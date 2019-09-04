package main

import (
	"context"

	"github.com/mjosc/greetme/pkg/client"
	"github.com/mjosc/greetme/pkg/mocks/mockserver/mockapi"
	"github.com/mjosc/greetme/pkg/restapi"
	"github.com/mjosc/greetme/pkg/server"
	"github.com/mjosc/greetme/pkg/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		restapi.FXOptions(),
		mockapi.FXOptions(),
		client.FXOptions(),
		service.FXOptions(),
		// This must be called last in order to properly configure the api (see the invoke methods within each package)
		server.FXOptions(),
	)

	app.Start(context.Background())
	<-app.Done()
	app.Stop(context.Background())
}
