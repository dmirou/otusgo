package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GenerateIntCase describes a test case for the GenerateInt method
type GenerateIntCase struct {
	Min            int
	Max            int
	PossibleValues []int
}

// TestGenerateInt checks that a generated number is in a correct range
func TestGenerateInt(t *testing.T) {
	var testCases = []GenerateIntCase{
		{
			Min:            0,
			Max:            0,
			PossibleValues: []int{0},
		},
		{
			Min:            0,
			Max:            1,
			PossibleValues: []int{0, 1},
		},
		{
			Min:            1,
			Max:            3,
			PossibleValues: []int{1, 2, 3},
		},
	}
	for _, testCase := range testCases {
		actual, err := GenerateInt(testCase.Min, testCase.Max)
		assert.Containsf(t, testCase.PossibleValues, actual, "Value is out of range")
		assert.Nilf(t, err, "Error should be nil")
	}
}

// TestGenerateSlice checks that the method GenerateSlice generates slice with a specified length
func TestGenerateSlice(t *testing.T) {
	var lengths = []int{1, 2, 5, 10}
	for _, length := range lengths {
		result, err := GenerateSlice(length)
		assert.Lenf(t, result, length, "Slice length is invalid")
		assert.Nilf(t, err, "Error should be nil: expected length %d", length)
	}
}
