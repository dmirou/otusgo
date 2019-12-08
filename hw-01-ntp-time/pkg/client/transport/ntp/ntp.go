package ntp

import (
	"github.com/beevik/ntp"
	"time"
)

type Transport struct {
	Host string
}

func NewTransport(host string) Transport {
	return Transport{
		Host: host,
	}
}

func (c Transport) GetTime() (time.Time, error) {
	result, err := ntp.Time(c.Host)
	if err != nil {
		return time.Time{}, err
	}
	return result, nil
}
