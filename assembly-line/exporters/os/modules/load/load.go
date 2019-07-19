package load

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
)

var (
	once sync.Once
)

type Load struct {
	modules.BaseModule
	loadClient load.LoadService
}

func (p *Load) Init(opts option.Options) error {
	p.InitB()
	p.CollectorName = opts.Collector.Name
	p.Interval = opts.Load.Interval
	p.loadClient = load.NewLoadService(p.CollectorName, opts.Collector.Client)

	return nil
}

func (p *Load) Push() (err error) {
	err = p.pushAvgStat()
	return err
}
