package load

import (
	"database/sql"
	"github.com/micro-in-cn/platform-web/internal/tools"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro/go-micro/util/log"
)

type storage struct {
	db *sql.DB
}

func (s *storage) saveLoadAvgStat(loadAvgStats []*proto.LoadAvgStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO load_avg_stat (
                           time, ip, node_name, load1, load5, 
                           load15)
            VALUES (
            $1, $2, $3, $4, $5,
	        $6)`)

	if err != nil {
		log.Logf("[saveLoadAvgStat] db prepare error, %s", err)
		return
	}

	for _, item := range loadAvgStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.Load1, item.Load5,
			item.Load15,
		)

		if err != nil {
			log.Logf("[saveLoadAvgStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveLoadAvgStat] db close error, %s", err)
		return
	}

	return
}
