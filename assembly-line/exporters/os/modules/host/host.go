package host

import (
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	hostClient proto.HostService
}

func (p *Pusher) Init(opts modules.Options) error {
	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.hostClient = proto.NewHostService(p.CollectorName, opts.Client)

	return nil
}

func (p *Pusher) Push() (err error) {
	err = p.pushInfo()
	return err
}
