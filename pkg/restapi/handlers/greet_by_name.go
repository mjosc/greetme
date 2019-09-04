package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mjosc/greetme/pkg/restapi/operations/greetings"
)

func NewGreetByName() greetings.GreetByNameHandler {
	return &GreetByName{}
}

type GreetByName struct {
}

func (h GreetByName) Handle(greetings.GreetByNameParams) middleware.Responder {
	return greetings.NewGreetByNameOK().WithPayload(&greetings.GreetByNameOKBody{
		Name:     "matt",
		Greeting: "Hello",
	})
}
