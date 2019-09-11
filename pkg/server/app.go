package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/mocks/mockserver/mockapi/mockops"
	"github.com/mjosc/greetme/pkg/restapi/handlers"
	"github.com/mjosc/greetme/pkg/restapi/operations"
	"go.uber.org/fx"
)

const port = 8080

var devMode bool

func Register(config *app.Config) fx.Option {
	devMode = config.DevMode
	return fx.Options(
		fx.Invoke(
			setup,
		),
	)
}

func setup(lc fx.Lifecycle, api *operations.GreetmeAPI, mockapi *mockops.MockAPI, handlers handlers.Handlers) {
	apiHandler := api.Serve(nil)
	mockapiHandler := mockapi.Serve(nil)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", apiHandler)

	if devMode {
		// We could also just register this handler with a different server instance with a different port and
		// run it in parallel.
		mux.Handle("/api/v1/mocks/", mockapiHandler)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					if err := server.ListenAndServe(); err != nil {
						if err == http.ErrServerClosed {
							log.Println("Server stopped")
						} else {
							log.Println(err, "Error starting server")
						}
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				if err := server.Shutdown(ctx); err != nil {
					log.Println(err, "Error shutting down server")
					return err
				}
				return nil
			},
		},
	)
}
