package cpu

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	cpu2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
)

func (c *CPU) pushTimes() (err error) {
	vv, err := cpu.Times(true)
	if err != nil {
		return fmt.Errorf("[pushTimes] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := make([]*cpu2.TimesStat, 0, len(vv))

	for _, v := range vv {
		data = append(data, &cpu2.TimesStat{
			Timestamp: t,
			CPU:       v.CPU,
			User:      v.User,
			System:    v.System,
			Idle:      v.Idle,
			Nice:      v.Nice,
			Iowait:    v.Iowait,
			Irq:       v.Irq,
			Softirq:   v.Softirq,
			Steal:     v.Steal,
			Guest:     v.Guest,
			GuestNice: v.GuestNice,
		})
	}

	req := &cpu2.CPURequest{
		Timestamp: t,
		IP:        c.opts.IP,
		NodeName:  c.opts.NodeName,
		TimesStat: data,
	}

	_, err = c.cpuClient.PushCPUTimesStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushInfo] push error: %s", err)
	}

	return
}
