package net

import (
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
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
	if err = p.pushConnectionStat(); err != nil {
		log.Logf("[Push] pushConnectionStat err: %s", err)
	}

	if err = p.pushIOCountersStat(); err != nil {
		log.Logf("[Push] pushIOCountersStat err: %s", err)
	}

	return err
}
