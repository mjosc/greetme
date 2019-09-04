package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mjosc/greetme/pkg/shared"
)

func NewUserService(client shared.HTTPClient) shared.UserService {
	return &UserService{
		HTTPClient: client,
	}
}

type UserService struct {
	HTTPClient shared.HTTPClient
}

func (s *UserService) GetByName(n string) (*shared.UserServiceResponse, error) {

	url := fmt.Sprintf("http://localhost:8080/api/v1/mocks/users/%v", n) // TODO: Fix path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("request returned status %v", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result = new(shared.UserServiceResponse)
	if err = json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return result, nil
}
