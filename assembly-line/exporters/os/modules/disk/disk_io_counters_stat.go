package disk

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/disk"
	disk2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/disk"
)

func (d *Disk) pushIOCounters() (err error) {
	vv, err := disk.IOCounters()
	if err != nil {
		return fmt.Errorf("[pushIOCounters] get infos error: %s", err)
	}

	data := make([]*disk2.IOCountersStat, 0, len(vv))

	t := ptypes.TimestampNow()
	for _, v := range vv {
		data = append(data, &disk2.IOCountersStat{
			Timestamp:        t,
			ReadCount:        v.ReadCount,
			MergedReadCount:  v.MergedReadCount,
			WriteCount:       v.WriteCount,
			MergedWriteCount: v.MergedWriteCount,
			ReadBytes:        v.ReadBytes,
			WriteBytes:       v.WriteBytes,
			ReadTime:         v.ReadTime,
			WriteTime:        v.WriteTime,
			IopsInProgress:   v.IopsInProgress,
			IoTime:           v.IoTime,
			WeightedIO:       v.WeightedIO,
			Name:             v.Name,
			SerialNumber:     v.SerialNumber,
			Label:            v.Label,
		})
	}

	req := &disk2.DiskRequest{
		Timestamp:      t,
		IP:             d.IP,
		NodeName:       d.NodeName,
		IoCountersStat: data,
	}

	_, err = d.diskClient.PushIOCountersStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushIOCounters] push error: %s", err)
	}

	return
}
