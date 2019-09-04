package handlers

import (
	"github.com/mjosc/greetme/pkg/mocks/mockapi/mockops/users"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	UserByName users.UserByNameHandler
}
