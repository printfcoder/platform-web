package nosj

// I am JSON

import (
	"encoding/json"
	"github.com/micro-in-cn/platform-web/modules"
	"net/http"
)

func WriteJsonData(w http.ResponseWriter, data interface{}) {

	rsp := &modules.Rsp{
		Data:    data,
		Success: true,
	}

	b, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func WriteError(w http.ResponseWriter, msg string) {

	rsp := &modules.Rsp{
		Error:   msg,
		Success: false,
	}

	b, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
