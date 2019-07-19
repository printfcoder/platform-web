package cmd

import "github.com/micro/cli"

func (app *c) advFlags() {
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "config_file",
			Usage:  "path to config file, but the configs in file are less weight than those passed on command",
			EnvVar: "MICRO_WEB_PLATFORM_CONFIG_FILE",
		},
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
		cli.StringFlag{
			Name:   "group_name",
			Usage:  "name of server group",
			EnvVar: "MICRO_WEB_PLATFORM_GROUP",
		},
	)
}

func (app *c) parseFlags(ctx *cli.Context) {
	if len(ctx.String("enable_cpu")) > 0 {
		app.opts.CPU.Enabled = ctx.Bool("enable_cpu")
	}

	if len(ctx.String("enable_disk")) > 0 {
		app.opts.Disk.Enabled = ctx.Bool("enable_disk")
	}

	if len(ctx.String("enable_docker")) > 0 {
		app.opts.Docker.Enabled = ctx.Bool("enable_docker")
	}

	if len(ctx.String("enable_host")) > 0 {
		app.opts.Host.Enabled = ctx.Bool("enable_host")
	}

	if len(ctx.String("enable_load")) > 0 {
		app.opts.Load.Enabled = ctx.Bool("enable_load")
	}

	if len(ctx.String("enable_mem")) > 0 {
		app.opts.Mem.Enabled = ctx.Bool("enable_mem")
	}

	if len(ctx.String("enable_net")) > 0 {
		app.opts.Net.Enabled = ctx.Bool("enable_net")
	}

	if len(ctx.String("enable_process")) > 0 {
		app.opts.Process.Enabled = ctx.Bool("enable_process")
	}

	if ctx.Int("collector") > 0 {
		app.opts.Collector.Name = ctx.String("collector")
	}

	if app.opts.Disk.Enabled && len(ctx.StringSlice("disk_paths")) != 0 {
		diskPaths = ctx.StringSlice("disk_paths")
	}

	if app.opts.Net.Enabled && len(ctx.StringSlice("net_kinds")) != 0 {
		netKinds = ctx.StringSlice("net_kinds")
	}
}
