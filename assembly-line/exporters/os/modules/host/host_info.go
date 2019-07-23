package host

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/host"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
)

func (h *Host) pushInfo() (err error) {
	v, err := host.Info()
	if err != nil {
		return fmt.Errorf("[pushInfo] get info error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := []*proto.HostInfo{
		{
			Timestamp:            ptypes.TimestampNow(),
			Hostname:             v.Hostname,
			Uptime:               v.Uptime,
			BootTime:             v.BootTime,
			Procs:                v.Procs,
			OS:                   v.OS,
			Platform:             v.Platform,
			PlatformFamily:       v.PlatformFamily,
			PlatformVersion:      v.PlatformVersion,
			KernelVersion:        v.KernelVersion,
			VirtualizationSystem: v.VirtualizationSystem,
			VirtualizationRole:   v.VirtualizationRole,
		},
	}

	req := &proto.HostRequest{
		Timestamp: t,
		IP:        h.IP,
		NodeName:  h.NodeName,
		HostInfo:  data,
	}

	_, err = h.hostClient.PushHostInfo(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushInfo] push error: %s", err)
	}

	return
}
