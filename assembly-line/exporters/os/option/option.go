package option

import (
	"time"
)

type Options struct {
	PushInterval  time.Duration   `json:"push-interval"`
	AppName       string          `json:"app-name"`
	AppVersion    string          `json:"app-version"`
	CollectorName string          `json:"collector-name"`
	CPU           *CPUOptions     `json:"cpu"`
	Disk          *DiskOptions    `json:"disk"`
	Docker        *DockerOptions  `json:"docker"`
	Host          *HostOptions    `json:"host"`
	Load          *LoadOptions    `json:"load"`
	Mem           *MemOptions     `json:"mem"`
	Net           *NetOptions     `json:"net"`
	Process       *ProcessOptions `json:"process"`
}

type CPUOptions struct {
	Enabled bool `json:"enable"`
}

type DiskOptions struct {
	Enabled bool     `json:"enable"`
	Paths   []string `json:"paths"`
}

type DockerOptions struct {
	Enabled bool `json:"enable"`
}

type HostOptions struct {
	Enabled bool `json:"enable"`
}

type LoadOptions struct {
	Enabled bool `json:"enable"`
}

type MemOptions struct {
	Enabled bool `json:"enable"`
}

type NetOptions struct {
	Enabled bool     `json:"enable"`
	Kinds   []string `json:"kinds"`
}

type ProcessOptions struct {
	Enabled bool `json:"enable"`
}

type Option func(*Options)
