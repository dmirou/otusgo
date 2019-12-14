package client

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/dmirou/otusgo/hw01ntptime/pkg/client/transport/mock"
)

func TestGetTimeSuccess(t *testing.T) {
	expectedResult := time.Now()
	var expectedErr error = nil
	mockTransport := mock.NewTransport(expectedResult, expectedErr)
	client := NewClient(mockTransport)
	result, err := client.GetTime()
	assert.EqualValuesf(t, expectedResult, result,
		"GetTime results are different\nexpected: %s, actual: %s", expectedResult, result)
	assert.EqualValuesf(t, expectedErr, err,
		"GetTime errors are different\nexpected: %s, actual %s", expectedErr, err)
}

func TestGetTimeWithError(t *testing.T) {
	expectedResult := time.Time{}
	expectedErr := errors.New("test error")
	mockTransport := mock.NewTransport(expectedResult, expectedErr)
	client := NewClient(mockTransport)
	result, err := client.GetTime()
	assert.EqualValuesf(t, expectedResult, result,
		"GetTime results are different\nexpected: %s, actual: %s", expectedResult, result)
	assert.EqualValuesf(t, expectedErr, err,
		"GetTime errors are different\nexpected: %s, actual %s", expectedErr, err)
}
