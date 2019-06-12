package disk

import (
	"context"
	"database/sql"

	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	ds *storage
)

type collector struct {
}

func (c *collector) PushDiskUsageStat(ctx context.Context, req *proto.DiskRequest, rsp *proto.DiskResponse) (err error) {
	err = ds.saveDiskUsageStat(req.UsageStat, req.IP, req.NodeName)
	return
}

func (c *collector) PushPartitionStat(ctx context.Context, req *proto.DiskRequest, rsp *proto.DiskResponse) (err error) {
	err = ds.savePartitionStat(req.PartitionStat, req.IP, req.NodeName)
	return
}

func (c *collector) PushIOCountersStat(ctx context.Context, req *proto.DiskRequest, rsp *proto.DiskResponse) (err error) {
	err = ds.saveIOCountersStat(req.IoCountersStat, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterDiskServiceHandler(server, new(collector))
	o = db.GetPG()
	ds = new(storage)
}
