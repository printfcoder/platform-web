package mem

import (
	"database/sql"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro/go-log"
)

type storage struct {
	db *sql.DB
}

func (s *storage) saveMemStat(memStats []*proto.MemoryStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO memory_stat (
                         time, ip, node_name, active_bytes, compressed_bytes, 
                         inactive_bytes, wired_bytes, free_bytes, swapped_in_bytes_total, swapped_out_bytes_total, 
                         total_bytes)
            VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8, $9, $10,
            $11)`)

	if err != nil {
		log.Logf("[saveMemStat] db prepare error, %s", err)
		return
	}

	for _, item := range memStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.ActiveBytes, item.CompressedBytes,
			item.InactiveBytes, item.WiredBytes, item.FreeBytes, item.SwappedInBytesTotal, item.SwappedOutBytesTotal,
			item.TotalBytes,
		)

		if err != nil {
			log.Logf("[saveMemStat] db exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveMemStat] db close error, %s", err)
		return
	}

	return
}
