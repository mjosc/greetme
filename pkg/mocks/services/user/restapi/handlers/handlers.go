package handlers

import (
	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi/operations/users"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	UserByName users.UserByNameHandler
}
