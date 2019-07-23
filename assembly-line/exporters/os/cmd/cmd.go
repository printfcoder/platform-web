package cmd

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
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
	opts    *modules.Options
}

// Init app
func Init(ops ...modules.Option) {
	once.Do(func() {
		app := newApp(ops...)
		app.advFlags()
		app.run()
	})
}

func newApp(ops ...modules.Option) (app *c) {
	app = &c{
		App: cmd.App(),
	}

	for _, o := range ops {
		o(app.opts)
	}

	return app
}
