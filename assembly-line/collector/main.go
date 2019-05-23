package main

import (
	"github.com/micro-in-cn/platform-web/assembly-line/collector/cpu"
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
	)

	s.Init(
		micro.Action(func(ctx *cli.Context) {
			cpu.Init(s.Server(), ctx)
		}),
	)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
