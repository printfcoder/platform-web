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

func (c *cpuStorage) saveTimesStat(data []*cpu.TimesStat) (err error) {
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

func (c *cpuStorage) savePercent(data []*cpu.Percent) (err error) {
	return
}

func (c *cpuStorage) saveCounts(data []*cpu.Counts) (err error) {
	return
}
