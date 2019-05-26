package cpu

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	"github.com/micro/go-log"
)

func (p *pusher) pushInfo() {

	log.Logf("[pushInfo] %s", 2)

	v, err := cpu.Info()
	if err != nil {

	}

	log.Logf("[pushInfo] %s", v)
}
