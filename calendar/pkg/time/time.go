package time

import "time"

type Time struct {
	t time.Time
}

// New creates new date in UTC location
func New(year, month, day, hour, min int) *Time {
	return &Time{
		t: time.Date(year, time.Month(month), day, hour, min, 0, 0, time.UTC),
	}
}

func (t *Time) After(u Time) bool {
	return t.t.After(u.t)
}

func (t *Time) Before(u Time) bool {
	return t.t.Before(u.t)
}

func (t *Time) HasDate(year, month, day int) bool {
	start := New(year, month, day-1, 23, 59)
	end := New(year, month, day+1, 0, 0)

	return t.After(*start) && t.Before(*end)
}

// Inside checks that t inside (start, end), not including borders.
func (t *Time) Inside(start, end Time) bool {
	if start.Equals(end) || t.Equals(start) || t.Equals(end) {
		return false
	}

	if start.Before(end) {
		switch {
		case t.Before(start):
			return false
		case t.After(end):
			return false
		}

		return true
	}

	switch {
	case t.Before(end):
		return false
	case t.After(start):
		return false
	}

	return true
}

func (t *Time) String() string {
	return t.t.Format("2006-01-02 15:04 -0700 MST")
}

func (t *Time) Equals(a Time) bool {
	return t.String() == a.String()
}
