package os

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/util/log"
)

type LoadAvgStat struct {
	load.LoadAvgStat
	BaseItem
}

func (o *api) loadAvgStat(w http.ResponseWriter, r *http.Request) {
	start, end := tools.TimeFixRange(r.URL.Query().Get("startTime"), r.URL.Query().Get("endTime"),
		-time.Second*30, time.Second*30)
	ips := r.URL.Query().Get("ips")

	if ips == "" {
		err := fmt.Errorf("[loadAvgStat] err: ip is illegals")
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}

	stmt, err := db.GetPG().Prepare(`SELECT time, ip, node_name, load1, load5, load15
    FROM load_avg_stat WHERE ip = ANY($1) AND time BETWEEN $2 AND $3`)
	if err != nil {
		err = fmt.Errorf("[loadAvgStat] prepare err: %s", err)
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(pq.Array(strings.Split(ips, ",")), start, end)
	if err != nil {
		err = fmt.Errorf("[loadAvgStat] query err: %s", err)
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}

	var data []*LoadAvgStat
	for rows.Next() {
		item := &LoadAvgStat{}
		if err = rows.Scan(&item.Time, &item.IP, &item.NodeName, &item.Load1, &item.Load5,
			&item.Load15);
			err != nil {
			err = fmt.Errorf("[loadAvgStat] scan err: %s", err)
			log.Logf("[ERR] %s", err)
			nosj.WriteError(w, err)
			return
		}
		data = append(data, item)
	}

	nosj.WriteJsonData(w, data)
	return
}
