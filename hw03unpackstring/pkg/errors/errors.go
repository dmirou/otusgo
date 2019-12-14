package errors

import "fmt"

type InvalidArgError struct {
	Value string
}

func (iae InvalidArgError) Error() string {
	return fmt.Sprintf("Invalid argument received: %s", iae.Value)
}
