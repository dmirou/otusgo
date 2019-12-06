package client

import (
	"github.com/beevik/ntp"
	"time"
)

type Client struct {
	Host string
}

func NewClient(host string) *Client {
	return &Client{
		Host: host,
	}
}

func (c *Client) GetDate() (*time.Time, error) {
	response, err := ntp.Query(c.Host)
	if err != nil {
		return nil, err
	}
	result := time.Now().Add(response.ClockOffset)
	return &result, nil
}
