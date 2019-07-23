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
		c := cpu.CPU{}
		app.opts.CPU.Collector = app.opts.Collector
		c.Init(app.opts)

		app.modules = append(app.modules, &c)
	}

	// disk
	if app.opts.Disk.Enabled {
		d := disk.Disk{}
		app.opts.Disk.Collector = app.opts.Collector
		d.Init(app.opts)

		app.modules = append(app.modules, &d)
	}

	// host
	if app.opts.Host.Enabled {
		h := host.Host{}
		app.opts.Host.Collector = app.opts.Collector
		h.Init(app.opts)

		app.modules = append(app.modules, &h)
	}

	// load
	if app.opts.Load.Enabled {
		l := load.Load{}
		app.opts.Load.Collector = app.opts.Collector
		l.Init(app.opts)

		app.modules = append(app.modules, &l)
	}

	// mem
	if app.opts.Mem.Enabled {
		m := mem.Mem{}
		app.opts.Mem.Collector = app.opts.Collector
		m.Init(app.opts)

		app.modules = append(app.modules, &m)
	}

	// net
	if app.opts.Net.Enabled {
		n := net.Net{}
		app.opts.Net.Collector = app.opts.Collector
		n.Init(app.opts)

		app.modules = append(app.modules, &n)
	}
}
