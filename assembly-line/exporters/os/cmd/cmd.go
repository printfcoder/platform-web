package cmd

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"sync"
)

var (
	name      = "go.micro.srv.platform_exporter_os"
	version   = "1.0.0"
	collector = "go.micro.srv.platform_collector"
	app       = &c{App: cmd.App()}
	once      sync.Once
)

type c struct {
	*cli.App
	ops []option.Option
}

// Init app
func Init(ops ...option.Option) {

	once.Do(func() {
		app.load(ops)
		app.run()
	})
}
