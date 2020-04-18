package event

import (
	"context"
	"time"
)

type Repository interface {
	Begin() (Transact, error)
	Create(ctx context.Context, e *Event) error
	GetByID(ctx context.Context, id ID) (*Event, error)
	Update(ctx context.Context, e *Event) error
	Delete(ctx context.Context, id ID) error
	FindByDate(ctx context.Context, userID UserID, year, month, day int) ([]*Event, error)
	FindCrossing(ctx context.Context, userID UserID, start, end time.Time) ([]*Event, error)
}

type Transact interface {
	Commit() error
	Rollback() error
}
