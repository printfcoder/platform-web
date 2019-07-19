package load

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/load"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/load"
)

func (p *Load) pushAvgStat() (err error) {
	v, err := load.Avg()
	if err != nil {
		return fmt.Errorf("[pushAvgStat] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := []*proto.LoadAvgStat{
		{
			Timestamp: ptypes.TimestampNow(),
			Load1:     v.Load1,
			Load5:     v.Load5,
			Load15:    v.Load15,
		},
	}

	req := &proto.LoadRequest{
		Timestamp:   t,
		IP:          p.IP,
		NodeName:    p.NodeName,
		LoadAvgStat: data,
	}

	_, err = p.loadClient.PushLoadAvgStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushAvgStat] push error: %s", err)
	}

	return
}
