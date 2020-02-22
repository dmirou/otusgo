package event

import "time"

type ID string

type Event struct {
	ID     ID
	Title  string
	Desc   string
	Start  *time.Time
	End    *time.Time
	AllDay bool
}
