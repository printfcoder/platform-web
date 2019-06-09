package load

import (
	"context"
	"database/sql"

	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	ls *loadStorage
)

type collector struct {
}

func (c *collector) PushLoadAvgStat(ctx context.Context, req *proto.LoadRequest, rsp *proto.LoadResponse) (err error) {
	err = ls.saveLoadAvgStat(req.LoadAvgStat, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterLoadServiceHandler(server, new(collector))
	o = db.GetPG()
	ls = new(loadStorage)
}
