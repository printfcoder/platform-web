package cmd

import (
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
	opts option.Options
}

// Init app
func Init(ops ...option.Option) {

	once.Do(func() {
		app := newApp(ops...)
		app.parseFlags()
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
		},
	}

	for _, o := range ops {
		o(&app.opts)
	}

	return app
}
