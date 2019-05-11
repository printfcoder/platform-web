package cmd

import (
	z "github.com/micro-in-cn/micro-web/internal/zap"
	"github.com/micro-in-cn/micro-web/modules"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-web"
	"go.uber.org/zap"
	"net/http"
)

var (
	Address   = ":8082"
	name      = "go.micro.web.dashboard"
	Version   = "1.0.1"
	RootPath  = "/v1/api"
	Namespace = "go.micro.web.v2"
	logger    = z.GetLogger()
)

// Init app
func Init(ops ...modules.Option) {

	app := cmd.App()
	app.Flags = append(app.Flags,
		cli.StringFlag{
			Name:   "root_path",
			Usage:  "Set the root path of micro web",
			EnvVar: "MICRO_WEB_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "static_dir",
			Usage:  "Set the static dir of micro web",
			EnvVar: "MICRO_WEB_STATIC_DIR",
		},
	)

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

	if len(ctx.String("root_path")) > 0 {
		RootPath = ctx.String("root_path")
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

	logger.Info("loading modules", zap.Any("modules", modules.Modules()))

	// Init modules
	for _, p := range modules.Modules() {

		p.Init(ctx)
		r := p.Path()

		if r == "/" {
			r = ""
		}

		for k, h := range p.Handlers() {

			route := RootPath + r + k

			if h.IsFunc() {
				logger.Info("handler Func", zap.Any("route", route), zap.Any("func", h.Func))
				s.HandleFunc(route, h.Func)
			} else {
				logger.Info("handler Handle", zap.Any("route", route), zap.Any("handle", h.Hld))
				s.Handle(route, h.Hld)
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
