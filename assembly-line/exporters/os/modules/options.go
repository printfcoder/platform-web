package modules

import (
	"time"

	"github.com/micro/go-micro/client"
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
	IP         string
	NodeName   string
}

type ModuleBaseOptions struct {
	// make modules use Collector easily
	Collector *Collector    `json:"collector"`
	Enabled   bool          `json:"enable"`
	Interval  time.Duration `json:"interval"`
	IP        string        `json:"ip"`
	NodeName  string        `json:"nodeName"`
}

type Collector struct {
	Name   string `json:"name"`
	Client client.Client
}

type CPUOptions struct {
	*ModuleBaseOptions
}

type DiskOptions struct {
	*ModuleBaseOptions
	Paths []string `json:"paths"`
}

type DockerOptions struct {
	*ModuleBaseOptions
}

type HostOptions struct {
	*ModuleBaseOptions
}

type LoadOptions struct {
	*ModuleBaseOptions
}

type MemOptions struct {
	*ModuleBaseOptions
}

type NetOptions struct {
	*ModuleBaseOptions
	Kinds  []string `json:"kinds"`
	Ifaces []string `json:"ifaces"`
}

type ProcessOptions struct {
	*ModuleBaseOptions
}

type Option func(*Options)
