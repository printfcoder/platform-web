package os

import (
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro/cli"
	"sync"
)

var (
	once sync.Once
)

type api struct {
}

func newAPI() *api {
	return &api{}
}

func (o *api) init(ctx *cli.Context) {
	db.Init(ctx)
}
