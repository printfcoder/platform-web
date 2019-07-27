package load

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
	"github.com/micro/go-micro/util/log"
)

var (
	once sync.Once
)

type Load struct {
	loadClient load.LoadService
	opts       *modules.LoadOptions
}

func (l *Load) Init(opts *modules.Options) {
	l.opts = opts.Load
	l.opts.NodeName = opts.NodeName
	l.opts.IP = opts.IP
	l.loadClient = load.NewLoadService(opts.Collector.Name, opts.Collector.Client)

	return
}

func (l *Load) Push() (err error) {
	err = l.pushAvgStat()
	return err
}

func (l *Load) Start() (err error) {
	go func() {
		t := time.NewTicker(time.Second * l.Interval())
		for {
			select {
			case <-t.C:
				if err = l.Push(); err != nil {
					log.Logf("[ERR] [Start] load push err: %s", err)
				}
			}
		}
	}()

	return nil
}

func (l *Load) Interval() time.Duration {
	return l.opts.Interval
}

func (l *Load) String() string {
	return "load"
}
