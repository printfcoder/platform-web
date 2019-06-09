package net

import (
	"database/sql"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
)

type netStorage struct {
	db *sql.DB
}

func (c *netStorage) saveConnectionStat(connectionStat []*proto.ConnectionStat, ip, nodeName string) (err error) {
	return
}

func (c *netStorage) saveIOCountersStat(ioCountersStat []*proto.IOCountersStat, ip, nodeName string) (err error) {
	return
}
