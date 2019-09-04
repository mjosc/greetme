package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mjosc/greetme/pkg/db/entities"
	"github.com/mjosc/greetme/pkg/mocks/mockapi/mockops/users"
)

const UserNotFound = "UserNotFound"

var database = map[string]*entities.User{
	"Matt": &entities.User{
		ID:   1,
		Name: "Matt",
	},
}

func NewUserByName() users.UserByNameHandler {
	return &UserByName{}
}

type UserByName struct {
}

func (h UserByName) Handle(params users.UserByNameParams) middleware.Responder {
	var responder = users.NewUserByNameOK()

	user, ok := database[params.Name]
	if !ok {
		return responder.WithPayload(&users.UserByNameOKBody{
			Valid:   false,
			Error:   UserNotFound,
			Message: fmt.Sprintf("A user with the name %v does not exist", params.Name),
		})
	}
	return users.NewUserByNameOK().WithPayload(&users.UserByNameOKBody{
		Valid: true,
		ID:    user.ID,
		Name:  user.Name,
	})
}
