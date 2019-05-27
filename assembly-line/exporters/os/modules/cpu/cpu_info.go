package cpu

import (
	"context"
	"fmt"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	cpu2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
	"time"
)

func (p *Pusher) pushInfo() (err error) {

	vv, err := cpu.Info()
	if err != nil {
		return fmt.Errorf("[pushInfo] get infos error: %s", err)
	}

	for _, v := range vv {
		req := &cpu2.CPURequest{
			Timestamp: time.Now().UnixNano(),
			IP:        "",
			NodeName:  "",
			InfoStat: &cpu2.InfoStat{
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
			},
		}

		_, err = p.cpuClient.PushCPUInfoStat(context.Background(), req)
		if err != nil {
			return fmt.Errorf("[pushInfo] push error: %s", err)
		}
	}

	return
}
