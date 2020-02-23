package localcache

import (
	"context"
	"errors"
	"testing"

	"github.com/dmirou/otusgo/calendar/pkg/event"
)

var testData = map[event.ID]*event.Event{
	"1": {
		ID:    "1",
		Title: "Daily meeting",
	},
	"2": {
		ID:    "2",
		Title: "Lunch",
	},
	"3": {
		ID:    "3",
		Title: "OTUS webinar",
	},
}

func TestCreate(t *testing.T) {
	lc := New()
	e := testData["1"]

	if err := lc.Create(context.Background(), e); err != nil {
		t.Errorf("unexpected result from Create method: %v, event: %v", err, e)
	}

	if len(lc.events) != 1 {
		t.Errorf("unexpected length after Create method: %d, expected: %d", len(lc.events), 1)
	}

	if lc.events[e.ID] != e {
		t.Errorf("unexpected event created: %v, expected: %v", lc.events[e.ID], e)
	}

	e2 := testData["2"]

	if err := lc.Create(context.Background(), e2); err != nil {
		t.Errorf("unexpected result from Create method: %v, event: %v", err, e2)
	}

	if len(lc.events) != 2 {
		t.Errorf("unexpected length after Create method: %d, expected: %d", len(lc.events), 2)
	}

	if lc.events[e2.ID] != e2 {
		t.Errorf("unexpected event created: %v, expected: %v", lc.events[e2.ID], e2)
	}

	if lc.events[e.ID] != e {
		t.Errorf("first event not found: %v, expected: %v", lc.events[e.ID], e)
	}
}

func TestGetByID(t *testing.T) {
	lc := New()

	for _, e := range testData {
		if err := lc.Create(context.Background(), e); err != nil {
			t.Errorf("unexpected result from Create method: %v", err)
			continue
		}
	}

	for id := range testData {
		e, err := lc.GetByID(context.Background(), id)
		if err != nil {
			t.Errorf("unexpected result from GetByID method: %v", err)
			continue
		}

		if id != e.ID {
			t.Fatalf("unexpected event ID received from GetByID method: %s, expected: %s", e.ID, id)
			continue
		}
	}

	var nonexistent event.ID = "nonexistent id"

	var expected = &event.NotFoundError{EventID: nonexistent}

	e, err := lc.GetByID(context.Background(), nonexistent)
	if !errors.As(err, &expected) {
		t.Errorf("unexpected error returned from GetByID: %v, expected %v", err, expected)
	}

	if e != nil {
		t.Errorf("unexpected event returned from GetByID: %v, expected %v", e, nil)
	}
}

func TestDelete(t *testing.T) {
	lc := New()

	for _, e := range testData {
		if err := lc.Create(context.Background(), e); err != nil {
			t.Errorf("unexpected result from Create method: %v", err)
			continue
		}
	}

	var count int

	for id := range testData {
		count = len(lc.events)

		if err := lc.Delete(context.Background(), id); err != nil {
			t.Errorf("unexpected result from Delete method: %v", err)
			continue
		}

		if _, ok := lc.events[id]; ok {
			t.Errorf("event %s not deleted in Delete method", id)
			continue
		}

		if len(lc.events) != count-1 {
			t.Errorf("unexpected count after Delete method: %d, expected: %d", len(lc.events), count-1)
			continue
		}
	}

	var (
		nonexistent event.ID = "nonexistent id"
		expected             = &event.NotFoundError{EventID: nonexistent}
	)

	err := lc.Delete(context.Background(), nonexistent)
	if !errors.As(err, &expected) {
		t.Errorf("unexpected error returned from Delete: %v, expected %v", err, expected)
	}
}
