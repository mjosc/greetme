package service

import (
	"github.com/mjosc/greetme/internal"
	"go.uber.org/fx"
)

var devMode bool

func FXOptions(config *internal.Config) fx.Option {
	devMode = config.DevMode
	return fx.Options(
		fx.Provide(
			NewUserService,
		),
	)
}
