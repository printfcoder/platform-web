package modules

import (
	"github.com/micro/go-log"
	"github.com/micro/util/go/lib/addr"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	once         sync.Once
	ip, nodeName string
)

type Pusher interface {
	Init(opts Options) error
	Push() error
}

type BasePusher struct {
	CollectorName string
	Interval      time.Duration
	NodeName      string
	IP            string
}

func (b *BasePusher) InitB() {

	once.Do(func() {

		var err error
		nodeName, err = os.Hostname()
		if err != nil {
			log.Logf("[InitB] get host name error: %s", err)
		}

		// then choose ip
		ips := addr.IPs()
		log.Logf("[InitB] got ips: %s", ips)

		// find the first one which is not prefix with 127
		for _, ipTemp := range ips {
			if strings.Index(ipTemp, "127") == 0 {
				continue
			}

			ip = ipTemp
			if nodeName == "" {
				nodeName = ipTemp
			}

			break
		}
	})

	b.IP = ip
	b.NodeName = nodeName

}

type Options struct {
	CollectorName string
	Interval      time.Duration
	DiskPath      []string
}

type Option func(*Options)
