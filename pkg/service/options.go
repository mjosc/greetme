package service

import (
	"github.com/mjosc/greetme/pkg/app"
	"go.uber.org/fx"
)

var devMode bool

func FXOptions(config *app.Config) fx.Option {
	devMode = config.DevMode
	return fx.Options(
		fx.Provide(
			NewUserService,
		),
	)
}
