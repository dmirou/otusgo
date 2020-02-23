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

func (d *Time) After(u Time) bool {
	return d.t.After(u.t)
}

func (d *Time) Before(u Time) bool {
	return d.t.Before(u.t)
}

func (d *Time) HasDate(year, month, day int) bool {
	start := New(year, month, day-1, 23, 59)
	end := New(year, month, day+1, 0, 0)

	return d.After(*start) && d.Before(*end)
}
