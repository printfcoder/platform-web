package option

import "time"

type Options struct {
	PushInterval  time.Duration
	AppName       string
	AppVersion    string
	CollectorName string
}

type Option func(*Options)
