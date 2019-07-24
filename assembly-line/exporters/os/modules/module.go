package modules

import (
	"sync"
	"time"
)

var (
	once         sync.Once
	ip, nodeName string
)

type Module interface {
	Init(options *Options)
	Interval() time.Duration
	Push() error
	Start() error
	String() string
}

type BaseModule struct {
	ModuleBaseOptions
	Err      chan error
	NodeName string
	IP       string
}
