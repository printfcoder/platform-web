package load

import (
	"database/sql"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

type memStorage struct {
	db *sql.DB
}

func (c *memStorage) saveVirtualMemoryStat(usageStat []*proto.VirtualMemoryStat, ip, nodeName string) (err error) {
	return
}

func (c *memStorage) saveSwapMemoryStat(usageStat []*proto.SwapMemoryStat, ip, nodeName string) (err error) {
	return
}
