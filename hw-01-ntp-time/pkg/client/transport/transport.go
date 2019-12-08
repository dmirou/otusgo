package client

import "time"

type Transport interface {
	GetTime() (time.Time, error)
}
