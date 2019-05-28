package cmd

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"sync"
	"time"
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
	modules []modules.Pusher
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
			AppName:       name,
			AppVersion:    version,
			PushInterval:  pushInterval,
			CollectorName: collector,
			EnableCPU:     true,
			EnableDisk:    true,
			EnableDocker:  true,
			EnableHost:    true,
			EnableLoad:    true,
			EnableMem:     true,
			EnableNet:     true,
			EnableProcess: true,
		},
	}

	for _, o := range ops {
		o(&app.opts)
	}

	return app
}