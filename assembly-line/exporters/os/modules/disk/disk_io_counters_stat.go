package disk

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/disk"
	cpu2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
)

func (p *Pusher) pushInfo() (err error) {

	vv, err := disk.IOCounters()
	if err != nil {
		return fmt.Errorf("[pushInfo] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	for _, v := range vv {
		req := &cpu2.CPURequest{
			Timestamp: t,
			IP:        p.IP,
			NodeName:  p.NodeName,
			InfoStat: []*cpu2.InfoStat{{
				Timestamp:  t,
				Cpu:        v.CPU,
				VendorId:   v.VendorID,
				Family:     v.Family,
				Model:      v.Model,
				Stepping:   v.Stepping,
				PhysicalId: v.PhysicalID,
				CoreId:     v.CoreID,
				Cores:      v.Cores,
				ModelName:  v.ModelName,
				Mhz:        v.Mhz,
				CacheSize:  v.CacheSize,
				Flags:      v.Flags,
				Microcode:  v.Microcode,
			}},
		}

		_, err = p.diskClient.PushIOCountersStat(context.Background(), req)
		if err != nil {
			return fmt.Errorf("[pushInfo] push error: %s", err)
		}
	}

	return
}
