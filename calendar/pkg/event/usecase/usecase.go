package usecase

import (
	"context"

	errors "github.com/dmirou/otusgo/calendar/pkg/error"
	"github.com/dmirou/otusgo/calendar/pkg/event"
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
			Name:   "e",
			Method: "CreateEvent",
			Desc:   "event should be not nil",
		}
	}

	if e.Title == "" {
		return &errors.InvalidArgError{
			Name:   "e",
			Method: "CreateEvent",
			Desc:   "event title should be not empty",
		}
	}

	if e.Start == nil {
		return &errors.InvalidArgError{
			Name:   "e",
			Method: "CreateEvent",
			Desc:   "event start date should be not empty",
		}
	}

	if e.End == nil {
		return &errors.InvalidArgError{
			Name:   "e",
			Method: "CreateEvent",
			Desc:   "event end date should be not empty",
		}
	}

	if e.Start.After(*e.End) {
		return &errors.InvalidArgError{
			Name:   "e",
			Method: "CreateEvent",
			Desc:   "event end date should be greater than start date",
		}
	}

	if !e.Start.HasDate(e.End.Year(), e.End.Month(), e.End.Day()) {
		return &errors.InvalidArgError{
			Name:   "e",
			Method: "CreateEvent",
			Desc:   "event end date should have the same date as event start date",
		}
	}

	events, err := uc.repo.FindCrossing(ctx, *e.Start, *e.End)
	if err != nil {
		return err
	}

	if len(events) != 0 {
		return &errors.DateBusyError{Start: e.Start, End: e.End}
	}

	return nil
}

func (uc *UseCase) GetEventByID(ctx context.Context, id event.ID) (*event.Event, error) {
	return uc.repo.GetByID(ctx, id)
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

func (uc *UseCase) DeleteEvent(ctx context.Context, id event.ID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UseCase) ListEventsByDate(ctx context.Context, year, month, day int) ([]*event.Event, error) {
	return uc.repo.FindByDate(ctx, year, month, day)
}