package cmd

import (
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro/cli"
)

func (app *c) load(ops []option.Option) {
	app.ops = ops

	app.flags()
	app.action()
    app.loadModules()
}

func (app *c) action() {
	app.Action = func(c *cli.Context) {
		app.run(c, app.ops...)
	}
}

func (app *c) flags() {
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
	)
}

func (app *c) run(ctx *cli.Context, opts ...option.Option) {

}

func (app *c) loadModules() {

}
