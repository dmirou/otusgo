// +build integration

package integration

import (
	"testing"
	"time"

	"github.com/dmirou/otusgo/calendar/pkg/helper"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

// nowDate returns protobuf timestamp with the start of the current day
// (zero hours, minutes, seconds)
func nowDate() *timestamp.Timestamp {
	now := time.Now()

	return helper.NewProtoDate(now.Year(), now.Month(), now.Day())
}

// checkFieldViolation checks that we received the error
// with the field violation in error's details
func checkFieldViolation(t *testing.T, err error, field string) {
	st := status.Convert(err)
	if len(st.Details()) != 1 {
		t.Fatalf("unexpected error details count: %v, expected: %v", len(st.Details()), 1)
	}

	detail := st.Details()[0]

	dt, ok := detail.(*errdetails.BadRequest)
	if !ok {
		t.Fatalf("unexpected detail: %v", detail)
	}

	fvs := dt.GetFieldViolations()
	if len(fvs) != 1 {
		t.Fatalf("unexpected field violations count: %v, expected: %v", len(fvs), 1)
	}

	if fvs[0].Field != field {
		t.Errorf(
			"got violation field: %s, but expected: %s",
			fvs[0].Field, field,
		)
	}
}

// find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
