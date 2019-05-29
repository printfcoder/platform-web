package cpu

import (
	"database/sql"
	"github.com/micro-in-cn/platform-web/assembly-line/collector/db"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
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
	VALUES (?, ?, ?, ?, ?, 
	        ?, ?, ?, ?, ?,
	        ?, ?, ?, ?, ?,
	        ?)`)

	if err != nil {

	}

	for _, info := range infos {
		_, err = stmt.Exec(
			info.Timestamp, ip, nodeName, info.Cpu, info.VendorId,
			info.Family, info.Model, info.Stepping, info.PhysicalId, info.CoreId,
			info.Cores, info.ModelName, info.Mhz, info.CacheSize, info.Flags,
			info.Microcode,
		)

		if err != nil {

		}
	}

	_, err = stmt.Exec()
	if err != nil {

	}

	err = stmt.Close()
	if err != nil {

	}

	return
}

func (c *cpuStorage) savePercent(data []*cpu.Percent) (err error) {
	return
}

func (c *cpuStorage) saveCounts(data []*cpu.Counts) (err error) {
	return
}
