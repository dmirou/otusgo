package localcache

import (
	"context"
	"sync"

	"github.com/dmirou/otusgo/calendar/pkg/event"
)

type LocalCache struct {
	events map[event.ID]*event.Event
	mu     *sync.Mutex
}

func New() *LocalCache {
	return &LocalCache{
		events: make(map[event.ID]*event.Event),
		mu:     new(sync.Mutex),
	}
}

func (lc *LocalCache) Create(ctx context.Context, e *event.Event) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.events[e.ID] = e

	return nil
}

func (lc *LocalCache) GetByID(ctx context.Context, id event.ID) (*event.Event, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	e, ok := lc.events[id]
	if !ok {
		return nil, &event.NotFoundError{EventID: id}
	}

	return e, nil
}

func (lc *LocalCache) Update(ctx context.Context, e *event.Event) error {
	return nil
}

func (lc *LocalCache) Delete(ctx context.Context, id event.ID) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if _, ok := lc.events[id]; !ok {
		return &event.NotFoundError{EventID: id}
	}

	delete(lc.events, id)

	return nil
}

func (lc *LocalCache) FindByDate(ctx context.Context, year, month, day int) ([]*event.Event, error) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	var events = make([]*event.Event, 0)
	for _, e := range lc.events {
		if !e.Start.HasDate(year, month, day) {
			continue
		}

		events = append(events, e)
	}

	return events, nil
}
