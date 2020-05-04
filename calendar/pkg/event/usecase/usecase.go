package usecase

import (
	"context"
	"time"

	errors "github.com/dmirou/otusgo/calendar/pkg/error"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/helper"
)

type UseCase struct {
	repo event.Repository
}

func New(repo event.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) CreateEvent(ctx context.Context, e *event.Event) error {
	tx, err := uc.repo.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback() // nolint: errcheck

	if err := uc.validateEvent(ctx, e); err != nil {
		return err
	}

	if err := uc.repo.Create(ctx, e); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) validateEvent(ctx context.Context, e *event.Event) error {
	if e == nil {
		return &errors.InvalidArgError{
			Name:   "event",
			Method: "CreateEvent",
			Desc:   "event should be not nil",
		}
	}

	if e.UserID == "" {
		return &errors.InvalidArgError{
			Name:   "user ID",
			Method: "CreateEvent",
			Desc:   "event userID should be not empty",
		}
	}

	if e.Title == "" {
		return &errors.InvalidArgError{
			Name:   "title",
			Method: "CreateEvent",
			Desc:   "event title should be not empty",
		}
	}

	if e.Start.After(e.End) || e.End.Sub(e.Start) < 15*time.Minute {
		return &errors.InvalidArgError{
			Name:   "end date",
			Method: "CreateEvent",
			Desc:   "event end date should be greater than start date more than 15 minutes",
		}
	}

	if !helper.HasDate(e.Start, e.End) {
		return &errors.InvalidArgError{
			Name:   "end date",
			Method: "CreateEvent",
			Desc:   "event end date should have the same date as event start date",
		}
	}

	events, err := uc.repo.FindCrossing(ctx, e.UserID, e.Start, e.End)
	if err != nil {
		return err
	}

	if len(events) != 0 {
		return &errors.DateBusyError{Start: e.Start, End: e.End}
	}

	return nil
}

func (uc *UseCase) GetEventByID(ctx context.Context, userID, id string) (
	*event.Event, error,
) {
	return uc.repo.GetByID(ctx, userID, id)
}

func (uc *UseCase) UpdateEvent(ctx context.Context, e *event.Event) error {
	tx, err := uc.repo.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback() // nolint: errcheck

	if err := uc.validateEvent(ctx, e); err != nil {
		return err
	}

	if err := uc.repo.Update(ctx, e); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteEvent(ctx context.Context, userID, id string) error {
	return uc.repo.Delete(ctx, userID, id)
}

func (uc *UseCase) ListEventsPerDate(
	ctx context.Context,
	userID string,
	date time.Time,
) ([]*event.Event, error) {
	return uc.repo.FindByDate(ctx, userID, date)
}

func (uc *UseCase) ListEventsPerWeek(
	ctx context.Context,
	userID string,
	start time.Time,
) ([]*event.Event, error) {
	return uc.repo.FindInside(ctx, userID, start, helper.Week)
}

func (uc *UseCase) ListEventsPerMonth(
	ctx context.Context,
	userID string,
	start time.Time,
) ([]*event.Event, error) {
	return uc.repo.FindInside(ctx, userID, start, helper.Month)
}
