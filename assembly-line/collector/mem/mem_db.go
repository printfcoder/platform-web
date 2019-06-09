package mem

import (
	"database/sql"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

type memStorage struct {
	db *sql.DB
}

func (c *memStorage) saveVirtualMemoryStat(memStat []*proto.VirtualMemoryStat, ip, nodeName string) (err error) {
	return
}

func (c *memStorage) saveSwapMemoryStat(swapMemStat []*proto.SwapMemoryStat, ip, nodeName string) (err error) {
	return
}
