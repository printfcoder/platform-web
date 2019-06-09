package load

import (
	"database/sql"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
)

type loadStorage struct {
	db *sql.DB
}

func (c *loadStorage) saveLoadAvgStat(loadAvgStat []*proto.LoadAvgStat, ip, nodeName string) (err error) {
	return
}
