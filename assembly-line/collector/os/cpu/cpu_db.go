package cpu

import (
	"database/sql"

	"github.com/lib/pq"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro/go-log"
)

type storage struct {
	db *sql.DB
}

func (s *storage) saveTimesStat(times []*proto.TimesStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO cpu_times (
                       time, ip, node_name, cpu, x_user,
                       system, idle, nice, iowait, irq, 
                       softirq, steal, guest, guest_nice)
	VALUES ($1, $2, $3, $4, $5, 
	        $6, $7, $8, $9, $10,
	        $11, $12, $13, $14)`)

	if err != nil {
		log.Logf("[saveTimesStat] db prepare error, %s", err)
		return
	}

	for _, time := range times {
		_, err = stmt.Exec(
			tools.PTimestamp(time.Timestamp), ip, nodeName, time.CPU, time.User,
			time.System, time.Idle, time.Nice, time.Iowait, time.Irq,
			time.Softirq, time.Steal, time.Guest, time.GuestNice,
		)

		if err != nil {
			log.Logf("[saveTimesStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveTimesStat] db close error, %s", err)
		return
	}

	return
}

func (s *storage) saveInfoStat(infos []*proto.InfoStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO cpu_info (
                      time, ip, node_name, cpu, vendor_id, 
                      family, model, stepping, physical_id, core_id, 
                      cores, model_name, mhz, cache_size, flags, 
                      microcode)
	VALUES ($1, $2, $3, $4, $5, 
	        $6, $7, $8, $9, $10,
	        $11, $12, $13, $14, $15,
	        $16)`)

	if err != nil {
		log.Logf("[saveInfoStat] db prepare error, %s", err)
		return
	}

	for _, info := range infos {
		_, err = stmt.Exec(
			tools.PTimestamp(info.Timestamp), ip, nodeName, info.Cpu, info.VendorId,
			info.Family, info.Model, info.Stepping, info.PhysicalId, info.CoreId,
			info.Cores, info.ModelName, info.Mhz, info.CacheSize, pq.StringArray(info.Flags),
			info.Microcode,
		)

		if err != nil {
			log.Logf("[saveInfoStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveInfoStat] db close error, %s", err)
		return
	}

	return
}

func (s *storage) savePercent(percents []*proto.Percent, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO cpu_percent (
                         time, ip, node_name, percent
                         )
	VALUES ($1, $2, $3, $4)`)

	if err != nil {
		log.Logf("[savePercent] db prepare error, %s", err)
		return
	}

	for _, percent := range percents {
		_, err = stmt.Exec(
			tools.PTimestamp(percent.Timestamp), ip, nodeName, percent.Percent,
		)

		if err != nil {
			log.Logf("[savePercent] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[savePercent] db close error, %s", err)
		return
	}

	return
}

func (s *storage) saveCounts(data []*proto.Counts, ip, nodeName string) (err error) {
	return
}
