// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	"github.com/dmirou/otusgo/calendar/pkg/contracts/event"
	"github.com/dmirou/otusgo/calendar/pkg/contracts/request"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	defaultUserID   = "1"
	secondsInMinute = 60
	secondsInHour   = secondsInMinute * 60
	secondsInDay    = secondsInHour * 24
	secondsInWeek   = secondsInDay * 7
)

func setup(t *testing.T) (*grpc.ClientConn, event.EventServiceClient, context.Context) {
	addr := os.Getenv("CORE_SERVER_ADDR")
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Errorf("can not connect to grpc server: %v", err)
	}

	c := event.NewEventServiceClient(cc)

	md := metadata.New(map[string]string{
		"user-id": defaultUserID,
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return cc, c, ctx
}

func TestGetNotExistingEvent(t *testing.T) {
	cc, c, ctx := setup(t)
	defer cc.Close()

	_, err := c.GetEventByID(ctx, &request.ByID{Id: "not-existing-event-id"})

	st, ok := status.FromError(err)
	if !ok {
		t.Errorf("unexpected error in GetEventByID: %v", err)
	}

	if st.Code() != codes.NotFound {
		t.Errorf(
			"unexpected error code in GetEventByID: %v, expected: %v",
			st.Code(), codes.NotFound,
		)
	}
}

func TestCreateEvent(t *testing.T) {
	cc, c, ctx := setup(t)
	defer cc.Close()

	start := nowDate()
	end := &timestamp.Timestamp{
		Seconds: start.Seconds + 15*secondsInMinute,
		Nanos:   start.Nanos,
	}

	e := &event.Event{
		Title:        "Walking",
		Desc:         "At the park",
		Start:        start,
		End:          end,
		NotifyBefore: &duration.Duration{Seconds: 15 * secondsInMinute},
	}

	eupd, err := c.CreateEvent(ctx, e)
	if err != nil {
		t.Fatalf("unexpected error in CreateEvent: %v", err)
	}

	if eupd.Id == "" {
		t.Errorf("got empty event id")
	}

	e2, err := c.GetEventByID(ctx, &request.ByID{Id: eupd.Id})
	if err != nil {
		t.Fatalf("unexpected error in GetEventByID: %v", err)
	}

	diff := cmp.Diff(eupd, e2, cmp.Exporter(exportAll))
	if diff != "" {
		t.Fatalf("unexpected event in GetEventByID: %v", diff)
	}

	_, err = c.CreateEvent(ctx, e)
	if err == nil {
		t.Fatalf("got: nil, but expected: error with invalid arg 'start date'")
	}

	checkFieldViolation(t, err, "start date")
}

func TestCreateEventShortDuration(t *testing.T) {
	cc, c, ctx := setup(t)
	defer cc.Close()

	now := nowDate()
	start := &timestamp.Timestamp{
		Seconds: now.Seconds + secondsInHour,
		Nanos:   now.Nanos,
	}
	end := &timestamp.Timestamp{
		Seconds: start.Seconds + 14*secondsInMinute,
		Nanos:   start.Nanos,
	}

	e := &event.Event{
		Title:        "Walking",
		Desc:         "At the park",
		Start:        start,
		End:          end,
		NotifyBefore: &duration.Duration{Seconds: 60 * secondsInMinute},
	}

	_, err := c.CreateEvent(ctx, e)
	if err == nil {
		t.Fatalf("got: nil, but expected: error")
	}

	checkFieldViolation(t, err, "end date")
}

func TestListEventsPerDate(t *testing.T) {
	cc, c, ctx := setup(t)
	defer cc.Close()

	now := nowDate()
	baseDate := &timestamp.Timestamp{
		Seconds: now.Seconds + secondsInDay,
		Nanos:   now.Nanos,
	}

	start := &timestamp.Timestamp{
		Seconds: baseDate.Seconds,
		Nanos:   baseDate.Nanos,
	}
	end := &timestamp.Timestamp{
		Seconds: start.Seconds + 15*secondsInMinute,
		Nanos:   start.Nanos,
	}

	e := &event.Event{
		Title:        "title1",
		Desc:         "desc1",
		Start:        start,
		End:          end,
		NotifyBefore: &duration.Duration{Seconds: 15 * secondsInMinute},
	}

	eupd, err := c.CreateEvent(ctx, e)
	if err != nil {
		t.Fatalf("unexpected error in CreateEvent: %v", err)
	}

	if eupd.Id == "" {
		t.Errorf("got empty event id")
	}

	ids := []string{eupd.Id}

	resp, err := c.ListEventsPerDate(ctx, &request.ByDate{Date: baseDate})
	if err != nil {
		t.Fatalf("unexpected error in ListEventsPerDate: %v", err)
	}

	if len(resp.Events) != 1 {
		t.Fatalf("got events length: %v, but expected: %v", len(resp.Events), 1)
	}

	for _, e := range resp.Events {
		if _, ok := find(ids, e.Id); !ok {
			t.Fatalf("unexpected event id: %s", e.Id)
		}
	}

	start = &timestamp.Timestamp{
		Seconds: e.End.Seconds,
		Nanos:   e.End.Nanos,
	}
	end = &timestamp.Timestamp{
		Seconds: start.Seconds + 15*secondsInMinute,
		Nanos:   start.Nanos,
	}

	e2 := &event.Event{
		Title:        "title2",
		Desc:         "desc2",
		Start:        start,
		End:          end,
		NotifyBefore: &duration.Duration{Seconds: 15 * secondsInMinute},
	}

	e2upd, err := c.CreateEvent(ctx, e2)
	if err != nil {
		t.Fatalf("unexpected error in CreateEvent: %v", err)
	}

	if e2upd.Id == "" {
		t.Errorf("got empty event id")
	}

	ids = append(ids, e2upd.Id)

	resp, err = c.ListEventsPerDate(ctx, &request.ByDate{Date: baseDate})
	if err != nil {
		t.Fatalf("unexpected error in ListEventsPerDate: %v", err)
	}

	if len(resp.Events) != 2 {
		t.Fatalf("got events length: %v, but expected: %v", len(resp.Events), 2)
	}

	for _, e := range resp.Events {
		if _, ok := find(ids, e.Id); !ok {
			t.Fatalf("unexpected event id: %s", e.Id)
		}
	}

	start = &timestamp.Timestamp{
		Seconds: baseDate.Seconds + secondsInDay,
		Nanos:   baseDate.Nanos,
	}
	end = &timestamp.Timestamp{
		Seconds: start.Seconds + 15*secondsInMinute,
		Nanos:   start.Nanos,
	}

	e3 := &event.Event{
		Title:        "title3",
		Desc:         "desc3",
		Start:        start,
		End:          end,
		NotifyBefore: &duration.Duration{Seconds: 15 * secondsInMinute},
	}

	e3upd, err := c.CreateEvent(ctx, e3)
	if err != nil {
		t.Fatalf("unexpected error in CreateEvent: %v", err)
	}

	if e3upd.Id == "" {
		t.Errorf("got empty event id")
	}

	resp, err = c.ListEventsPerDate(ctx, &request.ByDate{Date: baseDate})
	if err != nil {
		t.Fatalf("unexpected error in ListEventsPerDate: %v", err)
	}

	if len(resp.Events) != 2 {
		t.Fatalf("got events length: %v, but expected: %v", len(resp.Events), 2)
	}

	for _, e := range resp.Events {
		if _, ok := find(ids, e.Id); !ok {
			t.Fatalf("unexpected event id: %s", e.Id)
		}
	}
}
