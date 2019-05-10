package cmd

import (
	"github.com/micro-in-cn/micro-web/modules"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-web"
	"net/http"
	"strings"
)

var (
	Address   = ":8082"
	name      = "go.micro.web.dashboard"
	Version   = "1.0.1"
	Path      = "/"
	Namespace = "go.micro.web.v2"
)

// Init app
func Init(ops ...modules.Option) {

	app := cmd.App()
	app.Commands = Commands(ops...)
	app.Action = func(c *cli.Context) {
		run(c)
	}

	cmd.Init(
		cmd.Name(name),
	)
}

// Setup sets up the web app

func run(ctx *cli.Context, srvOpts ...modules.Option) {

	if len(ctx.GlobalString("server_name")) > 0 {
		name = ctx.GlobalString("server_name")
	}

	if len(ctx.GlobalString("server_version")) > 0 {
		Version = ctx.GlobalString("server_version")
	}

	if len(ctx.GlobalString("address")) > 0 {
		Version = ctx.GlobalString("address")
	}

	if len(ctx.String("address")) > 0 {
		Address = ctx.String("address")
	}

	s := web.NewService(
		web.Name(name),
		web.Version(Version),
		web.Address(Address),
	)

	s.HandleFunc("/favicon.ico", faviconHandler)

	// Init plugins
	for _, p := range modules.Modules() {

		p.Init(ctx)
		r := p.Path()

		for k, h := range p.Handlers() {
			if strings.LastIndex(r, "/") != len(r)-1 && strings.Index(k, "/") != 0 {
				r += "/"
			}
			if h.IsFunc() {
				s.HandleFunc(r+k, h.Func)
			} else {
				s.Handle(r+k, h.Hld)
			}
		}
	}

	// 初始化服务
	if err := s.Init(
		web.Action(
			func(c *cli.Context) {
				// do something
			}),
	); err != nil {
		panic(err)
	}

	// 运行服务
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func Commands(options ...modules.Option) []cli.Command {
	command := cli.Command{
		Name:  "web",
		Usage: "Run the web dashboard",
		Action: func(c *cli.Context) {
			run(c, options...)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "address",
				Usage:  "Set the web UI address e.g 0.0.0.0:8082",
				EnvVar: "MICRO_WEB_ADDRESS",
			},
			cli.StringFlag{
				Name:   "namespace",
				Usage:  "Set the namespace used by the Web proxy e.g. com.example.web",
				EnvVar: "MICRO_WEB_NAMESPACE",
			},
			cli.StringFlag{
				Name:   "static_dir",
				Usage:  "Set the static dir of micro web",
				EnvVar: "MICRO_WEB_STATIC_DIR",
			},
		},
	}

	return []cli.Command{command}
}
