package host

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
)

var (
	once sync.Once
)

type Host struct {
	modules.BaseModule
	hostClient proto.HostService
}

func (p *Host) Init(opts option.Options) error {
	p.InitB()
	p.CollectorName = opts.Collector.Name
	p.Interval = opts.Host.Interval
	p.hostClient = proto.NewHostService(p.CollectorName, opts.Collector.Client)

	return nil
}

func (p *Host) Push() (err error) {
	err = p.pushInfo()
	return err
}
