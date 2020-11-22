package asana

import "time"

func WithBaseURL(url string) ClientOption {
	return func(c *Client) *Client {
		c.BaseURL = url
		return c
	}
}

func WithUserAgent(agent string) ClientOption {
	return func(c *Client) *Client {
		c.UserAgent = agent
		return c
	}
}

func WithToken(token string) ClientOption {
	return func(c *Client) *Client {
		c.Token = token
		return c
	}
}

func WithTimeout(t time.Duration) ClientOption {
	return func(c *Client) *Client {
		c.httpClient.Timeout = t
		return c
	}
}
