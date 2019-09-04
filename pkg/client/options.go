package client

import (
	"go.uber.org/fx"
)

func FXOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			NewHTTPClient,
		),
	)
}
