package main

import (
	"context"

	"github.com/mjosc/greetme/pkg/mocks/mockapi"
	"github.com/mjosc/greetme/pkg/restapi"
	"github.com/mjosc/greetme/pkg/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		// The order of invoke options (as contained in each FXOptions call) is critical.
		restapi.FXOptions(),
		mockapi.FXOptions(),
		server.FXOptions(),
	)

	app.Start(context.Background())
	<-app.Done()
	app.Stop(context.Background())
}
