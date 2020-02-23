package time

import (
	"testing"
)

func TestHasDate(t *testing.T) {
	testdata := []struct {
		date  *Time
		year  int
		month int
		day   int
		yes   bool
	}{
		{
			New(2020, 1, 15, 0, 0),
			2020,
			1,
			15,
			true,
		},
		{
			New(2021, 2, 16, 23, 59),
			2021,
			2,
			16,
			true,
		},
		{
			New(2022, 12, 1, 24, 0),
			2022,
			12,
			1,
			false,
		},
		{
			New(2023, 11, 25, 0, 0),
			2023,
			11,
			26,
			false,
		},
		{
			New(2024, 10, 15, 23, 59),
			2024,
			10,
			16,
			false,
		},
	}

	for _, td := range testdata {
		actual := td.date.HasDate(td.year, td.month, td.day)
		if actual != td.yes {
			t.Errorf(
				"unexpected result from HasDate method: %t, expected: %t,\ndate: %v, Y-m-d: %d-%d-%d",
				actual, td.yes, td.date.t, td.year, td.month, td.day,
			)
		}
	}
}

// nolint: funlen
func TestInside(t *testing.T) {
	testData := []struct {
		t     *Time
		start *Time
		end   *Time
		yes   bool
	}{
		{
			New(2020, 2, 3, 10, 0),
			New(2020, 2, 3, 10, 0),
			New(2020, 2, 3, 10, 0),
			false,
		},
		{
			New(2021, 2, 3, 10, 0),
			New(2021, 2, 3, 10, 0),
			New(2021, 2, 3, 10, 1),
			false,
		},
		{
			New(2022, 2, 3, 10, 0),
			New(2022, 2, 3, 10, 1),
			New(2022, 2, 3, 10, 1),
			false,
		},
		{
			New(2023, 2, 3, 9, 59),
			New(2023, 2, 3, 10, 0),
			New(2023, 2, 3, 10, 1),
			false,
		},
		{
			New(2023, 2, 3, 9, 59),
			New(2023, 2, 3, 10, 1),
			New(2023, 2, 3, 10, 0),
			false,
		},
		{
			New(2024, 2, 3, 10, 2),
			New(2024, 2, 3, 10, 0),
			New(2024, 2, 3, 10, 1),
			false,
		},
		{
			New(2025, 2, 3, 10, 2),
			New(2025, 2, 3, 10, 1),
			New(2025, 2, 3, 10, 0),
			false,
		},
		{
			New(2030, 2, 3, 10, 1),
			New(2030, 2, 3, 9, 59),
			New(2030, 2, 3, 10, 2),
			true,
		},
		{
			New(2031, 2, 3, 10, 1),
			New(2031, 2, 3, 10, 2),
			New(2031, 2, 3, 9, 59),
			true,
		},
	}

	for _, td := range testData {
		actual := td.t.Inside(*td.start, *td.end)
		if actual != td.yes {
			t.Errorf(
				"unexpected result from Inside method: %t, expected: %t,\ndate: %s,\nstart: %s,\nend %s",
				actual, td.yes, td.t, td.start, td.end,
			)
		}
	}
}
