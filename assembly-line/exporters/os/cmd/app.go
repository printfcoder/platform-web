package cmd

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/cpu"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"time"
)

func (app *c) advFlags() {

	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "enable_cpu",
			Usage:  "enables exporter to export CPU info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_CPU",
		},
		cli.StringFlag{
			Name:   "enable_disk",
			Usage:  "enables exporter to export disk info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_DISK",
		},
		cli.StringFlag{
			Name:   "enable_docker",
			Usage:  "enables exporter to export docker info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_DOCKER",
		},
		cli.StringFlag{
			Name:   "enable_host",
			Usage:  "enables exporter to export host info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_HOST",
		},
		cli.StringFlag{
			Name:   "enable_load",
			Usage:  "enables exporter to export load info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_LOAD",
		},
		cli.StringFlag{
			Name:   "enable_mem",
			Usage:  "enables exporter to export memory info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_MEM",
		},
		cli.StringFlag{
			Name:   "enable_net",
			Usage:  "enables exporter to export net info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_NET",
		},
		cli.StringFlag{
			Name:   "enable_process",
			Usage:  "enables exporter to export process info",
			EnvVar: "MICRO_WEB_PLATFORM_ENABLE_PROCESS",
		},
		cli.StringFlag{
			Name:   "push_interval",
			Usage:  "Push interval in seconds",
			EnvVar: "MICRO_WEB_PLATFORM_PUSH_INTERVAL",
		},
		cli.StringFlag{
			Name:   "collector",
			Usage:  "name of collector service",
			EnvVar: "MICRO_WEB_PLATFORM_COLLECTOR",
		},
	)
}

func (app *c) parseFlags(ctx *cli.Context) {

	if len(ctx.String("enable_cpu")) > 0 && !ctx.Bool("enable_cpu") {
		app.opts.EnableCPU = false
	}

	if len(ctx.String("enable_disk")) > 0 && !ctx.Bool("enable_disk") {
		app.opts.EnableDisk = false
	}

	if len(ctx.String("enable_docker")) > 0 && !ctx.Bool("enable_docker") {
		app.opts.EnableDocker = false
	}

	if len(ctx.String("enable_host")) > 0 && !ctx.Bool("enable_host") {
		app.opts.EnableHost = false
	}

	if len(ctx.String("enable_load")) > 0 && !ctx.Bool("enable_load") {
		app.opts.EnableLoad = false
	}

	if len(ctx.String("enable_mem")) > 0 && !ctx.Bool("enable_mem") {
		app.opts.EnableMem = false
	}

	if len(ctx.String("enable_net")) > 0 && !ctx.Bool("enable_net") {
		app.opts.EnableNet = false
	}

	if len(ctx.String("enable_process")) > 0 && !ctx.Bool("enable_process") {
		app.opts.EnableProcess = false
	}

	if ctx.Int("push_interval") > 0 {
		app.opts.PushInterval = time.Duration(ctx.Int("push_interval"))
	}

	if ctx.Int("collector") > 0 {
		app.opts.CollectorName = ctx.String("collector")
	}

}

func (app *c) loadModules() {

	// cpu
	if app.opts.EnableCPU {
		p := cpu.Pusher{}
		opts := modules.Options{
			CollectorName: app.opts.CollectorName,
		}
		p.Init(opts)

		app.modules = append(app.modules, &p)
	}
}

func (app *c) run() {
	s := micro.NewService(
		micro.Name(name),
		micro.Version(version),
	)

	s.Init(micro.Action(func(ctx *cli.Context) {
		app.parseFlags(ctx)
		app.loadModules()

		go func() {
			t := time.NewTicker(app.opts.PushInterval * time.Second)
			for {
				select {
				case <-t.C:
					log.Logf("push data, %s", time.Now())
					app.push()
				}
			}
		}()
	}))

	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (app *c) push() {
	for _, m := range app.modules {
		if err := m.Push(); err != nil {
			log.Logf("[push] error, %s", err)
		}
	}
}
