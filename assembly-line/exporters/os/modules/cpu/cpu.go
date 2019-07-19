package cpu

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/go-log"
)

var (
	once sync.Once
)

type CPU struct {
	modules.BaseModule
	cpuClient cpu.CPUService
}

func (p *CPU) Init(opts option.Options) error {
	p.InitB()
	p.CollectorName = opts.Collector.Name
	p.Interval = opts.CPU.Interval
	p.cpuClient = cpu.NewCPUService(p.CollectorName, opts.Collector.Client)

	return nil
}

func (p *CPU) Push() (err error) {
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


