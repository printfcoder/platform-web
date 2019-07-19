package net

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/net"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/net"
)

func (p *Net) pushConnectionStat() (err error) {
	t := ptypes.TimestampNow()
	data := make([]*proto.ConnectionStat, 0)
	for _, kind := range p.kinds {
		vv, err := net.Connections(kind)
		if err != nil {
			return fmt.Errorf("[pushConnectionStat] get infos error: %s", err)
		}

		for _, v := range vv {
			lAddr := &proto.Addr{
				IP:   v.Laddr.IP,
				Port: v.Laddr.Port,
			}
			rAddr := &proto.Addr{
				IP:   v.Raddr.IP,
				Port: v.Raddr.Port,
			}

			data = append(data, &proto.ConnectionStat{
				Timestamp: t,
				Fd:        v.Fd,
				Family:    v.Family,
				Type:      v.Type,
				Laddr:     lAddr,
				Raddr:     rAddr,
				Status:    v.Status,
				Uids:      v.Uids,
				Pid:       v.Pid,
			})
		}
	}

	req := &proto.NetRequest{
		Timestamp:      t,
		IP:             p.IP,
		NodeName:       p.NodeName,
		ConnectionStat: data,
	}

	_, err = p.netClient.PushConnectionStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushConnectionStat] push error: %s", err)
	}

	return
}
