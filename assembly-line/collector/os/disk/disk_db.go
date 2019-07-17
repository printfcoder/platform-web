package disk

import (
	"database/sql"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro/go-log"
)

type storage struct {
	db *sql.DB
}

func (s *storage) saveDiskUsageStat(usageStats []*proto.UsageStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO disk_usage_stat (
                             time, ip, node_name, path, fs_type, 
                             total, free, used, used_percent, inodes_total, 
                             inodes_used, inodes_free, inodes_used_percent)
	VALUES ($1, $2, $3, $4, $5,
	        $6, $7, $8, $9, $10, 
	        $11, $12, $13
	        )`)

	if err != nil {
		log.Logf("[saveDiskUsageStat] db prepare error, %s", err)
		return
	}

	for _, usageStat := range usageStats {
		_, err = stmt.Exec(
			tools.PTimestamp(usageStat.Timestamp), ip, nodeName, usageStat.Fstype, usageStat.Path,
			usageStat.Total, usageStat.Free, usageStat.Used, usageStat.UsedPercent, usageStat.InodesTotal,
			usageStat.InodesUsed, usageStat.InodesFree, usageStat.InodesUsedPercent,
		)

		if err != nil {
			log.Logf("[saveDiskUsageStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveDiskUsageStat] db close error, %s", err)
		return
	}

	return
}

func (s *storage) savePartitionStat(partitionStats []*proto.PartitionStat, ip, nodeName string) (err error) {
	// do nothing
	return
}

func (s *storage) saveIOCountersStat(ioCountersStats []*proto.IOCountersStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO disk_io_counters_stat (
                                   time, ip, node_name, read_count, merged_read_count,
                                   write_count, merged_write_count, read_bytes, write_bytes, read_time,
                                   write_time, iops_in_progress, io_time, weighted_io, name,
                                   serial_number, label)
    VALUES  ($1, $2, $3, $4, $5,
	        $6, $7, $8, $9, $10, 
	        $11, $12, $13, $14, $15,
            $16, $17
	        )`)

	if err != nil {
		log.Logf("[saveIOCountersStat] db prepare error, %s", err)
		return
	}

	for _, item := range ioCountersStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.ReadCount, item.MergedReadCount,
			item.WriteCount, item.MergedWriteCount, item.ReadBytes, item.WriteBytes, item.ReadTime,
			item.WriteTime, item.IopsInProgress, item.IoTime, item.WeightedIO, item.Name,
			item.SerialNumber, item.Label,
		)

		if err != nil {
			log.Logf("[saveIOCountersStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveIOCountersStat] db close error, %s", err)
		return
	}

	return
}
