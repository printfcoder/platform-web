package net

import (
	"database/sql"
	"github.com/lib/pq"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro/go-log"
)

type storage struct {
	db *sql.DB
}

func (s *storage) saveConnectionStat(connectionStats []*proto.ConnectionStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO net_connection_stat (
                                 time, ip, node_name, fd, family, 
                                 type, l_addr_ip, l_addr_port, r_addr_ip, r_addr_port,
                                 status, uids, Pid)
            VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8, $9, $10,
            $11, $12, $13)`)

	if err != nil {
		log.Logf("[saveConnectionStat] db prepare error, %s", err)
		return
	}

	for _, item := range connectionStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.Fd, item.Family,
			item.Type, item.Laddr.IP, item.Laddr.Port, item.Raddr.IP, item.Raddr.Port,
			item.Status, pq.Int64Array(tools.Int32ArrayTo64(item.Uids)), item.Pid,
		)

		if err != nil {
			log.Logf("[saveConnectionStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveConnectionStat] db close error, %s", err)
		return
	}

	return
}

func (s *storage) saveIOCountersStat(ioCountersStats []*proto.IOCountersStat, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO net_io_counters_stat (
                                  time, ip, node_name, name, bytes_sent,
                                  bytes_recv, packets_sent, packets_recv, err_in, err_out,
                                  drop_in, drop_out, fifo_in, fifo_out)
            VALUES (
            $1, $2, $3, $4, $5,
            $6, $7, $8, $9, $10,
            $11, $12, $13, $14)`)

	if err != nil {
		log.Logf("[saveIOCountersStat] db prepare error, %s", err)
		return
	}

	for _, item := range ioCountersStats {
		_, err = stmt.Exec(
			tools.PTimestamp(item.Timestamp), ip, nodeName, item.Name, item.BytesSent,
			item.BytesRecv, item.PacketsSent, item.PacketsRecv, item.Errin, item.Errout,
			item.Dropin, item.Dropout, item.Fifoin, item.Fifoout,
		)

		if err != nil {
			log.Logf("[saveIOCountersStat] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveIOCountersStat] db close error, %s", err)
		return
	}

	return
}
