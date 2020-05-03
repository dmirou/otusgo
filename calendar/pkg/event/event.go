package event

import (
	"time"
)

type Event struct {
	ID           string
	UserID       string
	Title        string
	Desc         string
	Start        time.Time
	End          time.Time
	NotifyBefore time.Duration
}
