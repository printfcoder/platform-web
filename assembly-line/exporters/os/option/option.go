package option

import (
	"github.com/micro/go-micro/client"
	"time"
)

type Options struct {
	AppName    string          `json:"app-name"`
	AppVersion string          `json:"app-version"`
	Collector  *Collector      `json:"collector"`
	CPU        *CPUOptions     `json:"cpu"`
	Disk       *DiskOptions    `json:"disk"`
	Docker     *DockerOptions  `json:"docker"`
	Host       *HostOptions    `json:"host"`
	Load       *LoadOptions    `json:"load"`
	Mem        *MemOptions     `json:"mem"`
	Net        *NetOptions     `json:"net"`
	Process    *ProcessOptions `json:"process"`
}

type Collector struct {
	Name   string `json:"name"`
	Client client.Client
}

type CPUOptions struct {
	Enabled  bool          `json:"enable"`
	Interval time.Duration `json:"interval"`
}

type DiskOptions struct {
	Enabled  bool          `json:"enable"`
	Paths    []string      `json:"paths"`
	Interval time.Duration `json:"interval"`
}

type DockerOptions struct {
	Enabled  bool          `json:"enable"`
	Interval time.Duration `json:"interval"`
}

type HostOptions struct {
	Enabled  bool          `json:"enable"`
	Interval time.Duration `json:"interval"`
}

type LoadOptions struct {
	Enabled  bool          `json:"enable"`
	Interval time.Duration `json:"interval"`
}

type MemOptions struct {
	Enabled  bool          `json:"enable"`
	Interval time.Duration `json:"interval"`
}

type NetOptions struct {
	Enabled  bool          `json:"enable"`
	Kinds    []string      `json:"kinds"`
	Interval time.Duration `json:"interval"`
}

type ProcessOptions struct {
	Enabled  bool          `json:"enable"`
	Interval time.Duration `json:"interval"`
}

type Option func(*Options)
