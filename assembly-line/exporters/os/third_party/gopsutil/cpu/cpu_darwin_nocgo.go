// +build darwin
// +build !cgo

package cpu

import "github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/internal/common"

func perCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func allCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}
