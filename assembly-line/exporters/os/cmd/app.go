package cmd

import (
	"runtime"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	diskPaths  []string
	netKinds   = []string{"all"}
	configFile = "./conf/os.yml"
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
	log.Log(app.opts.Collector.Name)
}

func (app *c) run() {
	s := micro.NewService(
		micro.Name(name),
		micro.Version(version),
	)

	s.Init(micro.Action(func(ctx *cli.Context) {
		app.loadConfig(ctx)
		app.parseFlags(ctx)
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
