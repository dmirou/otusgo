package event

import (
	"context"

	"github.com/dmirou/otusgo/calendar/pkg/time"
)

type Repository interface {
	Begin() (Transact, error)
	Create(ctx context.Context, e *Event) error
	GetByID(ctx context.Context, id ID) (*Event, error)
	Update(ctx context.Context, e *Event) error
	Delete(ctx context.Context, id ID) error
	FindByDate(ctx context.Context, year, month, day int) ([]*Event, error)
	FindCrossing(ctx context.Context, start, end time.Time) ([]*Event, error)
}

type Transact interface {
	Commit() error
	Rollback() error
}
