package v1

import (
	"github.com/micro-in-cn/micro-web/modules"
	"github.com/micro/cli"
	"sync"
)

func init() {
	m = &basicModule{
		name: "basic",
		path: "/b",
	}

	modules.Registry(m)
}

var (
	m *basicModule
	// Default address to bind to

	GatewayNamespaces  = []string{"go.micro.api"}
	WebNamespacePrefix = []string{"go.micro.web"}
)

// basicModule includes web, registry, CLI, Stats submodules.
type basicModule struct {
	name string
	path string
	sync.RWMutex
	api *api
}

func (m *basicModule) Name() string {
	return m.name
}

func (m *basicModule) Path() string {
	return m.path
}

func (m *basicModule) Init(*cli.Context) error {
	return nil
}

func (m *basicModule) Flags() []cli.Flag {
	return nil
}

func (m *basicModule) Handlers() (mp map[string]*modules.Handler) {

	m.Lock()
	mp = make(map[string]*modules.Handler)
	defer m.Unlock()

	mp["/services"] = &modules.Handler{
		Func:   m.api.services,
		Method: []string{"GET"},
	}

	mp["/micro-services"] = &modules.Handler{
		Func:   m.api.microServices,
		Method: []string{"GET"},
	}

	mp["/service/{name:[a-zA-Z0-9/.]+}"] = &modules.Handler{
		Func:   m.api.service,
		Method: []string{"GET"},
	}

	mp["/api-gateway-services"] = &modules.Handler{
		Func:   m.api.apiGatewayServices,
		Method: []string{"GET"},
	}

	mp["/service-details"] = &modules.Handler{
		Func:   m.api.serviceDetails,
		Method: []string{"GET"},
	}

	mp["/stats"] = &modules.Handler{
		Func:   m.api.stats,
		Method: []string{"GET"},
	}

	mp["/api-stats"] = &modules.Handler{
		Hld:    apiProxy(),
		Method: []string{"GET"},
	}

	mp["/web-services"] = &modules.Handler{
		Func:   m.api.webServices,
		Method: []string{"GET"},
	}

	mp["/rpc"] = &modules.Handler{
		Func:   m.api.rpc,
		Method: []string{"POST"},
	}

	mp["/health"] = &modules.Handler{
		Func:   m.api.health,
		Method: []string{"GET"},
	}

	return
}

func (m *basicModule) Commands(options ...modules.Option) []cli.Command {
	command := cli.Command{
		Name:  "web",
		Usage: "Run the web dashboard",
		Action: func(c *cli.Context) {
			// run(c, options...)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "address",
				Usage:  "Set the web UI address e.g 0.0.0.0:8082",
				EnvVar: "MICRO_WEB_ADDRESS",
			},
			cli.StringFlag{
				Name:   "namespace",
				Usage:  "Set the namespace used by the Web proxy e.g. com.example.web",
				EnvVar: "MICRO_WEB_NAMESPACE",
			},
			cli.StringFlag{
				Name:   "static_dir",
				Usage:  "Set the static dir of micro web",
				EnvVar: "MICRO_WEB_STATIC_DIR",
			},
		},
	}

	return []cli.Command{command}
}
