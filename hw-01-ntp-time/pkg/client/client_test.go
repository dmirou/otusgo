package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	host := "example.com"
	timeClient := NewClient(host)
	assert.NotNil(t, timeClient, "client should be not nil")
}
