package os

import (
	"sync"

	"github.com/micro-in-cn/platform-web/modules"
	"github.com/micro/cli"
)

var (
	m *module
)

// module includes cpu, disk, docker, host, load, mem, net
type module struct {
	name string
	path string
	sync.RWMutex
	api *api
}

func (m *module) Name() string {
	return m.name
}

func (m *module) Path() string {
	return m.path
}

func (m *module) Init(ctx *cli.Context) error {
	m.api.init(ctx)
	return nil
}

func (m *module) Flags() []cli.Flag {
	return nil
}

func (m *module) Handlers() (mp map[string]*modules.Handler) {
	m.Lock()
	mp = make(map[string]*modules.Handler)
	defer m.Unlock()

	mp["/ip-group"] = &modules.Handler{
		Func:   m.api.ipGroup,
		Method: []string{"GET"},
	}

	mp["/cpu/infos"] = &modules.Handler{
		Func:   m.api.cpuInfos,
		Method: []string{"GET"},
	}
	mp["/cpu/percent"] = &modules.Handler{
		Func:   m.api.cpuPercent,
		Method: []string{"GET"},
	}
	mp["/cpu/times"] = &modules.Handler{
		Func:   m.api.cpuTimes,
		Method: []string{"GET"},
	}

	mp["/disk"] = &modules.Handler{
		Func:   m.api.disk,
		Method: []string{"GET"},
	}

	mp["/host"] = &modules.Handler{
		Func:   m.api.host,
		Method: []string{"GET"},
	}

	mp["/load/avg-stat"] = &modules.Handler{
		Func:   m.api.loadAvgStat,
		Method: []string{"GET"},
	}

	mp["/mem/percents"] = &modules.Handler{
		Func:   m.api.memPercent,
		Method: []string{"GET"},
	}

	mp["/net"] = &modules.Handler{
		Func:   m.api.net,
		Method: []string{"GET"},
	}

	return
}

func (m *module) Commands(options ...modules.Option) []cli.Command {
	command := cli.Command{
		Action: func(c *cli.Context) {
			err := m.Init(c)
			if err != nil {
				panic(err)
			}
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "enable_os",
				Usage:  "set os module enabled",
				EnvVar: "MICRO_PLATFORM_WEB_ENABLE_OS",
				Value:  "true",
			},
		},
	}

	return []cli.Command{command}
}

func init() {
	m = &module{
		name: "os",
		path: "/os",
		api:  newAPI(),
	}

	modules.Registry(m)
}
