package cpu

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/go-log"
	"sync"
	"time"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	cpuClient cpu.CPUService
}

func (p *Pusher) Init(opts modules.Options) error {
	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.cpuClient = cpu.NewCPUService(p.CollectorName, opts.Client)

	return nil
}

func (p *Pusher) Push() (err error) {
	once.Do(func() {
		for {
			if err = p.pushInfo(); err == nil {
				break
			} else {
				log.Logf("[Push] cpu pushPercent error, %s", err)
				time.Sleep(2 * time.Second)
			}
		}
	})

	if err = p.pushPercent(); err != nil {
		log.Logf("[Push] cpu pushPercent error, %s", err)
		return
	}
	if err = p.pushTimes(); err != nil {
		log.Logf("[Push] cpu pushTimes error, %s", err)
		return
	}
	return err
}
