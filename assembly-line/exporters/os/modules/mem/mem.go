package mem

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

type Pusher struct {
	modules.BasePusher
	memClient proto.MemService
}

func (p *Pusher) Init(opts modules.Options) error {
	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.memClient = proto.NewMemService(p.CollectorName, opts.Client)

	return nil
}

func (p *Pusher) Push() (err error) {
	err = p.pushVirtualMemoryStat()
	if err != nil {
		return
	}

	err = p.pushSwapMemoryStat()
	if err != nil {
		return
	}

	return err
}
