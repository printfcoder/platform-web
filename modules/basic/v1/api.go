package v1

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/micro-in-cn/platform-web/internal/tools"
	"github.com/micro-in-cn/platform-web/modules/internal/helper"
	"github.com/micro-in-cn/platform-web/modules/internal/nosj"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/registry"
)

type api struct {
}

type rpcRequest struct {
	Service  string
	Endpoint string
	Method   string
	Address  string
	URL      string
	timeout  int
	Request  interface{}
}

// serviceAPIDetail is the service api detail
type serviceAPIDetail struct {
	Name      string               `json:"name,omitempty"`
	Endpoints []*registry.Endpoint `json:"endpoints,omitempty"`
}

func (api *api) webServices(w http.ResponseWriter, r *http.Request) {
	services, err := (*cmd.DefaultOptions().Registry).ListServices()
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	webServices := make([]*registry.Service, 0)
	for _, s := range services {
		for _, webN := range WebNamespacePrefix {
			if strings.Index(s.Name, webN) == 0 && len(strings.TrimPrefix(s.Name, webN)) > 0 {
				s.Name = strings.Replace(s.Name, webN+".", "", 1)
				webServices = append(webServices, s)
			}
		}
	}

	sort.Sort(tools.SortedServices{Services: services})

	nosj.WriteJsonData(w, webServices)

	return
}

func (api *api) services(w http.ResponseWriter, r *http.Request) {

	services, err := (*cmd.DefaultOptions().Registry).ListServices()
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	for _, service := range services {
		ss, err := (*cmd.DefaultOptions().Registry).GetService(service.Name)
		if err != nil {
			continue
		}
		if len(ss) == 0 {
			continue
		}

		for _, s := range ss {
			service.Nodes = append(service.Nodes, s.Nodes...)
		}
	}

	sort.Sort(tools.SortedServices{Services: services})

	nosj.WriteJsonData(w, services)
	return
}

func (api *api) microServices(w http.ResponseWriter, r *http.Request) {

	services, err := (*cmd.DefaultOptions().Registry).ListServices()
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	ret := make([]*registry.Service, 0)

	for _, srv := range services {
		temp, err := (*cmd.DefaultOptions().Registry).GetService(srv.Name)
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}

		for _, s := range temp {
			for _, n := range s.Nodes {
				if n.Metadata["registry"] != "" {
					ret = append(ret, s)
					break
				}
			}
		}
	}

	sort.Sort(tools.SortedServices{Services: ret})

	nosj.WriteJsonData(w, ret)
	return
}

func (api *api) serviceDetails(w http.ResponseWriter, r *http.Request) {
	services, err := (*cmd.DefaultOptions().Registry).ListServices()
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	sort.Sort(tools.SortedServices{Services: services})

	serviceDetails := make([]*serviceAPIDetail, 0)
	for _, service := range services {
		s, err := (*cmd.DefaultOptions().Registry).GetService(service.Name)
		if err != nil {
			continue
		}
		if len(s) == 0 {
			continue
		}

		serviceDetails = append(serviceDetails, &serviceAPIDetail{
			Name:      service.Name,
			Endpoints: s[0].Endpoints,
		})
	}

	nosj.WriteJsonData(w, serviceDetails)
	return
}

func (api *api) service(w http.ResponseWriter, r *http.Request) {

	serviceName := r.URL.Query().Get("service")

	if len(serviceName) > 0 {
		s, err := (*cmd.DefaultOptions().Registry).GetService(serviceName)
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}

		if len(s) == 0 {
			nosj.WriteError(w, fmt.Errorf("Service Is Not found %s: ", serviceName))
			return
		}

		nosj.WriteJsonData(w, s)
		return
	}

	return
}

func (api *api) apiGatewayServices(w http.ResponseWriter, r *http.Request) {

	services, err := (*cmd.DefaultOptions().Registry).ListServices()
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	sel := selector.NewSelector(
		selector.Registry(*cmd.DefaultOptions().Registry),
	)

	ret := make([]*registry.Service, 0)

	for _, service := range services {

		_, _ = sel.Select(service.Name, func(options *selector.SelectOptions) {

			filter := func(services []*registry.Service) []*registry.Service {
				for _, s := range services {
					for _, gwN := range GatewayNamespaces {
						if s.Name == gwN {
							ret = append(ret, s)
							break
						}
					}
				}
				return ret
			}

			options.Filters = append(options.Filters, filter)
		})
	}

	nosj.WriteJsonData(w, ret)
	return

}

func (api *api) rpc(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	rpcReq := &rpcRequest{}

	d := json.NewDecoder(r.Body)
	d.UseNumber()

	if err := d.Decode(&rpcReq); err != nil {
		nosj.WriteError(w, fmt.Errorf("rpc decode err %s: ", err))
		return
	}

	if len(rpcReq.Endpoint) == 0 {
		rpcReq.Endpoint = rpcReq.Method
	}

	rpcReq.timeout, _ = strconv.Atoi(r.Header.Get("Timeout"))
	rpcReq.URL = r.URL.Path

	rpc(w, helper.RequestToContext(r), rpcReq)
}

func (api *api) health(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rpcReq := &rpcRequest{
		Service:  r.URL.Query().Get("service"),
		Endpoint: "Debug.Health",
		Request:  "{}",
		URL:      r.URL.Path,
		Address:  r.URL.Query().Get("address"),
	}

	rpc(w, helper.RequestToContext(r), rpcReq)
}

func (api *api) stats(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	rpcReq := &rpcRequest{
		Service:  r.URL.Query().Get("service"),
		Endpoint: "Debug.Stats",
		Request:  "{}",
		URL:      r.URL.Path,
		Address:  r.URL.Query().Get("address"),
	}

	rpc(w, helper.RequestToContext(r), rpcReq)
	return
}
