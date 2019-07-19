package mem

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/prometheus/node_exporter/collector/mem"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

func (p *Mem) pushMemoryStat() (err error) {
	m, err := mem.GetMemInfo()
	if err != nil {
		return fmt.Errorf("[pushMemoryStat] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := []*proto.MemoryStat{{
		Timestamp:            t,
		ActiveBytes:          m["active_bytes"],
		CompressedBytes:      m["compressed_bytes"],
		InactiveBytes:        m["inactive_bytes"],
		WiredBytes:           m["wired_bytes"],
		FreeBytes:            m["free_bytes"],
		SwappedInBytesTotal:  m["swapped_in_bytes_total"],
		SwappedOutBytesTotal: m["swapped_out_bytes_total"],
		TotalBytes:           m["total_bytes"],
	}}
	req := &proto.MemRequest{
		Timestamp:  t,
		IP:         p.IP,
		NodeName:   p.NodeName,
		MemoryStat: data,
	}

	_, err = p.memClient.PushMemoryStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushMemoryStat] push error: %s", err)
	}

	return
}
