package error

import (
	"fmt"
)

type InvalidArgError struct {
	Name   string
	Method string
	Desc   string
}

func (iae *InvalidArgError) Error() string {
	return fmt.Sprintf("invalid argument %q in method %q: %s", iae.Name, iae.Method, iae.Desc)
}
