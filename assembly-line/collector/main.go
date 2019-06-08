package main

import (
	"github.com/micro-in-cn/platform-web/assembly-line/collector/cpu"
	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	"github.com/micro-in-cn/platform-web/assembly-line/collector/disk"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

var (
	name    = "go.micro.srv.platform_collector"
	version = "v1"
)

func main() {
	s := micro.NewService(
		micro.Name(name),
		micro.Version(version),
		micro.Flags(
			cli.StringFlag{
				Name:   "pg_config_file",
				Usage:  "enables exporter to export docker info",
				EnvVar: "MICRO_WEB_PLATFORM_PG_CONFIG_FILE",
			},
		),
	)

	s.Init(
		micro.Action(func(ctx *cli.Context) {
			db.Init(ctx)
			cpu.Init(s.Server(), ctx)
			disk.Init(s.Server(), ctx)
		}),
	)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
