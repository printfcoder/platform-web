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

func (s *storage) saveVirtualMemoryStat(memStats []*proto.VirtualMemoryStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO mem_virtual_memory_stat (
                                     time, ip, node_name, total, available,
                                     used, used_percent, free, active, inactive,
                                     wired, laundry, buffers, cached, writeback,
                                     dirty, writeback_tmp, shared, slab, sreclaimable,
                                     page_tables, swap_cached, commit_limit, committed_as, high_total,
                                     high_free, low_total, low_free, swap_total, swap_free,
                                     mapped, v_malloc_total, v_malloc_used, v_malloc_chunk, huge_pages_total,
                                     huge_pages_free, huge_page_size)
            VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8, $9, $10,
            $11, $12, $13, $14, $15,
            $16, $17, $18, $19, $20,
            $21, $22, $23, $24, $25,
            $26, $27, $28, $29, $30,
            $31, $32, $33, $34, $35,
	        $36, $37)`)

	if err != nil {
		log.Logf("[saveVirtualMemoryStat] db prepare error, %s", err)
		return
	}

	for _, item := range memStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.Total, item.Available,
			item.Used, item.UsedPercent, item.Free, item.Active, item.Inactive,
			item.Wired, item.Laundry, item.Buffers, item.Cached, item.Writeback,
			item.Dirty, item.WritebackTmp, item.Shared, item.Slab, item.SReclaimable,
			item.PageTables, item.SwapCached, item.CommitLimit, item.CommittedAS, item.HighTotal,
			item.HighFree, item.LowTotal, item.LowFree, item.SwapTotal, item.SwapFree,
			item.Mapped, item.VMallocTotal, item.VMallocUsed, item.VMallocChunk, item.HugePagesTotal,
			item.HugePagesFree, item.HugePageSize,
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

func (s *storage) saveSwapMemoryStat(swapMemStats []*proto.SwapMemoryStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO mem_swap_memory_stat (
                                  time, ip, node_name, total, used, 
                                  free, used_percent, sin, sout, pg_in,
                                  pg_out, pg_fault)
            VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8, $9, $10,
            $11, $12)`)

	if err != nil {
		log.Logf("[saveSwapMemoryStat] db prepare error, %s", err)
		return
	}

	for _, item := range swapMemStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.Total, item.Used,
			item.Free, item.UsedPercent, item.Sin, item.Sout, item.PgIn,
			item.PgOut, item.PgFault,
		)

		if err != nil {
			log.Logf("[saveSwapMemoryStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveSwapMemoryStat] db close error, %s", err)
		return
	}

	return
}
