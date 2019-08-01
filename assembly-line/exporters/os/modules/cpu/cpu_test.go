package cpu

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	"testing"
)

func TestCPUTimes(t *testing.T) {
	vvPer, _ := cpu.Times(true)

	for _, v := range vvPer {
		t.Logf("user %f, system: %f, idle: %f",
			v.User/(v.User+v.System+v.Idle), v.System/(v.User+v.System+v.Idle), v.Idle/(v.User+v.System+v.Idle))
	}
}
