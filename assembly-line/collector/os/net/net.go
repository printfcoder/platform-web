package net

import (
	"context"
	"database/sql"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	ns *storage
)

type collector struct {
}

func (c *collector) PushIOCountersStat(ctx context.Context, req *proto.NetRequest, rsp *proto.NetResponse) (err error) {
	err = ns.saveIOCountersStat(req.IOCountersStat, req.IP, req.NodeName)
	return
}

func (c *collector) PushConnectionStat(ctx context.Context, req *proto.NetRequest, rsp *proto.NetResponse) (err error) {
	err = ns.saveConnectionStat(req.ConnectionStat, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterNetServiceHandler(server, new(collector))
	o = db.GetPG()
	ns = new(storage)
}
