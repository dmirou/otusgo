package helper

import "time"

// NewTime creates new date in UTC location
func NewTime(year, month, day, hour, min int) time.Time {
	return time.Date(year, time.Month(month), day, hour, min, 0, 0, time.UTC)
}

// HasDate checks if t has specified year, month and day
func HasDate(t time.Time, year, month, day int) bool {
	start := NewTime(year, month, day-1, 23, 59)
	end := NewTime(year, month, day+1, 0, 0)

	return t.After(start) && t.Before(end)
}

// TimeInside checks that t inside (start, end), not including borders.
func TimeInside(t, start, end time.Time) bool {
	if start == end || t == start || t == end {
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
