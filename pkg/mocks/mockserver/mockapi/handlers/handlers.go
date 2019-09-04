package handlers

import (
	"github.com/mjosc/greetme/pkg/mocks/mockserver/mockapi/mockops/users"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	UserByName users.UserByNameHandler
}
