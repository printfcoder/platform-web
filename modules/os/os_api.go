package os

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/micro-in-cn/platform-web/internal/config"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/cli"
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

func (o *api) ipGroup(w http.ResponseWriter, r *http.Request) {
	v := config.GetConfig("os", "ip-group")
	ret := map[string]interface{}{}
	err := json.Unmarshal(v.Bytes(), &ret)
	if err != nil {
		nosj.WriteError(w, fmt.Errorf("[ipGroup] err: %s", err))
		return
	}
	nosj.WriteJsonData(w, ret)
	return
}
