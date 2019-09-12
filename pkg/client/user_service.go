package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mjosc/greetme/pkg/shared"
)

func NewUserService(client shared.HTTPClient) shared.UserService {
	var url string
	// This will obviously error with an empty url outside of dev mode
	if configuration.DevMode {
		url = fmt.Sprintf("http://localhost:%v/api/v1", configuration.MockUserServicePort)
	}
	return &UserService{
		HTTPClient: client,
		URL:        url,
	}
}

type UserService struct {
	HTTPClient shared.HTTPClient
	URL        string
}

func (s *UserService) GetByName(n string) (*shared.UserServiceResponse, error) {

	url := fmt.Sprintf("%v/users/%v", s.URL, n)
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
