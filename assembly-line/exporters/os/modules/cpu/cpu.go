package cpu

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/go-micro/client"
	"sync"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	cpuClient cpu.CPUService
}

func (p *Pusher) Init(opts modules.Options) error {

	p.CollectorName = opts.CollectorName
	p.cpuClient = cpu.NewCPUService(p.CollectorName, client.DefaultClient)

	return nil
}

func (p *Pusher) Push() (err error) {

	if err = p.pushInfo(); err != nil {
		return err
	}

	p.pushPercent()
	p.pushStatus()
	p.pushTimes()

	return err
}
