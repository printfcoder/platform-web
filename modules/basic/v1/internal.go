package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro-in-cn/platform-web/internal/proxy"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/errors"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func rpc(w http.ResponseWriter, ctx context.Context, rpcReq *rpcRequest) {

	if len(rpcReq.Service) == 0 {
		nosj.WriteError(w, fmt.Errorf("Service Is Not found "))
		return
	}

	if len(rpcReq.Endpoint) == 0 {
		nosj.WriteError(w, fmt.Errorf("Endpoint Is Not found err "))
		return
	}

	// decode rpc request param body
	if req, ok := rpcReq.Request.(string); ok {
		d := json.NewDecoder(strings.NewReader(req))
		d.UseNumber()

		if err := d.Decode(&rpcReq.Request); err != nil {
			nosj.WriteError(w, fmt.Errorf("error decoding request string err: %s", err))
			return
		}
	}

	// create request/response
	var response json.RawMessage
	var err error
	req := (*cmd.DefaultOptions().Client).NewRequest(rpcReq.Service, rpcReq.Endpoint, rpcReq.Request, client.WithContentType("application/json"))

	var opts []client.CallOption

	// set timeout
	if rpcReq.timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(rpcReq.timeout)*time.Second))
	}

	// remote call
	if len(rpcReq.Address) > 0 {
		opts = append(opts, client.WithAddress(rpcReq.Address))
	}

	// remote call
	err = (*cmd.DefaultOptions().Client).Call(ctx, req, &response, opts...)
	if err != nil {
		ce := errors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = "go.micro.rpc"
			ce.Status = http.StatusText(500)
			ce.Detail = "error during request: " + ce.Detail
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Code))
		}
		w.Write([]byte(ce.Error()))
		return
	}

	if strings.Contains(rpcReq.URL, "/v1/rpc") {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		nosj.WriteJsonData(w, response)
	}
}

func apiProxy() http.Handler {
	director := func(r *http.Request) {
		kill := func() {
			r.URL.Host = ""
			r.URL.Path = ""
			r.URL.Scheme = ""
			r.Host = ""
			r.RequestURI = ""
		}

		parts := strings.Split(r.URL.Path, "/v1/api-stats")
		if len(parts) < 2 {
			kill()
			return
		}

		address := r.URL.Query().Get("address")
		r.Header.Set(proxy.BasePathHeader, "/v1/api-stats")
		r.Header.Set("Content-Type", "application/json")
		r.URL.Host = fmt.Sprintf("%s", address)
		r.URL.Path = "/stats" + strings.Join(parts[2:], "/")
		r.URL.Scheme = "http"
		r.Host = r.URL.Host
	}

	return &proxy.Proxy{
		Default:  &httputil.ReverseProxy{Director: director},
		Director: director,
	}
}
