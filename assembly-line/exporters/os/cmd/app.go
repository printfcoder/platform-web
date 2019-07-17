package cmd

import (
	"runtime"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/cpu"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/disk"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/host"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/load"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/mem"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules/net"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	diskPaths  []string
	netKinds   = []string{"all"}
	configFile = "./conf/os.yml"
)

func init() {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	diskPaths = []string{path}
}

func (app *c) loadConfig(ctx *cli.Context) {
	if len(ctx.String("config_file")) > 0 {
		configFile = ctx.String("config_file")
	}

	if err := config.Load(file.NewSource(file.WithPath(configFile))); err != nil {
		panic(err)
	}

	v := config.Get("micro", "platform-web", "assembly-line", "exporters")
	log.Log(string(v.Bytes()))
	err := v.Scan(&app.opts)
	if err != nil {
		panic(err)
	}
	log.Log(app.opts.PushInterval)
	log.Log(app.opts.Collector.Name)
}

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
	if len(ctx.String("enable_cpu")) > 0 && !ctx.Bool("enable_cpu") {
		log.Logf("[Info] disabled cpu")
		app.opts.CPU.Enabled = false
	}

	if len(ctx.String("enable_disk")) > 0 && !ctx.Bool("enable_disk") {
		app.opts.Disk.Enabled = false
	}

	if len(ctx.String("enable_docker")) > 0 && !ctx.Bool("enable_docker") {
		app.opts.Docker.Enabled = false
	}

	if len(ctx.String("enable_host")) > 0 && !ctx.Bool("enable_host") {
		app.opts.Host.Enabled = false
	}

	if len(ctx.String("enable_load")) > 0 && !ctx.Bool("enable_load") {
		app.opts.Load.Enabled = false
	}

	if len(ctx.String("enable_mem")) > 0 && !ctx.Bool("enable_mem") {
		app.opts.Mem.Enabled = false
	}

	if len(ctx.String("enable_net")) > 0 && !ctx.Bool("enable_net") {
		app.opts.Net.Enabled = false
	}

	if len(ctx.String("enable_process")) > 0 && !ctx.Bool("enable_process") {
		app.opts.Process.Enabled = false
	}

	if ctx.Int("push_interval") > 0 {
		app.opts.PushInterval = time.Duration(ctx.Int("push_interval"))
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

func (app *c) loadModules(client client.Client) {
	opts := modules.Options{
		CollectorName: app.opts.Collector.Name,
		Interval:      app.opts.PushInterval,
		Client:        client,
	}

	log.Logf("[INFO] loadModules")

	// cpu
	if app.opts.CPU.Enabled {
		log.Logf("[INFO] cpu enabled")
		p := cpu.Pusher{}
		_ = p.Init(opts)

		app.modules = append(app.modules, &p)
	}

	// disk
	if app.opts.Disk.Enabled {
		p := disk.Pusher{}
		opts.DiskPaths = diskPaths
		_ = p.Init(opts)

		app.modules = append(app.modules, &p)
	}

	// host
	if app.opts.Host.Enabled {
		p := host.Pusher{}
		_ = p.Init(opts)

		app.modules = append(app.modules, &p)
	}

	// load
	if app.opts.Load.Enabled {
		p := load.Pusher{}
		_ = p.Init(opts)

		app.modules = append(app.modules, &p)
	}

	// mem
	if app.opts.Mem.Enabled {
		p := mem.Pusher{}
		_ = p.Init(opts)

		app.modules = append(app.modules, &p)
	}

	// net
	if app.opts.Net.Enabled {
		p := net.Pusher{}
		opts.NetKinds = netKinds
		_ = p.Init(opts)

		app.modules = append(app.modules, &p)
	}
}

func (app *c) run() {
	s := micro.NewService(
		micro.Name(name),
		micro.Version(version),
	)

	s.Init(micro.Action(func(ctx *cli.Context) {
		app.loadConfig(ctx)
		app.parseFlags(ctx)
		app.loadModules(s.Client())

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
