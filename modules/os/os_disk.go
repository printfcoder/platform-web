package os

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/util/log"
)

type DiskUsageStat struct {
	disk.UsageStat
	BaseItem
}

func (o *api) diskUsageStat(w http.ResponseWriter, r *http.Request) {
	start, end := tools.TimeFixRange(r.URL.Query().Get("startTime"), r.URL.Query().Get("endTime"),
		-time.Second*30, time.Second*30)
	ips := r.URL.Query().Get("ips")

	if ips == "" {
		err := fmt.Errorf("[diskUsageStat] err: ip is illegals")
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}

	stmt, err := db.GetPG().Prepare(`SELECT 
       time, ip, node_name, total, free, 
       used, used_percent
    FROM disk_usage_stat WHERE ip = ANY($1) AND time BETWEEN $2 AND $3 ORDER BY time DESC`)
	if err != nil {
		err = fmt.Errorf("[diskUsageStat] prepare err: %s", err)
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(pq.Array(strings.Split(ips, ",")), start, end)
	if err != nil {
		err = fmt.Errorf("[diskUsageStat] query err: %s", err)
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}

	var data []*DiskUsageStat
	for rows.Next() {
		item := &DiskUsageStat{}
		if err = rows.Scan(&item.Time, &item.IP, &item.NodeName, &item.Total, &item.Free,
			&item.Used, &item.UsedPercent);
			err != nil {
			err = fmt.Errorf("[diskUsageStat] scan err: %s", err)
			log.Logf("[ERR] %s", err)
			nosj.WriteError(w, err)
			return
		}
		data = append(data, item)
	}

	nosj.WriteJsonData(w, data)
	return
}

func (o *api) diskIOStat(w http.ResponseWriter, r *http.Request) {
	return
}

