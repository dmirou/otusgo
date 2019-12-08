package client

import (
	"time"

	client "github.com/dmirou/otus-go/hw-01-ntp-time/pkg/client/transport"
)

type Client struct {
	transport client.Transport
}

func NewClient(transport client.Transport) Client {
	return Client{
		transport: transport,
	}
}

func (c Client) GetTime() (time.Time, error) {
	return c.transport.GetTime()
}
