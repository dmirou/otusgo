package event

import (
	"context"
	"time"
)

type UseCase interface {
	CreateEvent(ctx context.Context, e *Event) error
	GetEventByID(ctx context.Context, id ID) (*Event, error)
	UpdateEvent(ctx context.Context, e *Event) error
	DeleteEvent(ctx context.Context, id ID) error
	ListEventsPerDate(ctx context.Context, userID UserID, date time.Time) ([]*Event, error)
	ListEventsPerWeek(ctx context.Context, userID UserID, start time.Time) ([]*Event, error)
	ListEventsPerMonth(ctx context.Context, userID UserID, start time.Time) ([]*Event, error)
}
