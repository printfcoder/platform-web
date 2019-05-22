package cmd

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
)

var (
	name      = "go.micro.web.platform_exporter_os"
	collector = "go.micro.srv.platform_collector"
	app       = &c{
		App: cmd.App(),
	}
)

type c struct {
	*cli.App
	ops []option.Option
}

// Init app
func Init(ops ...option.Option) {

	app.load(ops)
	
	if err := cmd.Init(cmd.Name(name)); err != nil {
		panic(err)
	}
}




