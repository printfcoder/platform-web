package net

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/net"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
)

func (n *Net) pushIOCountersStat() (err error) {
	t := ptypes.TimestampNow()
	data := make([]*proto.IOCountersStat, 0)
	vv, err := net.IOCounters(true)
	if err != nil {
		return fmt.Errorf("[pushIOCountersStat] get infos error: %s", err)
	}

	for _, v := range vv {
		if !n.ifaceMap[v.Name] {
			continue
		}

		data = append(data, &proto.IOCountersStat{
			Timestamp:   t,
			Name:        v.Name,
			BytesSent:   v.BytesSent,
			BytesRecv:   v.BytesRecv,
			PacketsSent: v.PacketsSent,
			PacketsRecv: v.PacketsRecv,
			Errin:       v.Errin,
			Errout:      v.Errout,
			Dropin:      v.Dropin,
		})
	}

	req := &proto.NetRequest{
		Timestamp:      t,
		IP:             n.IP,
		NodeName:       n.NodeName,
		IOCountersStat: data,
	}

	_, err = n.netClient.PushIOCountersStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushIOCountersStat] push error: %s", err)
	}

	return
}
