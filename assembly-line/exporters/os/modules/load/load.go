package load

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
	"sync"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	loadClient load.LoadService
}

func (p *Pusher) Init(opts modules.Options) error {
	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.loadClient = load.NewLoadService(p.CollectorName, opts.Client)

	return nil
}

func (p *Pusher) Push() (err error) {
	err = p.pushAvgStat()
	return err
}
