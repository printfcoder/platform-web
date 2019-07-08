package disk

import (
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro/go-log"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	path       []string
	diskClient disk2.DiskService
}

func (p *Pusher) Init(opts modules.Options) error {
	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.path = opts.DiskPaths
	p.diskClient = disk2.NewDiskService(p.CollectorName, opts.Client)

	return nil
}

func (p *Pusher) Push() (err error) {
	if err = p.pushIOCounters(); err != nil {
		log.Logf("[Push] pushIOCounters err: %s", err)
	}

	if err = p.pushUsage(); err != nil {
		log.Logf("[Push] pushUsage err: %s", err)
	}

	if err = p.pushPartition(); err != nil {
		log.Logf("[Push] pushPartition err: %s", err)
	}

	return err
}
