package dispatcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Tasks       []Task // tasks to execute
	N           int    // count of workers
	M           int    // max count of errors
	ReturnError bool
}

func TestRun(t *testing.T) {
	testCases := []TestCase{
		{
			Tasks:       GenerateTasks(1, 0),
			N:           1,
			M:           0,
			ReturnError: false,
		},
		{
			Tasks:       GenerateTasks(2, 0),
			N:           1,
			M:           1,
			ReturnError: false,
		},
		{
			Tasks:       GenerateTasks(2, 1),
			N:           1,
			M:           1,
			ReturnError: true,
		},
		{
			Tasks:       GenerateTasks(2, 2),
			N:           1,
			M:           1,
			ReturnError: true,
		},
		{
			Tasks:       GenerateTasks(10, 3),
			N:           2,
			M:           2,
			ReturnError: true,
		},
	}

	for _, testCase := range testCases {
		err := Run(testCase.Tasks, testCase.N, testCase.M)
		if testCase.ReturnError {
			assert.Errorf(t, err, "result should be an error")
		} else {
			assert.Nilf(t, err, "result should be nil")
		}
	}
}
