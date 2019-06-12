package cpu

import (
	"context"
	"database/sql"
	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	cs *storage
)

type cpuCollector struct {
}

func (c *cpuCollector) PushCPUTimesStat(ctx context.Context, req *proto.CPURequest, rsp *proto.CPUResponse) (err error) {
	err = cs.saveTimesStat(req.TimesStat, req.IP, req.NodeName)
	return
}

func (c *cpuCollector) PushCPUInfoStat(ctx context.Context, req *proto.CPURequest, rsp *proto.CPUResponse) (err error) {
	err = cs.saveInfoStat(req.InfoStat, req.IP, req.NodeName)
	return
}

func (c *cpuCollector) PushCPUPercent(ctx context.Context, req *proto.CPURequest, rsp *proto.CPUResponse) (err error) {
	err = cs.savePercent(req.Percent, req.IP, req.NodeName)
	return
}

func (c *cpuCollector) PushCPUCounts(ctx context.Context, req *proto.CPURequest, rsp *proto.CPUResponse) (err error) {
	err = cs.saveCounts(req.Counts, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterCPUServiceHandler(server, new(cpuCollector))
	o = db.GetPG()
	cs = new(storage)
}
