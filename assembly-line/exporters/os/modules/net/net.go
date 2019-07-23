package net

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
	"github.com/micro/go-micro/util/log"
)

var (
	once sync.Once
)

type Net struct {
	modules.BaseModule
	opts      *modules.NetOptions
	netClient proto.NetService
	ifaceMap  map[string]bool
}

func (n *Net) Init(opts *modules.Options) {
	n.opts = opts.Net
	n.InitB()
	n.netClient = proto.NewNetService(n.opts.Collector.Name, n.opts.Collector.Client)

	// for search ifraces are used to be exported
	for _, i := range n.opts.Ifaces {
		n.ifaceMap[i] = true
	}

	return
}

func (n *Net) Push() (err error) {
	if err = n.pushConnectionStat(); err != nil {
		log.Logf("[ERR] [Push] pushConnectionStat err: %s", err)
	}

	if err = n.pushIOCountersStat(); err != nil {
		log.Logf("[ERR] [Push] pushIOCountersStat err: %s", err)
	}

	return err
}

func (n *Net) Interval() time.Duration {
	return n.opts.Interval
}

func (n *Net) String() string {
	return "net"
}
