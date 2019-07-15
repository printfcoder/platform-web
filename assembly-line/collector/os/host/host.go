package host

import (
	"context"
	"database/sql"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	ds *storage
)

type collector struct {
}

func (c *collector) PushHostInfo(ctx context.Context, req *proto.HostRequest, rsp *proto.HostResponse) (err error) {
	err = ds.saveHostInfo(req.HostInfo, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterHostServiceHandler(server, new(collector))
	o = db.GetPG()
	ds = new(storage)
}
