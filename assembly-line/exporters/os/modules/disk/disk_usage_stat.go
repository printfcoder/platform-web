package disk

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/disk"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
)

func (d *Disk) pushUsage() (err error) {
	data := make([]*disk2.UsageStat, 0, len(d.opts.Paths))
	t := ptypes.TimestampNow()

	for _, path := range d.opts.Paths {
		v, err := disk.Usage(path)
		if err != nil {
			return fmt.Errorf("[pushUsage] get usage error: %s", err)
		}

		data = append(data, &disk2.UsageStat{
			Timestamp:         t,
			Path:              v.Path,
			Fstype:            v.Fstype,
			Total:             v.Total,
			Free:              v.Free,
			Used:              v.Used,
			UsedPercent:       v.UsedPercent,
			InodesTotal:       v.InodesTotal,
			InodesUsed:        v.InodesUsed,
			InodesFree:        v.InodesFree,
			InodesUsedPercent: v.InodesUsedPercent,
		})
	}

	req := &disk2.DiskRequest{
		Timestamp: t,
		IP:        d.opts.IP,
		NodeName:  d.opts.NodeName,
		UsageStat: data,
	}

	_, err = d.diskClient.PushDiskUsageStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushUsage] push error: %s", err)
	}

	return
}
