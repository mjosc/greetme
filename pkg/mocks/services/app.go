package services

import (
	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/mocks/services/user"
	"go.uber.org/fx"
)

func Register(config *app.Config) fx.Option {
	return fx.Options(
		user.Register(config),
	)
}
