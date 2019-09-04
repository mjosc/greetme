package client

import (
	"net/http"
	"time"

	"github.com/mjosc/greetme/pkg/shared"
)

func NewHTTPClient() shared.HTTPClient {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
