package disk

import (
	"database/sql"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
)

type diskStorage struct {
	db *sql.DB
}

func (c *diskStorage) saveDiskUsageStat(usageStat []*disk.UsageStat, ip, nodeName string) (err error) {
	return
}

func (c *diskStorage) savePartitionStat(usageStat []*disk.PartitionStat, ip, nodeName string) (err error) {
	return
}

func (c *diskStorage) saveIOCountersStat(usageStat []*disk.IOCountersStat, ip, nodeName string) (err error) {
	return
}
