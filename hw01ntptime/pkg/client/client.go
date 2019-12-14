package client

import (
	"time"

	client "github.com/dmirou/otusgo/hw01ntptime/pkg/client/transport"
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
