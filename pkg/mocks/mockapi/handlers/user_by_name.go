package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mjosc/greetme/pkg/mocks/mockapi/mockops/users"
)

func NewUserByName() users.UserByNameHandler {
	return &UserByName{}
}

type UserByName struct {
}

func (h UserByName) Handle(users.UserByNameParams) middleware.Responder {
	return users.NewUserByNameOK().WithPayload(&users.UserByNameOKBody{
		ID:   1,
		Name: "matt",
	})
}
