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
	opts      *modules.NetOptions
	netClient proto.NetService
	ifaceMap  map[string]bool
}

func (n *Net) Init(opts *modules.Options) {
	n.opts = opts.Net
	n.opts.NodeName = opts.NodeName
	n.opts.IP = opts.IP
	n.netClient = proto.NewNetService(n.opts.Collector.Name, n.opts.Collector.Client)
	n.ifaceMap = map[string]bool{}

	// ifraces which are used to be exported
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

func (n *Net) Start() (err error) {
	go func() {
		t := time.NewTicker(time.Second * n.Interval())
		for {
			select {
			case <-t.C:
				if err = n.Push(); err != nil {
				}
			}
		}
	}()

	return nil
}

func (n *Net) Interval() time.Duration {
	return n.opts.Interval
}

func (n *Net) String() string {
	return "net"
}
