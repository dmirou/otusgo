package event

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, e *Event) error
	GetByID(ctx context.Context, id ID) (*Event, error)
	Update(ctx context.Context, e *Event) error
	Delete(ctx context.Context, id ID) error
	FindByDate(ctx context.Context, year, month, day int) ([]*Event, error)
}