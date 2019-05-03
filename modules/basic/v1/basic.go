package v1

import (
	"github.com/gorilla/mux"
	"sync"
)

func init() {
	m = &basicModule{
		name: "basic",
		path: "/",
	}
}

var (
	m *basicModule
)

// basicModule includes web, registry, CLI, Stats submodules.
type basicModule struct {
	name string
	path string
	sync.RWMutex
	api *api
}

func (m *basicModule) Handlers() {

}

func (m *basicModule) initHandlers(r *mux.Router) {
	r.HandleFunc("/api/v1/services", m.api.services).Methods("GET")
	r.HandleFunc("/api/v1/micro-services", m.api.microServices).Methods("GET")

	r.HandleFunc("/api/v1/service/{name:[a-zA-Z0-9/.]+}", m.api.service).Methods("GET")
	r.HandleFunc("/api/v1/api-gateway-services", m.api.apiGatewayServices).Methods("GET")

	r.HandleFunc("/api/v1/service-details", m.api.serviceDetails).Methods("GET")
	r.HandleFunc("/api/v1/stats", m.api.stats).Methods("GET")
	r.Path("/api/v1/api-stats").Handler(apiProxy()).Methods("GET")
	r.HandleFunc("/api/v1/web-services", m.api.webServices).Methods("GET")
	r.HandleFunc("/api/v1/rpc", m.api.rpc).Methods("POST")
	r.HandleFunc("/api/v1/health", m.api.health).Methods("GET")
}
