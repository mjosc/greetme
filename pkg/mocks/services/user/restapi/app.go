package restapi

import (
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi/handlers"
	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi/operations"
	"go.uber.org/fx"
)

func Register(*app.Config) fx.Option {
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

func newAPI() (*operations.MockUserServiceAPI, error) {
	spec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	return operations.NewMockUserServiceAPI(spec), nil
}

func setupAPI(api *operations.MockUserServiceAPI, handlers handlers.Handlers) {
	errors.DefaultHTTPCode = 400
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.UsersUserByNameHandler = handlers.UserByName
}
