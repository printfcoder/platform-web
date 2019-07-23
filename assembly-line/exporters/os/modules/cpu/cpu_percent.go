package cpu

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	cpu2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
)

func (c *CPU) pushPercent() (err error) {
	vv, err := cpu.Percent(c.opts.Interval, true)
	if err != nil {
		return fmt.Errorf("[pushPercent] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := make([]*cpu2.Percent, 0, len(vv))

	for _, v := range vv {
		data = append(data, &cpu2.Percent{
			Timestamp: t,
			Percent:   v,
		})
	}

	req := &cpu2.CPURequest{
		Timestamp: t,
		IP:        c.IP,
		NodeName:  c.NodeName,
		Percent:   data,
	}

	_, err = c.cpuClient.PushCPUPercent(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushPercent] push error: %s", err)
	}

	return
}
