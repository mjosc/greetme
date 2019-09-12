package user

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi"

	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi/handlers"
	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi/operations"
	"go.uber.org/fx"
)

var configuration *app.Config

// Register is the entry-point to register all components of the mock user service with FX. It should
// be called from services.Register which will make it accessible to the main application.
func Register(config *app.Config) fx.Option {
	configuration = config
	return fx.Options(
		restapi.Register(config),
		fx.Invoke(
			setup,
		),
	)
}

func setup(lc fx.Lifecycle, api *operations.MockUserServiceAPI, handlers handlers.Handlers) {
	apiHandler := api.Serve(nil)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", apiHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", configuration.MockUserServicePort),
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
