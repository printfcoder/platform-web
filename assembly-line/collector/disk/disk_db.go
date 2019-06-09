package disk

import (
	"database/sql"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
)

type diskStorage struct {
	db *sql.DB
}

func (c *diskStorage) saveDiskUsageStat(usageStat []*proto.UsageStat, ip, nodeName string) (err error) {
	return
}

func (c *diskStorage) savePartitionStat(usageStat []*proto.PartitionStat, ip, nodeName string) (err error) {
	return
}

func (c *diskStorage) saveIOCountersStat(usageStat []*proto.IOCountersStat, ip, nodeName string) (err error) {
	return
}
