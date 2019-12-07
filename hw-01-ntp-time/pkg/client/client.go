package client

import (
	"time"
)

type Client interface {
	GetTime() (time.Time, error)
}
