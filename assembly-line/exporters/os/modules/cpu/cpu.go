package cpu

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/go-log"
	"sync"
	"time"
)

var (
	collectorName = "go.micro.srv.platform_collector"
	p             *pusher
	once          sync.Once
	once2         sync.Once
	cpuClient     cpu.CPUService
)

type pusher struct {
	interval     time.Duration
	pushFucQueue []func()
}

func Init(opts option.Options) {

	once.Do(func() {
		p = &pusher{
			interval: opts.PushInterval,
		}

	})

	prepareDo()

	go func() {
		t := time.NewTicker(p.interval * time.Second)
		for {
			select {
			case <-t.C:
				log.Logf("push data, %s", time.Now())
				p.push()
			}
		}
	}()
}

func prepareDo() {

}

func (p *pusher) push() {
	p.pushInfo()
	p.pushPercent()
	p.pushStatus()
	p.pushTimes()
}
