package handlers

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mjosc/greetme/pkg/mocks/services/user/restapi/operations/users"
)

func NewUserByName() users.UserByNameHandler {
	return &UserByName{}
}

type UserByName struct {
}

func (h UserByName) Handle(params users.UserByNameParams) middleware.Responder {
	var responder = users.NewUserByNameOK()
	name := strings.Title(params.Name)
	if name != "Matt" {
		return responder.WithPayload(&users.UserByNameOKBody{
			Valid:   false,
			Error:   "UserNotFound",
			Message: fmt.Sprintf("A user with the name %v does not exist", name),
			Name:    name,
		})
	}
	return responder.WithPayload(&users.UserByNameOKBody{
		Valid: true,
		ID:    1,
		Name:  name,
	})
}
