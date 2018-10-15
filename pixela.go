package pixela

import (
	"net/http"
	"time"
)

const baseURL = "https://pixe.la/v1"

//Client is Client for access to pixe.la service
type Client struct {
	HTTPClient http.Client
	URL        string
	UserName   string
	Token      string
}

//NewClient is return Cilent
func NewClient() *Client {
	return &Client{
		HTTPClient: http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
		URL: baseURL,
	}
}
