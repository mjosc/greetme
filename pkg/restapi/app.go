package restapi

import (
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/restapi/handlers"
	"github.com/mjosc/greetme/pkg/restapi/operations"
	"go.uber.org/fx"
)

func Register(*app.Config) fx.Option {
	return fx.Options(
		fx.Provide(
			newAPI,
			handlers.NewGreetByName,
		),
		fx.Invoke(
			setupAPI,
		),
	)
}

func newAPI() (*operations.GreetmeAPI, error) {
	spec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	return operations.NewGreetmeAPI(spec), nil
}

func setupAPI(api *operations.GreetmeAPI, handlers handlers.Handlers) {
	errors.DefaultHTTPCode = 400
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.GreetingsGreetByNameHandler = handlers.GreetByName
}
