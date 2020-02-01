package dispatcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Tasks  []Task // tasks to execute
	N      int    // count of workers
	M      int    // max count of errors
	Result error
}

func TestRun(t *testing.T) {
	testCases := []TestCase{
		{
			Tasks:  []Task{},
			N:      5,
			M:      0,
			Result: nil,
		},
	}

	for _, testCase := range testCases {
		err := Run(testCase.Tasks, testCase.N, testCase.M)
		assert.Equalf(t, testCase.Result, err, "results not equal")
	}
}
