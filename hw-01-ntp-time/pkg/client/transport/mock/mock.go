package mock

import (
	"time"
)

type GetTime func() (time.Time, error)

type Transport struct {
	ResultTime time.Time
	ResultErr  error
}

func NewTransport(resultTime time.Time, err error) Transport {
	return Transport{
		ResultTime: resultTime,
		ResultErr:  err,
	}
}

func (c *Transport) SetResultTime(resultTime time.Time) {
	c.ResultTime = resultTime
}

func (c *Transport) SetResultErr(err error) {
	c.ResultErr = err
}

func (c Transport) GetTime() (time.Time, error) {
	return c.ResultTime, c.ResultErr
}
