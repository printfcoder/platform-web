package host

import (
	"database/sql"
	"github.com/micro-in-cn/platform-web/internal/tools"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro/go-log"
)

type storage struct {
	db *sql.DB
}

func (s *storage) saveHostInfo(infos []*proto.HostInfo, ip, nodeName string) (err error) {
	o := db.GetPG()

	// transaction is no need here

	stmt, err := o.Prepare(`INSERT INTO host_info (
                       time, ip, node_name, hostname, up_time,
                       boot_Time, procs, os, platform, platform_family,
                       platform_version, kernel_version, virtualization_system, virtualization_role, host_id) 
	VALUES ($1, $2, $3, $4, $5,
	        $6, $7, $8, $9, $10, 
	        $11, $12, $13, $14, $15
	        )`)

	if err != nil {
		log.Logf("[saveHostInfo] db prepare error, %s", err)
		return
	}

	for _, info := range infos {
		_, err = stmt.Exec(
			tools.PTimestamp(info.Timestamp), ip, nodeName, info.Hostname, info.Uptime,
			info.BootTime, info.Procs, info.OS, info.Platform, info.PlatformFamily,
			info.PlatformVersion, info.KernelVersion, info.VirtualizationSystem, info.VirtualizationRole, info.HostID,
		)

		if err != nil {
			log.Logf("[saveHostInfo] db prepare exec error, %s", err)
			return
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Logf("[saveHostInfo] db close error, %s", err)
		return
	}

	return
}
