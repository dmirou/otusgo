package ntp

import (
	"github.com/beevik/ntp"
	"time"
)

type Client struct {
	Host string
}

func NewClient(host string) Client {
	return Client{
		Host: host,
	}
}

func (c Client) GetTime() (time.Time, error) {
	result, err := ntp.Time(c.Host)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
