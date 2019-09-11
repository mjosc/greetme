package client

import (
	"github.com/mjosc/greetme/pkg/app"
	"go.uber.org/fx"
)

var devMode bool

func Register(config *app.Config) fx.Option {
	devMode = config.DevMode
	return fx.Options(
		fx.Provide(
			NewHTTPClient,
			NewUserService,
		),
	)
}
