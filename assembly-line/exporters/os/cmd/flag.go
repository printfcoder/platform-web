package cmd

import "github.com/micro/cli"

func (app *c) advFlags() {
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "config_file",
			Usage:  "path to config file, but the configs in file are less weight than those passed on command",
			EnvVar: "MICRO_WEB_PLATFORM_CONFIG_FILE",
		},
	)
}
