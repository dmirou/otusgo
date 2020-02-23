package event

import (
	"fmt"
)

type NotFoundError struct {
	EventID ID
	Err     error
}

func (nfe *NotFoundError) Error() string {
	return fmt.Sprintf("event %s not found: %v", nfe.EventID, nfe.Err)
}
