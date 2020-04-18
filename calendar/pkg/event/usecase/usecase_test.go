package usecase

import (
	"context"
	goerrors "errors"
	"testing"
	"time"

	errors "github.com/dmirou/otusgo/calendar/pkg/error"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/event/repository/mock"
	"github.com/dmirou/otusgo/calendar/pkg/helper"
)

// nolint: funlen
func TestCreateEvent(t *testing.T) {
	repo := mock.New()
	uc := New(repo)

	e := &event.Event{
		Title: "Breakfast",
		Start: helper.NewTime(2020, 2, 29, 8, 30),
		End:   helper.NewTime(2020, 2, 29, 8, 45),
	}

	if err := uc.CreateEvent(context.Background(), e); err != nil {
		t.Errorf("unexpected error in CreateEvent: %v", err)
	}

	if !repo.CreateCalled {
		t.Errorf("CreateEvent repository not called but should")
	}

	expected := &errors.InvalidArgError{Name: "e", Method: "CreateEvent"}
	if err := uc.CreateEvent(context.Background(), nil); !goerrors.As(err, &expected) {
		t.Errorf("unexpected error in CreateEvent: %v", err)
	}

	eventWoTitle := *e
	eventWoTitle.Title = ""

	if err := uc.CreateEvent(context.Background(), &eventWoTitle); !goerrors.As(err, &expected) {
		t.Errorf("unexpected error in CreateEvent: %v", err)
	}

	startBeforeEnd := *e
	startBeforeEnd.Start, startBeforeEnd.End = startBeforeEnd.End, startBeforeEnd.Start

	repo.FindCrossingFn = func(ctx context.Context, start, end time.Time) ([]*event.Event, error) {
		return []*event.Event{e}, nil
	}
	repo.ResetCalled()

	errdb := &errors.DateBusyError{}
	if err := uc.CreateEvent(context.Background(), e); !goerrors.As(err, &errdb) {
		t.Errorf("unexpected error in CreateEvent: %v", err)
	}

	if !repo.FindCrossingCalled {
		t.Errorf("FindCrossing of repository not called but should")
	}

	if repo.CreateCalled {
		t.Errorf("CreateEvent of repository called but should not")
	}
}
