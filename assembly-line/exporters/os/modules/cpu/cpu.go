package cpu

import (
	"fmt"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/option"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/mem"
)

var ()

func Init(ops []option.Option) {

	// cores
	// v, _ := cpu.Counts(true)

	// percent

	// status

	// times

	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
