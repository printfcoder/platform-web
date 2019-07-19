package mem

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

type Mem struct {
	modules.BaseModule
	memClient proto.MemService
}

func (p *Mem) Init(opts option.Options) error {
	p.InitB()
	p.CollectorName = opts.Collector.Name
	p.Interval = opts.Mem.Interval
	p.memClient = proto.NewMemService(p.CollectorName, opts.Collector.Client)

	return nil
}

func (p *Mem) Push() (err error) {
	err = p.pushMemoryStat()
	if err != nil {
		return
	}

	return err
}
