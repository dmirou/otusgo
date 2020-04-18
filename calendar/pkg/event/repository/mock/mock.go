package mock

import (
	"context"
	"time"

	"github.com/dmirou/otusgo/calendar/pkg/event"
)

type TxMock struct{}

func (tm *TxMock) Commit() error {
	return nil
}

func (tm *TxMock) Rollback() error {
	return nil
}

type Mock struct {
	BeginCalled bool
	BeginFn     func() (event.Transact, error)

	CreateCalled bool
	CreateFn     func(ctx context.Context, e *event.Event) error

	GetByIDCalled bool
	GetByIDFn     func(ctx context.Context, id event.ID) (*event.Event, error)

	UpdateCalled bool
	UpdateFn     func(ctx context.Context, e *event.Event) error

	DeleteCalled bool
	DeleteFn     func(ctx context.Context, id event.ID) error

	FindByDateCalled bool
	FindByDateFn     func(ctx context.Context, year, month, day int) ([]*event.Event, error)

	FindCrossingCalled bool
	FindCrossingFn     func(ctx context.Context, start, end time.Time) ([]*event.Event, error)
}

func New() *Mock {
	m := Mock{}
	m.BeginFn = func() (event.Transact, error) {
		return &TxMock{}, nil
	}
	m.CreateFn = func(ctx context.Context, e *event.Event) error {
		return nil
	}
	m.GetByIDFn = func(ctx context.Context, id event.ID) (*event.Event, error) {
		return nil, nil
	}
	m.UpdateFn = func(ctx context.Context, e *event.Event) error {
		return nil
	}
	m.DeleteFn = func(ctx context.Context, id event.ID) error {
		return nil
	}
	m.FindByDateFn = func(ctx context.Context, year, month, day int) ([]*event.Event, error) {
		return []*event.Event{}, nil
	}
	m.FindCrossingFn = func(ctx context.Context, start, end time.Time) ([]*event.Event, error) {
		return []*event.Event{}, nil
	}

	return &m
}

func (m *Mock) ResetCalled() {
	m.BeginCalled = false
	m.CreateCalled = false
	m.GetByIDCalled = false
	m.UpdateCalled = false
	m.DeleteCalled = false
	m.FindByDateCalled = false
	m.FindCrossingCalled = false
}

func (m *Mock) Begin() (event.Transact, error) {
	m.BeginCalled = true

	return m.BeginFn()
}

func (m *Mock) Create(ctx context.Context, e *event.Event) error {
	m.CreateCalled = true

	return m.CreateFn(ctx, e)
}

func (m *Mock) GetByID(ctx context.Context, id event.ID) (*event.Event, error) {
	m.GetByIDCalled = true

	return m.GetByIDFn(ctx, id)
}

func (m *Mock) Update(ctx context.Context, e *event.Event) error {
	m.UpdateCalled = true

	return m.UpdateFn(ctx, e)
}

func (m *Mock) Delete(ctx context.Context, id event.ID) error {
	m.DeleteCalled = true

	return m.DeleteFn(ctx, id)
}

func (m *Mock) FindByDate(ctx context.Context, userID event.UserID, year, month, day int) ([]*event.Event, error) {
	m.FindByDateCalled = true

	return m.FindByDateFn(ctx, year, month, day)
}

func (m *Mock) FindCrossing(ctx context.Context, userID event.UserID, start, end time.Time) ([]*event.Event, error) {
	m.FindCrossingCalled = true

	return m.FindCrossingFn(ctx, start, end)
}
