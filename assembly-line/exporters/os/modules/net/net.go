package net

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"sync"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
	"github.com/micro/go-log"
)

var (
	once sync.Once
)

type Net struct {
	modules.BaseModule
	kinds     []string
	netClient proto.NetService
}

func (p *Net) Init(opts option.Options) error {
	p.InitB()
	p.CollectorName = opts.Collector.Name
	p.Interval = opts.Net.Interval
	p.kinds = opts.Net.Kinds
	p.netClient = proto.NewNetService(p.CollectorName, opts.Collector.Client)

	return nil
}

func (p *Net) Push() (err error) {
	if err = p.pushConnectionStat(); err != nil {
		log.Logf("[Push] pushConnectionStat err: %s", err)
	}

	if err = p.pushIOCountersStat(); err != nil {
		log.Logf("[Push] pushIOCountersStat err: %s", err)
	}

	return err
}
