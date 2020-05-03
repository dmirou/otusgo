package helper

import (
	cevent "github.com/dmirou/otusgo/calendar/pkg/contracts/event"
	"github.com/dmirou/otusgo/calendar/pkg/event"
)

// EventFromProtobuf convert protobuf Event to Event
func EventFromProtobuf(e *cevent.Event) (*event.Event, error) {
	re := &event.Event{}

	re.ID = e.Id
	re.UserID = e.UserId
	re.Title = e.Title
	re.Desc = e.Desc

	start, err := ProtobufToTime(*e.Start)
	if err != nil {
		return nil, err
	}

	re.Start = start

	end, err := ProtobufToTime(*e.End)
	if err != nil {
		return nil, err
	}

	re.End = end

	nb, err := ProtobufToDuration(*e.NotifyBefore)
	if err != nil {
		return nil, err
	}

	re.NotifyBefore = nb

	return re, nil
}

// EventToProtobuf convert Event to protobuf Event
func EventToProtobuf(e *event.Event) *cevent.Event {
	ce := &cevent.Event{}

	ce.Id = e.ID
	ce.UserId = e.UserID
	ce.Title = e.Title
	ce.Desc = e.Desc

	ce.Start = TimeToProtobuf(e.Start)
	ce.End = TimeToProtobuf(e.End)
	ce.NotifyBefore = DurationToProtobuf(e.NotifyBefore)

	return ce
}
