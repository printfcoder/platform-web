package os

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/util/log"
)

type BaseItem struct {
	IP       string    `json:"ip"`
	NodeName string    `json:"nodeName"`
	Time     time.Time `json:"time"`
}

type CPUInfo struct {
	cpu.InfoStat
	BaseItem
}

type CPUTimes struct {
	cpu.TimesStat
	BaseItem
}

type CPUPercent struct {
	cpu.Percent
	BaseItem
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
	ips := r.URL.Query().Get("ips")

	if ips == "" {
		err := fmt.Errorf("[cpuPercent] err: ip is illegals")
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}

	stmt, err := db.GetPG().Prepare(`SELECT time, ip, node_name, percent FROM cpu_percent WHERE ip = ANY($1) AND time BETWEEN $2 AND $3 ORDER BY time DESC`)
	if err != nil {
		err = fmt.Errorf("[cpuPercent] prepare err: %s", err)
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(pq.Array(strings.Split(ips, ",")), start, end)
	if err != nil {
		err = fmt.Errorf("[cpuPercent] query err: %s", err)
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}

	var data []*CPUPercent
	for rows.Next() {
		item := &CPUPercent{}
		if err = rows.Scan(&item.Time, &item.IP, &item.NodeName, &item.Percent.Percent);
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
	start, end := tools.TimeFixRange(r.URL.Query().Get("startTime"), r.URL.Query().Get("endTime"),
		-time.Second*30, time.Second*30)
	ips := r.URL.Query().Get("ips")

	if ips == "" {
		err := fmt.Errorf("[cpuTimes] err: ip is illegal")
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}

	stmt, err := db.GetPG().Prepare(`SELECT
       time, ip, node_name, cpu, x_user, 
       system, idle, nice, iowait, irq, 
       softirq, steal, guest, guest_nice
       FROM cpu_times WHERE ip = ANY($1) AND time BETWEEN $2 AND $3 ORDER BY time DESC`)
	if err != nil {
		err = fmt.Errorf("[cpuTimes] prepare err: %s", err)
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(pq.Array(strings.Split(ips, ",")), start, end)
	if err != nil {
		err = fmt.Errorf("[cpuTimes] query err: %s", err)
		log.Log(err)
		nosj.WriteError(w, err)
		return
	}

	var data []*CPUTimes
	for rows.Next() {
		item := &CPUTimes{}
		if err = rows.Scan(&item.Time, &item.IP, &item.NodeName, &item.CPU, &item.User,
			&item.System, &item.Idle, &item.Nice, &item.Iowait, &item.Irq,
			&item.Softirq, &item.Steal, &item.Guest, &item.GuestNice);
			err != nil {
			err = fmt.Errorf("[cpuTimes] scan err: %s", err)
			log.Log(err)
			nosj.WriteError(w, err)
			return
		}
		data = append(data, item)
	}

	nosj.WriteJsonData(w, data)
	return
}
