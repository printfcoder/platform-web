package load

import (
	"context"
	"database/sql"

	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	ds *memStorage
)

type collector struct {
}

func (c *collector) PushVirtualMemoryStat(ctx context.Context, req *proto.MemRequest, rsp *proto.MemResponse) (err error) {
	err = ds.saveVirtualMemoryStat(req.VirtualMemoryStat, req.IP, req.NodeName)
	return
}

func (c *collector) PushSwapMemoryStat(ctx context.Context, req *proto.MemRequest, rsp *proto.MemResponse) (err error) {
	err = ds.saveSwapMemoryStat(req.SwapMemoryStat, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterDiskServiceHandler(server, new(collector))
	o = db.GetPG()
	ds = new(memStorage)
}
