package handlers

import (
	"github.com/mjosc/greetme/pkg/restapi/operations/greetings"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	GreetByName greetings.GreetByNameHandler
}
