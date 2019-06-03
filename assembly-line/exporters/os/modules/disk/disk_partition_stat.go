package disk

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/disk"
	cpu2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
)

func (p *Pusher) pushPercent() (err error) {

	vv, err := disk.Partitions(true)
	if err != nil {
		return fmt.Errorf("[pushPercent] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	for _, v := range vv {
		req := &cpu2.CPURequest{
			Timestamp: t,
			IP:        p.IP,
			NodeName:  p.NodeName,
			Percent: []*cpu2.Percent{{
				Timestamp: t,
				Percent:   v,
			}},
		}

		_, err = p.diskClient.PushDiskCounts(context.Background(), req)
		if err != nil {
			return fmt.Errorf("[pushPercent] push error: %s", err)
		}
	}

	return
}
