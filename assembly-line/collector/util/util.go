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

func Int32ArrayTo64(in []int32) []int64 {
	ret := make([]int64, len(in))
	for _, v := range in {
		ret = append(ret, int64(v))
	}

	return ret
}
