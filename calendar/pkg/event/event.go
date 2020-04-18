package event

import (
	"time"
)

type ID string
type UserID string

type Event struct {
	ID           ID
	UserID       UserID
	Title        string
	Desc         string
	Start        time.Time
	End          time.Time
	NotifyBefore *time.Duration
}
