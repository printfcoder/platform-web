package mem

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/mem"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

func (p *Pusher) pushSwapMemoryStat() (err error) {
	v, err := mem.SwapMemory()
	if err != nil {
		return fmt.Errorf("[pushSwapMemoryStat] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := []*proto.SwapMemoryStat{{
		Timestamp:   t,
		Total:       v.Total,
		Used:        v.Used,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
		Sin:         v.Sin,
		Sout:        v.Sout,
		PgIn:        v.PgIn,
		PgOut:       v.PgOut,
		PgFault:     v.PgFault,
	}}

	req := &proto.MemRequest{
		Timestamp:      t,
		IP:             p.IP,
		NodeName:       p.NodeName,
		SwapMemoryStat: data,
	}

	_, err = p.memClient.PushSwapMemoryStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushSwapMemoryStat] push error: %s", err)
	}

	return
}
