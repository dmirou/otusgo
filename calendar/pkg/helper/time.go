package helper

import "time"

const (
	Day   = time.Hour * 24
	Week  = Day * 7
	Month = Week * 4
)

// NewDate creates new date in UTC location
func NewDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// NewTime creates new time in UTC location
func NewTime(year int, month time.Month, day, hour, min int) time.Time {
	return time.Date(year, month, day, hour, min, 0, 0, time.UTC)
}

// HasDate checks if t has specified year, month and day from date
func HasDate(t, date time.Time) bool {
	return t.Year() == date.Year() &&
		t.Month() == date.Month() &&
		t.Day() == date.Day()
}

// TimeInside checks that t inside (start, end), not including borders.
func TimeInside(t, start, end time.Time) bool {
	return timeInside(t, start, end, false)
}

// TimeInsideOrEqual checks that t inside [start, end], including borders.
func TimeInsideOrEqual(t, start, end time.Time) bool {
	return timeInside(t, start, end, true)
}

// timeInside checks that t inside (start, end), not including borders,
// or t inside [start, end], if orEqual is true.
func timeInside(t, start, end time.Time, orEqual bool) bool {
	switch {
	case orEqual && (t == start || t == end):
		return true
	case !orEqual && (start == end || t == start || t == end):
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
