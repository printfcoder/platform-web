package modules

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro/go-log"
	"github.com/micro/util/go/lib/addr"
)

var (
	once         sync.Once
	ip, nodeName string
)

type Module interface {
	Init(opts option.Options) error
	Push() error
	Start() error
}

type BaseModule struct {
	Module
	CollectorName string
	Interval      time.Duration
	NodeName      string
	IP            string
	Err           chan error
}

func (b *BaseModule) InitB() {
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

func (b *BaseModule) Start() (err error) {
	go func() {
		t := time.NewTicker(b.Interval * time.Second)
		for {
			select {
			case <-t.C:
				log.Logf("push data, %s", time.Now())
				if err = b.Push(); err != nil {
					b.Err <- err
				}
			}
		}
	}()

	return nil
}
