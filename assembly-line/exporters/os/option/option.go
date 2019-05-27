package option

import (
	"time"
)

type Options struct {
	PushInterval  time.Duration
	AppName       string
	AppVersion    string
	CollectorName string
	EnableCPU     bool
	EnableDisk    bool
	EnableDocker  bool
	EnableHost    bool
	EnableLoad    bool
	EnableMem     bool
	EnableNet     bool
	EnableProcess bool
}

type Option func(*Options)
