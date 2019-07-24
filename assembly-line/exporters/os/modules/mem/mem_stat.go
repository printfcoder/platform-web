package mem

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/util/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/prometheus/node_exporter/collector/mem"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

func (m *Mem) pushMemoryStat() (err error) {
	mem, err := mem.GetMemInfo()
	if err != nil {
		err = fmt.Errorf("[ERR] [pushMemoryStat] get infos error: %s", err)
		log.Log(err)
		return err
	}

	t := ptypes.TimestampNow()
	data := []*proto.MemoryStat{{
		Timestamp:            t,
		ActiveBytes:          mem["active_bytes"],
		CompressedBytes:      mem["compressed_bytes"],
		InactiveBytes:        mem["inactive_bytes"],
		WiredBytes:           mem["wired_bytes"],
		FreeBytes:            mem["free_bytes"],
		SwappedInBytesTotal:  mem["swapped_in_bytes_total"],
		SwappedOutBytesTotal: mem["swapped_out_bytes_total"],
		TotalBytes:           mem["total_bytes"],
	}}
	req := &proto.MemRequest{
		Timestamp:  t,
		IP:         m.opts.IP,
		NodeName:   m.opts.NodeName,
		MemoryStat: data,
	}

	_, err = m.memClient.PushMemoryStat(context.Background(), req)
	if err != nil {
		err = fmt.Errorf("[ERR] [pushMemoryStat] push error: %s", err)
		log.Log(err)
		return err
	}

	return
}
