package mock

import (
	"time"
)

type GetTime func() (time.Time, error)

type Client struct {
	callback GetTime
}

func NewClient(callback GetTime) Client {
	return Client{
		callback: callback,
	}
}

func (c Client) GetTime() (time.Time, error) {
	return c.callback()
}
