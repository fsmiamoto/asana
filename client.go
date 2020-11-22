package asana

import (
	"io"
	"net/http"
	"time"
)

type ClientOption func(c *Client) *Client

type Client struct {
	BaseURL    string
	Token      string
	UserAgent  string
	Timeout    time.Duration
	httpClient *http.Client
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		httpClient: &http.Client{},
	}
	for _, option := range opts {
		c = option(c)
	}
	return c
}

func (c *Client) buildRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}
