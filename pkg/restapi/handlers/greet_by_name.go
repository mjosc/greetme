package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mjosc/greetme/pkg/restapi/operations/greetings"
	"github.com/mjosc/greetme/pkg/shared"
)

func NewGreetByName(userService shared.UserService) greetings.GreetByNameHandler {
	return &GreetByName{
		UserService: userService,
	}
}

type GreetByName struct {
	UserService shared.UserService
}

func (h GreetByName) Handle(params greetings.GreetByNameParams) middleware.Responder {
	res, err := h.UserService.GetByName(params.Name)
	if err != nil {
		fmt.Println(err)
		return greetings.NewGreetByNameInternalServerError()
	}
	if !res.Valid {
		return greetings.NewGreetByNameBadRequest().WithPayload(&greetings.GreetByNameBadRequestBody{
			Reason:      UserNotFound,
			Description: fmt.Sprintf("A user with the name %v does not exist", params.Name),
		})
	}

	return greetings.NewGreetByNameOK().WithPayload(&greetings.GreetByNameOKBody{
		Greeting: fmt.Sprintf("Hello, %v!", res.Name),
	})
}
