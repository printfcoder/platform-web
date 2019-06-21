package os

import "sync"

// module includes cpu, disk, docker, host, load, mem, net
type module struct {
	sync.RWMutex
	api *api
}

type api struct {
}
