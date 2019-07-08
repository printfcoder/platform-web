package mem

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/mem"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
)

func (p *Pusher) pushVirtualMemoryStat() (err error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("[pushVirtualMemoryStat] get infos error: %s", err)
	}

	t := ptypes.TimestampNow()
	data := []*proto.VirtualMemoryStat{{
		Timestamp:      t,
		Total:          v.Total,
		Available:      v.Available,
		Used:           v.Used,
		UsedPercent:    v.UsedPercent,
		Free:           v.Free,
		Active:         v.Active,
		Inactive:       v.Inactive,
		Wired:          v.Wired,
		Laundry:        v.Laundry,
		Buffers:        v.Buffers,
		Cached:         v.Cached,
		Writeback:      v.Writeback,
		Dirty:          v.Dirty,
		WritebackTmp:   v.WritebackTmp,
		Shared:         v.Shared,
		Slab:           v.Slab,
		SReclaimable:   v.SReclaimable,
		PageTables:     v.PageTables,
		SwapCached:     v.SwapCached,
		CommitLimit:    v.CommitLimit,
		CommittedAS:    v.CommittedAS,
		HighTotal:      v.HighTotal,
		HighFree:       v.HighFree,
		LowTotal:       v.LowTotal,
		LowFree:        v.LowFree,
		SwapTotal:      v.SwapTotal,
		SwapFree:       v.SwapFree,
		Mapped:         v.Mapped,
		VMallocTotal:   v.VMallocTotal,
		VMallocUsed:    v.VMallocUsed,
		VMallocChunk:   v.VMallocChunk,
		HugePagesTotal: v.HugePagesTotal,
		HugePagesFree:  v.HugePagesFree,
		HugePageSize:   v.HugePageSize,
	}}

	req := &proto.MemRequest{
		Timestamp:         t,
		IP:                p.IP,
		NodeName:          p.NodeName,
		VirtualMemoryStat: data,
	}

	_, err = p.memClient.PushVirtualMemoryStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushVirtualMemoryStat] push error: %s", err)
	}

	return
}
