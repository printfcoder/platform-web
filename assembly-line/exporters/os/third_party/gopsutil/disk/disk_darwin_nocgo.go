// +build darwin
// +build !cgo

package disk

import (
	"context"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/internal/common"
)

func IOCounters(names ...string) (map[string]IOCountersStat, error) {
	return IOCountersWithContext(context.Background(), names...)
}

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
