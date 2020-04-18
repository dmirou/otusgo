package error

import (
	"fmt"
	"time"

	"github.com/dmirou/otusgo/calendar/pkg/event"
)

type InvalidArgError struct {
	Name   string
	Method string
	Desc   string
}

func (ia *InvalidArgError) Error() string {
	return fmt.Sprintf("invalid argument %q in method %q: %s", ia.Name, ia.Method, ia.Desc)
}

type EventNotFoundError struct {
	EventID event.ID
}

func (enf *EventNotFoundError) Error() string {
	return fmt.Sprintf("event not found by id: %q", enf.EventID)
}

type DateBusyError struct {
	Start time.Time
	End   time.Time
}

func (db *DateBusyError) Error() string {
	return fmt.Sprintf("date already busy: (%s, %s)", db.Start, db.End)
}
