package services

import (
	"github.com/mjosc/greetme/pkg/app"
	"github.com/mjosc/greetme/pkg/mocks/services/user"
	"go.uber.org/fx"
)

// Register is the entry-point to register all mock services with FX. Any additional mock services
// should be registered here, rather than external to this package.
func Register(config *app.Config) fx.Option {
	return fx.Options(
		user.Register(config),
	)
}
