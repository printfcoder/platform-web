package disk

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro/go-log"
)

var (
	once sync.Once
)

type Disk struct {
	modules.BaseModule
	path       []string
	diskClient disk2.DiskService
}

func (p *Disk) Init(opts option.Options) error {
	p.InitB()
	p.CollectorName = opts.Collector.Name
	p.Interval = opts.Disk.Interval
	p.path = opts.Disk.Paths
	p.diskClient = disk2.NewDiskService(p.CollectorName, opts.Collector.Client)

	return nil
}

func (p *Disk) Push() (err error) {
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
