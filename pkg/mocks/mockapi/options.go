package mockapi

import (
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	"github.com/mjosc/greetme/pkg/mocks/mockapi/handlers"
	"github.com/mjosc/greetme/pkg/mocks/mockapi/mockops"
	"go.uber.org/fx"
)

func FXOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			newAPI,
			handlers.NewUserByName,
		),
		fx.Invoke(
			setupAPI,
		),
	)
}

func newAPI() (*mockops.MockAPI, error) {
	spec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	return mockops.NewMockAPI(spec), nil
}

func setupAPI(api *mockops.MockAPI, handlers handlers.Handlers) {
	errors.DefaultHTTPCode = 400
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.UsersUserByNameHandler = handlers.UserByName
}
