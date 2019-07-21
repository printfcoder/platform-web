package os

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/util/log"
)

type MemPercent struct {
	mem.MemoryStat
	BaseItem
}

func (o *api) memPercent(w http.ResponseWriter, r *http.Request) {
	start, end := tools.TimeFixRange(r.URL.Query().Get("startTime"), r.URL.Query().Get("endTime"),
		-time.Second*30, time.Second*30)
	ips := r.URL.Query().Get("ips")

	if ips == "" {
		err := fmt.Errorf("[memPercent] err: ip is illegals")
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}

	stmt, err := db.GetPG().Prepare(`SELECT time, ip, node_name, active_bytes, compressed_bytes,
       inactive_bytes,  wired_bytes, free_bytes, swapped_in_bytes_total, swapped_out_bytes_total,
       total_bytes
    FROM memory_stat WHERE ip = ANY($1) AND time BETWEEN $2 AND $3`)
	if err != nil {
		err = fmt.Errorf("[memPercent] prepare err: %s", err)
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(pq.Array(strings.Split(ips, ",")), start, end)
	if err != nil {
		err = fmt.Errorf("[memPercent] query err: %s", err)
		log.Logf("[ERR] %s", err)
		nosj.WriteError(w, err)
		return
	}

	var data []*MemPercent
	for rows.Next() {
		item := &MemPercent{}
		if err = rows.Scan(&item.Time, &item.IP, &item.NodeName, &item.MemoryStat.ActiveBytes, &item.MemoryStat.CompressedBytes,
			&item.MemoryStat.InactiveBytes, &item.MemoryStat.WiredBytes, &item.MemoryStat.FreeBytes, &item.MemoryStat.SwappedInBytesTotal, &item.MemoryStat.SwappedOutBytesTotal,
			&item.MemoryStat.TotalBytes);
			err != nil {
			err = fmt.Errorf("[memPercent] scan err: %s", err)
			log.Logf("[ERR] %s", err)
			nosj.WriteError(w, err)
			return
		}
		data = append(data, item)
	}

	nosj.WriteJsonData(w, data)
	return
}
