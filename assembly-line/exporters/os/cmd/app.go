package cmd

import (
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/addr"
	"github.com/micro/go-micro/util/log"
)

var (
	diskPaths  []string
	netKinds   = []string{"all"}
	configFile = "./conf/os.yml"
	once2      sync.Once
)

func init() {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	diskPaths = []string{path}
}

func (app *c) loadConfig(ctx *cli.Context) {
	if len(ctx.String("config_file")) > 0 {
		configFile = ctx.String("config_file")
	}

	if err := config.Load(file.NewSource(file.WithPath(configFile))); err != nil {
		panic(err)
	}

	v := config.Get("micro", "platform-web", "assembly-line", "exporters")
	err := v.Scan(&app.opts)
	if err != nil {
		panic(err)
	}

	var ip string
	nodeName, err := os.Hostname()
	if err != nil {
		log.Logf("[ERR] [loadConfig] get host name error: %s", err)
	}

	// then choose ip
	ips := addr.IPs()
	log.Logf("[INFO] [loadConfig] got ips: %s", ips)

	// find the first one which is not prefix with 127 and not ipv6
	for _, ipTemp := range ips {
		if strings.Index(ipTemp, "127") == 0 || strings.Count(ipTemp, ":") > 1 {
			continue
		}

		ip = ipTemp
		if nodeName == "" {
			nodeName = ipTemp
		}

		log.Logf("[INFO] [loadConfig] node ip: %s and name: %s", ip, nodeName)

		break
	}

	app.opts.NodeName = nodeName
	app.opts.IP = ip
}

func (app *c) run() {
	s := micro.NewService(
		micro.Name(name),
		micro.Version(version),
	)

	s.Init(micro.Action(func(ctx *cli.Context) {
		app.loadConfig(ctx)
		app.loadModules(s.Client())
		app.start()
	}))

	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (app *c) start() {
	for _, m := range app.modules {
		if err := m.Start(); err != nil {
			log.Logf("[start] error, %s", err)
		}
	}
}
