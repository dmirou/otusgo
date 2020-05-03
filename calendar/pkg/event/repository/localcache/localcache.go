package localcache

import (
	"context"
	"strconv"
	"sync"
	"time"

	errors "github.com/dmirou/otusgo/calendar/pkg/error"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/helper"
)

type TxMock struct{}

func (tm *TxMock) Commit() error {
	return nil
}

func (tm *TxMock) Rollback() error {
	return nil
}

type LocalCache struct {
	events map[string]*event.Event
	mu     *sync.Mutex
}

func New() *LocalCache {
	return &LocalCache{
		events: make(map[string]*event.Event),
		mu:     new(sync.Mutex),
	}
}

func (lc *LocalCache) Begin() (event.Transact, error) {
	return &TxMock{}, nil
}

func (lc *LocalCache) Create(ctx context.Context, e *event.Event) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	ecopy := *e
	ecopy.ID = strconv.Itoa(len(lc.events) + 1)
	lc.events[ecopy.ID] = &ecopy
	e.ID = ecopy.ID

	return nil
}

func (lc *LocalCache) GetByID(ctx context.Context, userID, id string) (*event.Event, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	e, ok := lc.events[id]
	if !ok {
		return nil, &errors.EventNotFoundError{EventID: id}
	}

	ecopy := *e

	return &ecopy, nil
}

func (lc *LocalCache) Update(ctx context.Context, e *event.Event) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	actual, ok := lc.events[e.ID]
	if !ok {
		return &errors.EventNotFoundError{EventID: e.ID}
	}

	if actual.Title != e.Title {
		actual.Title = e.Title
	}

	if actual.Desc != e.Desc {
		actual.Desc = e.Desc
	}

	if actual.Start != e.Start {
		actual.Start = e.Start
	}

	if actual.End != e.End {
		actual.End = e.End
	}

	return nil
}

func (lc *LocalCache) Delete(ctx context.Context, userID, id string) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if _, ok := lc.events[id]; !ok {
		return &errors.EventNotFoundError{EventID: id}
	}

	delete(lc.events, id)

	return nil
}

func (lc *LocalCache) FindByDate(
	ctx context.Context,
	userID string,
	date time.Time,
) ([]*event.Event, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var events = make([]*event.Event, 0)
	for _, e := range lc.events {
		if e.UserID != userID {
			continue
		}

		if !helper.HasDate(e.Start, date) {
			continue
		}

		ecopy := *e
		events = append(events, &ecopy)
	}

	return events, nil
}

func (lc *LocalCache) FindInside(
	ctx context.Context,
	userID string,
	start time.Time,
	d time.Duration,
) ([]*event.Event, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	end := start.Add(d)

	var events = make([]*event.Event, 0)
	for _, e := range lc.events {
		if e.UserID != userID {
			continue
		}

		if !helper.TimeInsideOrEqual(e.Start, start, end) {
			continue
		}

		ecopy := *e
		events = append(events, &ecopy)
	}

	return events, nil
}

func (lc *LocalCache) FindCrossing(
	ctx context.Context,
	userID string,
	start,
	end time.Time,
) ([]*event.Event, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var events = make([]*event.Event, 0)
	for _, e := range lc.events {
		if e.UserID != userID {
			continue
		}

		if helper.TimeInside(e.Start, start, end) || helper.TimeInside(e.End, start, end) {
			ecopy := *e
			events = append(events, &ecopy)
		}
	}

	return events, nil
}
