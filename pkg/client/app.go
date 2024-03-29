package client

import (
	"github.com/mjosc/greetme/pkg/app"
	"go.uber.org/fx"
)

var configuration *app.Config

func Register(config *app.Config) fx.Option {
	configuration = config
	return fx.Options(
		fx.Provide(
			NewHTTPClient,
			NewUserService,
		),
	)
}
