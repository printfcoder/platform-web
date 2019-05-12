package cmd

import (
	z "github.com/micro-in-cn/micro-web/internal/zap"
	"github.com/micro-in-cn/micro-web/modules"
	"github.com/micro/cli"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-web"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	Address          = ":9082"
	name             = "go.micro.web.platform"
	Version          = "1.0.1-beta"
	RootPath         = "/platform"
	apiPath          = "/api/v1"
	StaticDir        = "webapp"
	Namespace        = "go.micro.web.platform"
	registerTTL      = 30 * time.Second
	registerInterval = 10 * time.Second
	logger           = z.GetLogger()
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

	if len(ctx.GlobalString("register_ttl")) > 0 {
		registerTTL = ctx.Duration("register_ttl")
	}

	if len(ctx.GlobalString("register_interval")) > 0 {
		registerInterval = ctx.Duration("register_interval")
	}

	s := web.NewService(
		web.Name(name),
		web.Version(Version),
		web.Address(Address),
		web.RegisterTTL(registerTTL),
		web.RegisterInterval(registerInterval),
	)

	// favicon.ico
	s.HandleFunc("/favicon.ico", faviconHandler)

	// static dir
	s.Handle(RootPath+"/", http.StripPrefix(RootPath+"/", http.FileServer(http.Dir(StaticDir))))

	// init modules
	for _, m := range modules.Modules() {

		logger.Info("loading module：", zap.Any("module", m.Name()))

		m.Init(ctx)
		r := m.Path()

		for k, h := range m.Handlers() {

			route := RootPath + apiPath + r + k

			if h.IsFunc() {
				logger.Info("handler Func", zap.Any("route", route))
				s.HandleFunc(route, h.Func)
			} else {
				logger.Info("handler Handle", zap.Any("route", route))
				s.Handle(route, h.Hld)
			}
		}
	}

	if err := s.Init(
		web.Action(
			func(c *cli.Context) {
				// do something
			}),
	); err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}
