package localcache

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/time"
)

var testData = map[event.ID]*event.Event{
	"1": {
		ID:    "1",
		Title: "Daily meeting",
		Start: time.New(2020, 2, 15, 5, 00),
		End:   time.New(2020, 2, 15, 5, 30),
	},
	"2": {
		ID:    "2",
		Title: "Lunch",
		Start: time.New(2020, 2, 15, 7, 00),
		End:   time.New(2020, 2, 15, 8, 00),
	},
	"3": {
		ID:    "3",
		Title: "Running",
		Start: time.New(2020, 2, 18, 0, 00),
		End:   time.New(2020, 2, 18, 0, 15),
	},
	"4": {
		ID:    "4",
		Title: "OTUS webinar",
		Start: time.New(2020, 2, 18, 15, 00),
		End:   time.New(2020, 2, 18, 18, 00),
	},
	"5": {
		ID:    "5",
		Title: "Learning English",
		Start: time.New(2020, 2, 19, 2, 00),
		End:   time.New(2020, 2, 19, 2, 30),
	},
}

var testDataByDate = []struct {
	year  int
	month int
	day   int
	ids   []event.ID
}{
	{
		2020,
		2,
		15,
		[]event.ID{"1", "2"},
	},
	{
		2020,
		2,
		18,
		[]event.ID{"3", "4"},
	},
	{
		2020,
		2,
		19,
		[]event.ID{"5"},
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

	if err := fill(lc); err != nil {
		t.Fatalf("unexpected result in fill method: %v", err)
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

func fill(lc *LocalCache) error {
	for _, e := range testData {
		if err := lc.Create(context.Background(), e); err != nil {
			return err
		}
	}

	return nil
}

func TestDelete(t *testing.T) {
	lc := New()

	if err := fill(lc); err != nil {
		t.Fatalf("unexpected result in fill method: %v", err)
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

func TestFindByDate(t *testing.T) {
	lc := New()

	if err := fill(lc); err != nil {
		t.Fatalf("unexpected result in fill method: %v", err)
	}

	for _, td := range testDataByDate {
		date := fmt.Sprintf("%d-%d-%d", td.year, td.month, td.day)

		evs, err := lc.FindByDate(context.Background(), td.year, td.month, td.day)
		if err != nil {
			t.Errorf("unexpected result in FindByDate method: %v, date: %s", err, date)
		}

		if len(evs) != len(td.ids) {
			t.Errorf("unexpected count in FindByDate method: %d, expected: %d", len(evs), len(td.ids))
		}

		for _, id := range td.ids {
			found := false

			for _, ev := range evs {
				if ev.ID == id {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("event %s not found in FindByDate method, but should", id)
			}
		}
	}

	free := fmt.Sprintf("%d-%d-%d", 2000, 1, 2)

	evs, err := lc.FindByDate(context.Background(), 2000, 1, 2)
	if err != nil {
		t.Errorf("unexpected result in FindByDate method: %v, date: %s", err, free)
	}

	if len(evs) != 0 {
		t.Errorf("unexpected count in FindByDate method: %d, expected: %d", len(evs), 0)
	}
}