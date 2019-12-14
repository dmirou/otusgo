package strings

import (
	"fmt"
)

// InvalidArgError is a invalid argument error
type InvalidArgError struct {
	Value string
}

// Error converts error to string
func (iae InvalidArgError) Error() string {
	return fmt.Sprintf("Invalid argument received: %s", iae.Value)
}
