package cmd

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/cpu"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/disk"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/host"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/load"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/mem"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/net"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
)

func (app *c) loadModules(client client.Client) {
	log.Logf("[INFO] loadModules")

	// cpu
	if app.opts.CPU.Enabled {
		log.Logf("[INFO] cpu enabled")
		p := cpu.CPU{}
		_ = p.Init(app.opts)

		app.modules = append(app.modules, &p)
	}

	// disk
	if app.opts.Disk.Enabled {
		p := disk.Disk{}
		_ = p.Init(app.opts)

		app.modules = append(app.modules, &p)
	}

	// host
	if app.opts.Host.Enabled {
		p := host.Host{}
		_ = p.Init(app.opts)

		app.modules = append(app.modules, &p)
	}

	// load
	if app.opts.Load.Enabled {
		p := load.Load{}
		_ = p.Init(app.opts)

		app.modules = append(app.modules, &p)
	}

	// mem
	if app.opts.Mem.Enabled {
		p := mem.Mem{}
		_ = p.Init(app.opts)

		app.modules = append(app.modules, &p)
	}

	// net
	if app.opts.Net.Enabled {
		p := net.Net{}
		_ = p.Init(app.opts)

		app.modules = append(app.modules, &p)
	}
}
