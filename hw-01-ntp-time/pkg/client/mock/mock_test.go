package mock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	callback := func() (time.Time, error) {
		return time.Time{}, nil
	}
	timeClient := NewClient(callback)
	assert.NotNil(t, timeClient, "client should be not nil")
}

func TestGetTimeSuccess(t *testing.T) {
	expected := time.Now()
	callback := func() (time.Time, error) {
		return expected, nil
	}
	timeClient := NewClient(callback)
	actual, err := timeClient.GetTime()
	assert.Nil(t, err, "GetTime error should be nil")
	assert.EqualValuesf(t, expected, actual,
		"GetTime result is invalid, expected %s, actual %s", actual, expected)
}
