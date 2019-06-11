package net

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
	"github.com/micro/go-micro/client"
	"sync"
)

var (
	once sync.Once
)

type Pusher struct {
	modules.BasePusher
	kinds     []string
	netClient proto.NetService
}

func (p *Pusher) Init(opts modules.Options) error {
	p.InitB()
	p.CollectorName = opts.CollectorName
	p.Interval = opts.Interval
	p.kinds = opts.NetKinds
	p.netClient = proto.NewNetService(p.CollectorName, client.DefaultClient)

	return nil
}

func (p *Pusher) Push() (err error) {
	err = p.pushConnectionStat()
	return err
}
