package event

import "github.com/dmirou/otusgo/calendar/pkg/time"

type ID string

type Event struct {
	ID    ID
	Title string
	Desc  string
	Start *time.Time
	End   *time.Time
}
