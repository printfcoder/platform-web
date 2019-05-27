package modules

type Pusher interface {
	Init(opts Options) error
	Push() error
}

type BasePusher struct {
	CollectorName string
}

type Options struct {
	CollectorName string
}

type Option func(*Options)
