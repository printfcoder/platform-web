package cmd

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro/cli"
	"github.com/micro/go-micro/config/cmd"
)

var (
	name         = "go.micro.srv.platform_exporter_os"
	version      = "v1"
	collector    = "go.micro.srv.platform_collector"
	pushInterval = time.Duration(3)
	once         sync.Once
)

type c struct {
	*cli.App
	modules []modules.Module
	opts    option.Options
}

// Init app
func Init(ops ...option.Option) {
	once.Do(func() {
		app := newApp(ops...)
		app.advFlags()
		app.run()
	})
}

func newApp(ops ...option.Option) (app *c) {
	app = &c{
		App: cmd.App(),
		opts: option.Options{
			AppName:    name,
			AppVersion: version,
			Collector: &option.Collector{
				Name: collector,
			},
			CPU: &option.CPUOptions{
				Enabled: true,
			},
			Disk: &option.DiskOptions{
				Enabled: true,
			},
			Docker: &option.DockerOptions{
				Enabled: true,
			},
			Host: &option.HostOptions{
				Enabled: true,
			},
			Load: &option.LoadOptions{
				Enabled: true,
			},
			Mem: &option.MemOptions{
				Enabled: true,
			},
			Net: &option.NetOptions{
				Enabled: true,
			},
			Process: &option.ProcessOptions{
				Enabled: true,
			},
		},
	}

	for _, o := range ops {
		o(&app.opts)
	}

	return app
}
