package cpu

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	"github.com/micro-in-cn/platform-web/assembly-line/collector/util"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro/go-log"
)

type cpuStorage struct {
	db *sql.DB
}

func (c *cpuStorage) saveTimesStat(times []*cpu.TimesStat, ip, nodeName string) (err error) {

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
			util.PTimestamp(time.Timestamp), ip, nodeName, time.CPU, time.User,
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

	return
}

func (c *cpuStorage) saveInfoStat(infos []*cpu.InfoStat, ip, nodeName string) (err error) {

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
			util.PTimestamp(info.Timestamp), ip, nodeName, info.Cpu, info.VendorId,
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

func (c *cpuStorage) savePercent(percents []*cpu.Percent, ip, nodeName string) (err error) {

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
			util.PTimestamp(percent.Timestamp), ip, nodeName, percent.Percent,
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

func (c *cpuStorage) saveCounts(data []*cpu.Counts, ip, nodeName string) (err error) {
	return
}
