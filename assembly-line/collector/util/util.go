package util

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func PTimestamp(ts *timestamp.Timestamp) (t time.Time) {
	t, _ = ptypes.Timestamp(ts)
	return t
}
