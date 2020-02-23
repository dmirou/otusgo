package event

import (
	"fmt"
)

type NotFoundError struct {
	EventID ID
}

func (nfe *NotFoundError) Error() string {
	return fmt.Sprintf("event not found by id: %q", nfe.EventID)
}
