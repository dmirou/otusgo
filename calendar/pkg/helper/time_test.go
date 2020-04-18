package helper

import (
	"testing"
	"time"
)

func TestHasDate(t *testing.T) {
	testdata := []struct {
		date  time.Time
		year  int
		month int
		day   int
		yes   bool
	}{
		{
			NewTime(2020, 1, 15, 0, 0),
			2020,
			1,
			15,
			true,
		},
		{
			NewTime(2021, 2, 16, 23, 59),
			2021,
			2,
			16,
			true,
		},
		{
			NewTime(2022, 12, 1, 24, 0),
			2022,
			12,
			1,
			false,
		},
		{
			NewTime(2023, 11, 25, 0, 0),
			2023,
			11,
			26,
			false,
		},
		{
			NewTime(2024, 10, 15, 23, 59),
			2024,
			10,
			16,
			false,
		},
	}

	for _, td := range testdata {
		actual := HasDate(td.date, td.year, td.month, td.day)
		if actual != td.yes {
			t.Errorf(
				"unexpected result from HasDate method: %t, expected: %t,\ndate: %v, Y-m-d: %d-%d-%d",
				actual, td.yes, td.date, td.year, td.month, td.day,
			)
		}
	}
}

// nolint: funlen
func TestInside(t *testing.T) {
	testData := []struct {
		t     time.Time
		start time.Time
		end   time.Time
		yes   bool
	}{
		{
			NewTime(2020, 2, 3, 10, 0),
			NewTime(2020, 2, 3, 10, 0),
			NewTime(2020, 2, 3, 10, 0),
			false,
		},
		{
			NewTime(2021, 2, 3, 10, 0),
			NewTime(2021, 2, 3, 10, 0),
			NewTime(2021, 2, 3, 10, 1),
			false,
		},
		{
			NewTime(2022, 2, 3, 10, 0),
			NewTime(2022, 2, 3, 10, 1),
			NewTime(2022, 2, 3, 10, 1),
			false,
		},
		{
			NewTime(2023, 2, 3, 9, 59),
			NewTime(2023, 2, 3, 10, 0),
			NewTime(2023, 2, 3, 10, 1),
			false,
		},
		{
			NewTime(2023, 2, 3, 9, 59),
			NewTime(2023, 2, 3, 10, 1),
			NewTime(2023, 2, 3, 10, 0),
			false,
		},
		{
			NewTime(2024, 2, 3, 10, 2),
			NewTime(2024, 2, 3, 10, 0),
			NewTime(2024, 2, 3, 10, 1),
			false,
		},
		{
			NewTime(2025, 2, 3, 10, 2),
			NewTime(2025, 2, 3, 10, 1),
			NewTime(2025, 2, 3, 10, 0),
			false,
		},
		{
			NewTime(2030, 2, 3, 10, 1),
			NewTime(2030, 2, 3, 9, 59),
			NewTime(2030, 2, 3, 10, 2),
			true,
		},
		{
			NewTime(2031, 2, 3, 10, 1),
			NewTime(2031, 2, 3, 10, 2),
			NewTime(2031, 2, 3, 9, 59),
			true,
		},
	}

	for _, td := range testData {
		actual := TimeInside(td.t, td.start, td.end)
		if actual != td.yes {
			t.Errorf(
				"unexpected result from TimeInside method: %t, expected: %t,\ndate: %s,\nstart: %s,\nend %s",
				actual, td.yes, td.t, td.start, td.end,
			)
		}
	}
}
