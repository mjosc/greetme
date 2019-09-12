package mocks

import (
	"errors"

	"github.com/mjosc/greetme/pkg/shared"
)

func NewUserService() *UserService {
	return &UserService{}
}

type UserService struct {
	HandleGetByName func(string) (*shared.UserServiceResponse, error)
}

func (m *UserService) GetByName(n string) (*shared.UserServiceResponse, error) {
	if m.HandleGetByName == nil {
		return nil, errors.New("UserService.HandleGetByName is not intitialized")
	}
	return m.HandleGetByName(n)
}
