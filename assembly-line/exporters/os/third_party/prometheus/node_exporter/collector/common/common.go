package common

import (
	"github.com/prometheus/procfs"
	"gopkg.in/alecthomas/kingpin.v2"
	"path/filepath"
)

var (
	procPath = kingpin.Flag("path.procfs", "procfs mountpoint.").Default(procfs.DefaultMountPoint).String()
)

func ProcFilePath(name string) string {
	return filepath.Join(*procPath, name)
}
