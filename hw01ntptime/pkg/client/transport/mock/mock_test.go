package mock

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeSuccess(t *testing.T) {
	expectedResult := time.Now()
	var expectedErr error = nil
	mockTransport := NewTransport(expectedResult, expectedErr)
	result, err := mockTransport.GetTime()
	assert.EqualValuesf(t, expectedResult, result,
		"GetTime results are different\nexpected: %s, actual: %s", expectedResult, result)
	assert.EqualValuesf(t, expectedErr, err,
		"GetTime errors are different\nexpected: %s, actual %s", expectedErr, err)
}

func TestGetTimeWithError(t *testing.T) {
	expectedResult := time.Time{}
	expectedErr := errors.New("test error")
	mockTransport := NewTransport(expectedResult, expectedErr)
	result, err := mockTransport.GetTime()
	assert.EqualValuesf(t, expectedResult, result,
		"GetTime results are different\nexpected: %s, actual: %s", expectedResult, result)
	assert.EqualValuesf(t, expectedErr, err,
		"GetTime errors are different\nexpected: %s, actual %s", expectedErr, err)
}
