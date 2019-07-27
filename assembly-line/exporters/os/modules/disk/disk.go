package disk

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro/go-log"
)

var (
	once sync.Once
)

type Disk struct {
	diskClient disk2.DiskService
	opts *modules.DiskOptions
}

func (d *Disk) Init(opts *modules.Options) {
	d.opts = opts.Disk
	d.opts.NodeName = opts.NodeName
	d.opts.IP = opts.IP
	d.diskClient = disk2.NewDiskService(opts.Collector.Name, opts.Collector.Client)

	return
}

func (d *Disk) Push() (err error) {
	if err = d.pushIOCounters(); err != nil {
		log.Logf("[ERR] [Push] pushIOCounters err: %s", err)
	}

	if err = d.pushUsage(); err != nil {
		log.Logf("[ERR] [Push] pushUsage err: %s", err)
	}

	if err = d.pushPartition(); err != nil {
		log.Logf("[ERR] [Push] pushPartition err: %s", err)
	}

	return err
}

func (d *Disk) Start() (err error) {
	go func() {
		t := time.NewTicker(time.Second * d.Interval())
		for {
			select {
			case <-t.C:
				if err = d.Push(); err != nil {
					log.Logf("[ERR] [Start] disk push err: %s", err)
				}
			}
		}
	}()

	return nil
}

func (d *Disk) Interval() time.Duration {
	return d.opts.Interval
}

func (d *Disk) String() string {
	return "disk"
}
