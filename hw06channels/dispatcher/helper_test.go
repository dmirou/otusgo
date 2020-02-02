package dispatcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GenerateTasksCase describes a test case for TestGenerateTasks.
type GenerateTasksCase struct {
	SuccessCount int
	ErrorCount   int
}

// TestGenerateTasks checks that generateTasks makes a slice with the specified counts of tasks.
func TestGenerateTasks(t *testing.T) {
	testCases := []GenerateTasksCase{
		{
			SuccessCount: 0,
			ErrorCount:   0,
		},
		{
			SuccessCount: 1,
			ErrorCount:   0,
		},
		{
			SuccessCount: 3,
			ErrorCount:   0,
		},
		{
			SuccessCount: 0,
			ErrorCount:   1,
		},
		{
			SuccessCount: 0,
			ErrorCount:   5,
		},
		{
			SuccessCount: 1,
			ErrorCount:   1,
		},
		{
			SuccessCount: 3,
			ErrorCount:   5,
		},
	}

	for _, testCase := range testCases {
		tasks := generateTasks(testCase.SuccessCount, testCase.ErrorCount)
		actualSuccessCount := 0
		actualErrorCount := 0
		for _, task := range tasks {
			if task() != nil {
				actualErrorCount++
				continue
			}
			actualSuccessCount++
		}
		assert.Equalf(t, testCase.SuccessCount, actualSuccessCount, "Success counts are different")
		assert.Equalf(t, testCase.ErrorCount, actualErrorCount, "Error counts are different")
	}

}
