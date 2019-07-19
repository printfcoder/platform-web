package disk

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/disk"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
)

func (p *Disk) pushPartition() (err error) {
	vv, err := disk.Partitions(true)
	if err != nil {
		return fmt.Errorf("[pushPartition] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := make([]*disk2.PartitionStat, 0, len(vv))

	for _, v := range vv {
		data = append(data, &disk2.PartitionStat{
			Timestamp:  t,
			Device:     v.Device,
			Mountpoint: v.Mountpoint,
			Fstype:     v.Fstype,
			Opts:       v.Opts,
		})
	}

	req := &disk2.DiskRequest{
		Timestamp:     t,
		IP:            p.IP,
		NodeName:      p.NodeName,
		PartitionStat: data,
	}

	_, err = p.diskClient.PushPartitionStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushPartition] push error: %s", err)
	}

	return
}
