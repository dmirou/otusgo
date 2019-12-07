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

func (c *Client) GetTime() (*time.Time, error) {
	time, err := ntp.Time(c.Host)
	if err != nil {
		return nil, err
	}
	return &time, nil
}
