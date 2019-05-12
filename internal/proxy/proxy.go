package proxy

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/selector"
	"github.com/micro/micro/web/common"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
)

var (
	re = regexp.MustCompile("^[a-zA-Z0-9]+([a-zA-Z0-9-]*[a-zA-Z0-9]*)?$")
)

func proxy() http.Handler {
	sel := selector.NewSelector(
		selector.Registry((*cmd.DefaultOptions().Registry)),
	)

	director := func(r *http.Request) {
		kill := func() {
			r.URL.Host = ""
			r.URL.Path = ""
			r.URL.Scheme = ""
			r.Host = ""
			r.RequestURI = ""
		}

		parts := strings.Split(r.URL.Path, "/proxy/")
		if len(parts) < 2 {
			kill()
			return
		}
		if !re.MatchString(parts[1]) {
			kill()
			return
		}
		next, err := sel.Select(Namespace + "." + parts[1])
		if err != nil {
			kill()
			return
		}

		s, err := next()
		if err != nil {
			kill()
			return
		}

		r.Header.Set(common.BasePathHeader, "/"+parts[1])
		r.URL.Host = fmt.Sprintf("%s:%d", s.Address, s.Port)
		r.URL.Path = "/" + strings.Join(parts[2:], "/")
		r.URL.Scheme = "http"
		r.Host = r.URL.Host
	}

	return &common.Proxy{
		Default:  &httputil.ReverseProxy{Director: director},
		Director: director,
	}
}
