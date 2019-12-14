package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case defines a test case
type Case struct {
	In  string
	Out CaseOut
}

// CaseOut defines result returned by a function to be tested
type CaseOut struct {
	Str string
	Err error
}

func TestUnpack(t *testing.T) {
	cases := []Case{
		{
			In:  "a4bc2d5e",
			Out: CaseOut{Str: "aaaabccddddde", Err: nil},
		},
		{
			In:  "abcd",
			Out: CaseOut{Str: "abcd", Err: nil},
		},
		{
			In:  "45",
			Out: CaseOut{Str: "", Err: &InvalidArgError{Value: "45"}},
		},
		{
			In:  "",
			Out: CaseOut{Str: "", Err: nil},
		},
	}

	for _, testCase := range cases {
		outStr, err := Unpack(testCase.In)
		assert.EqualValuesf(t, testCase.Out.Str, outStr,
			"input: %s, out strings are different", testCase.In)
		assert.EqualValues(t, testCase.Out.Err, err,
			"input: %s, out errors are different", testCase.In)
	}
}
