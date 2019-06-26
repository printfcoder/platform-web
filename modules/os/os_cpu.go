package os

import (
	"fmt"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"time"
)

type CPUInfo struct {
	cpu.InfoStat
	IP       string `json:"ip"`
	NodeName string `json:"nodeName"`
}

func (o *api) cpuInfos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetPG().Query(`SELECT DISTINCT ip, node_name, cpu, vendor_id, family,
                model, stepping, physical_id, core_id, cores,
                model_name, mhz, cache_size, microcode 
    FROM cpu_info`)
	if err != nil {
		err = fmt.Errorf("[cpuInfos] query err: %s", err)
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}
	defer rows.Close()

	var data []*CPUInfo
	for rows.Next() {
		item := &CPUInfo{}
		if err = rows.Scan(&item.IP, &item.NodeName, &item.Cpu, &item.VendorId, &item.Family,
			&item.Model, &item.Stepping, &item.PhysicalId, &item.CoreId, &item.Cores,
			&item.ModelName, &item.Mhz, &item.CacheSize, &item.Microcode);
			err != nil {
			err = fmt.Errorf("[cpuInfos] scan err: %s", err)
			log.Log(err)
			nosj.WriteError(w, err)
			return
		}
		data = append(data, item)
	}

	nosj.WriteJsonData(w, data)
	return
}

func (o *api) cpuPercent(w http.ResponseWriter, r *http.Request) {
	start, end := tools.TimeFixRange(r.URL.Query().Get("startTime"), r.URL.Query().Get("endTime"),
		-time.Second*30, time.Second*30)
	ip := r.URL.Query().Get("ip")

	if ip == "" {
		err := fmt.Errorf("[cpuPercent] err: ip is illegal")
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}

	rows, err := db.GetPG().Query(`SELECT time, node_name, percent FROM cpu_percent WHERE ip = ? AND time BETWEEN ? AND ?`,
		ip, start, end)
	if err != nil {
		err = fmt.Errorf("[cpuPercent] query err: %s", err)
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}

	var data []*cpu.Percent
	for rows.Next() {
		item := &cpu.Percent{}
		if err = rows.Scan(&item.Timestamp, &item.Percent);
			err != nil {
			err = fmt.Errorf("[cpuPercent] scan err: %s", err)
			log.Log(err)
			nosj.WriteError(w, err)
			return
		}
		data = append(data, item)
	}

	nosj.WriteJsonData(w, data)
	return
}

func (o *api) cpuTimes(w http.ResponseWriter, r *http.Request) {

	return
}
