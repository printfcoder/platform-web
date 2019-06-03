package disk

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro/go-micro/client"
	"sync"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	diskClient disk2.DiskService
}

func (p *Pusher) Init(opts modules.Options) error {

	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.diskClient = disk2.NewDiskService(p.CollectorName, client.DefaultClient)

	return nil
}

func (p *Pusher) Push() (err error) {

	once.Do(func() {
		for {
			if err = p.pushInfo(); err == nil {
				break
			}
		}
	})

	p.pushPercent()
	p.pushTimes()

	return err
}
