package list

import (
	"errors"
	"math/rand"
	"time"
)

// GenerateSliceWithLength generates a new random slice with the length between min and max
func GenerateSliceWithLength(min, max int) ([]int, error) {
	length, err := GenerateInt(min, max)
	if err != nil {
		return nil, err
	}
	return GenerateSlice(length)
}

// GenerateSlice creates a new slice with n random integers
func GenerateSlice(n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("slice length should be greater or equal than one")
	}
	result := make([]int, n)
	for i := 0; i < n-1; i++ {
		result[i] = rand.Intn(n)
	}
	return result, nil
}

// GenerateInt returns a random integer between min and max [min, max]
// min and max are included to the possible result
func GenerateInt(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("min should be bigger or equal than max")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min, nil
}
