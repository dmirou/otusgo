package event

import (
	"context"
	"time"
)

type UseCase interface {
	CreateEvent(ctx context.Context, e *Event) error
	GetEventByID(ctx context.Context, userID, id string) (*Event, error)
	UpdateEvent(ctx context.Context, e *Event) error
	DeleteEvent(ctx context.Context, userID, id string) error
	ListEventsPerDate(ctx context.Context, userID string, date time.Time) ([]*Event, error)
	ListEventsPerWeek(ctx context.Context, userID string, start time.Time) ([]*Event, error)
	ListEventsPerMonth(ctx context.Context, userID string, start time.Time) ([]*Event, error)
}
