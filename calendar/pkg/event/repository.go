package event

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, event *Event) error
	GetByID(ctx context.Context, id ID) (*Event, error)
	Update(ctx context.Context, event *Event) error
	Delete(ctx context.Context, id ID) error
	Find(ctx context.Context) ([]*Event, error)
}
