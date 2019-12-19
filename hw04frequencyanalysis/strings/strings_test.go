package strings

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// WordCountCase is test case for WordsCount function
type WordCountCase struct {
	In  string
	Out map[string]int
}

// Top10Case is test case for Top10 function
type Top10Case struct {
	In  string
	Out []string
}

// TestWordsCount check a calculation count of the each word in the string
func TestWordsCount(t *testing.T) {
	cases := []WordCountCase{
		{
			In:  "",
			Out: map[string]int{},
		},
		{
			In: "один три, три - два, три, два",
			Out: map[string]int{
				"один": 1,
				"три":  3,
				"два":  2,
			},
		},
		{
			In: "cat and dog one dog two cats and one man",
			Out: map[string]int{
				"cat":  1,
				"and":  2,
				"dog":  2,
				"one":  2,
				"two":  1,
				"cats": 1,
				"man":  1,
			},
		},
		{
			In: "Кошка и собака - одна собака, две кошки и тут кошка",
			Out: map[string]int{
				"Кошка":  2,
				"и":      2,
				"собака": 2,
				"одна":   1,
				"две":    1,
				"кошки":  1,
				"тут":    1,
			},
		},
		{
			In: "Нога ударила мяч. Мяч и нога не друзья, они - подруги. Подруги побили друга.",
			Out: map[string]int{
				"Нога":    2,
				"ударила": 1,
				"мяч":     2,
				"и":       1,
				"не":      1,
				"друзья":  1,
				"они":     1,
				"подруги": 2,
				"побили":  1,
				"друга":   1,
			},
		},
	}

	for _, testCase := range cases {
		assert.Equalf(t, testCase.Out, WordsCount(testCase.In),
			"out maps are different, input:\n%s\n", testCase.In)
	}
}

// TestTop10 check cases when there are less than or equal to 10 common words in the string
func TestTop10(t *testing.T) {
	cases := []Top10Case{
		{
			In:  "",
			Out: nil,
		},
		{
			In: "один три, три - два, три, два",
			Out: []string{
				"три",
			},
		},
		{
			In: "cat and dog one dog two cats and one man",
			Out: []string{
				"and",
				"dog",
				"one",
			},
		},
		{
			In: "Кошка и собака - одна собака, две кошки и тут кошка",
			Out: []string{
				"Кошка",
				"собака",
				"и",
			},
		},
		{
			In: "Нога ударила мяч. Мяч и нога не друзья, они - подруги. Подруги побили друга.",
			Out: []string{
				"Нога",
				"мяч",
				"подруги",
			},
		},
	}

	for _, testCase := range cases {
		expected := testCase.Out
		actual := Top10(testCase.In)
		sort.Strings(expected)
		sort.Strings(actual)
		assert.EqualValuesf(t, testCase.Out, actual,
			"out arrays are different, input:\n%s\n", testCase.In)
	}
}

// TestTop10MoreThan10 check cases when there are more that 10 common words in the string
func TestTop10MoreThan10(t *testing.T) {
	cases := []Top10Case{
		{
			In: "1 - 2, 3. привет 4. 5, я 6 - 7, тень 8. 9 0, сок . 1 2 - 3 4 5 6 7 8 9 0 сок",
			Out: []string{
				"1",
				"2",
				"3",
				"4",
				"5",
				"6",
				"7",
				"8",
				"9",
				"0",
				"сок",
			},
		},
		{
			In: "a1 a b2 c c3 d d4 e e5 f6 g7 h8 i9 j10 k11 a1 b2 c3 d4 e5 f6 g7 h8 i9 j10 k11",
			Out: []string{
				"a1",
				"b2",
				"c3",
				"d4",
				"e5",
				"f6",
				"g7",
				"h8",
				"i9",
				"j10",
				"k11",
			},
		},
	}
	for _, testCase := range cases {
		expected := testCase.Out
		actual := Top10(testCase.In)
		assert.Len(t, actual, 10, "actual count is invalid")
		assert.Subset(t, testCase.Out, expected, actual,
			"actual is not a subset of expected, input:\n%s\n", testCase.In)
	}
}
