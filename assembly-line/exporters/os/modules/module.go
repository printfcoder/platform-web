package modules

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/util/go/lib/addr"
)

var (
	once         sync.Once
	ip, nodeName string
)

type Module interface {
	Init(options *Options)
	Interval() time.Duration
	Push() error
	Start() error
	String() string
}

type BaseModule struct {
	Err chan error
	Module
	NodeName string
	IP       string
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
		t := time.NewTicker(time.Second * b.Interval())
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
