package cpu

import (
	"context"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

type cpuCollector struct {
}

func (c *cpuCollector) PushCPUTimesStat(ctx context.Context, req *cpu.CPURequest, rsp *cpu.CPUResponse) (err error) {
	return
}

func (c *cpuCollector) PushCPUInfoStat(ctx context.Context, req *cpu.CPURequest, rsp *cpu.CPUResponse) (err error) {

	return
}

func (c *cpuCollector) PushCPUPercent(ctx context.Context, req *cpu.CPURequest, rsp *cpu.CPUResponse) (err error) {
	return
}

func (c *cpuCollector) PushCPUCounts(ctx context.Context, req *cpu.CPURequest, rsp *cpu.CPUResponse) (err error) {
	return
}

func Init(server server.Server, ctx *cli.Context) {
	cpu.RegisterCPUServiceHandler(server, new(cpuCollector))
}

