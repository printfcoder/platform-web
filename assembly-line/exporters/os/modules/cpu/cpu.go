package cpu

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/go-micro/util/log"
)

var (
	once sync.Once
)

type CPU struct {
	cpuClient cpu.CPUService
	opts      *modules.CPUOptions
}

func (c *CPU) Init(opts *modules.Options) {
	c.opts = opts.CPU
	c.opts.NodeName = opts.NodeName
	c.opts.IP = opts.IP
	c.cpuClient = cpu.NewCPUService(opts.Collector.Name, opts.Collector.Client)

	return
}

func (c *CPU) Push() (err error) {
	once.Do(func() {
		for {
			if err = c.pushInfo(); err == nil {
				break
			} else {
				log.Logf("[ERR] [Push] cpu pushPercent error, %s", err)
				time.Sleep(2 * time.Second)
			}
		}
	})

	/*if err = c.pushPercent(); err != nil {
		log.Logf("[ERR] [Push] cpu pushPercent error, %s", err)
		return
	}*/
	if err = c.pushTimes(); err != nil {
		log.Logf("[ERR] [Push] cpu pushTimes error, %s", err)
		return
	}
	return err
}

func (c *CPU) Start() (err error) {
	go func() {
		t := time.NewTicker(time.Second * c.Interval())
		for {
			select {
			case <-t.C:
				if err = c.Push(); err != nil {
					log.Logf("[ERR] [Start] cpu push err: %s", err)
				}
			}
		}
	}()

	return nil
}

func (c *CPU) Interval() time.Duration {
	return c.opts.Interval
}

func (c *CPU) String() string {
	return "cpu"
}
