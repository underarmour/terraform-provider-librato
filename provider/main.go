package provider

import (
	"fmt"
	"io"
	"net/http"
)

type Provider struct {
	user    string
	token   string
	baseUrl string
	Client  http.Client
}

func (p *Provider) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := p.baseUrl + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header["Content-Type"] = []string{"application/json"}
	return req, nil
}

func NewProvider(user, token, baseUrl string) *Provider {
	return &Provider{
		user:    user,
		token:   token,
		baseUrl: baseUrl,
		Client:  http.Client{},
	}
}
