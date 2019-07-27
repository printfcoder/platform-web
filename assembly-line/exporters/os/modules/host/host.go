package host

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
	"github.com/micro/go-micro/util/log"
)

var (
	once sync.Once
)

type Host struct {
	hostClient proto.HostService
	opts       *modules.HostOptions
}

func (h *Host) Init(opts *modules.Options) {
	h.opts = opts.Host
	h.opts.NodeName = opts.NodeName
	h.opts.IP = opts.IP
	h.hostClient = proto.NewHostService(opts.Collector.Name, opts.Collector.Client)

	return
}

func (h *Host) Push() (err error) {
	err = h.pushInfo()
	return err
}

func (h *Host) Start() (err error) {
	go func() {
		t := time.NewTicker(time.Second * h.Interval())
		for {
			select {
			case <-t.C:
				if err = h.Push(); err != nil {
					log.Logf("[ERR] [Start] host push err: %s", err)
				}
			}
		}
	}()

	return nil
}

func (h *Host) Interval() time.Duration {
	return h.opts.Interval
}

func (h *Host) String() string {
	return "host"
}
