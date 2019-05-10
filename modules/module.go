package modules

import (
	"github.com/micro/cli"
	"net/http"
)

// Module is the biggest unit of web component.
// a component is a webapp which will be registered on root path by Name function
type Module interface {
	// Name of module
	Name() string

	// Path returns the root path of module
	Path() string

	// Init initialize the module
	Init(*cli.Context) error

	// Handlers returns http handler of this module
	Handlers() map[string]*Handler

	// Global Flags
	Flags() []cli.Flag

	// Commands are the Sub-commands
	Commands(...Option) []cli.Command
}

type Handler struct {
	Func   func(w http.ResponseWriter, r *http.Request)
	Method []string
	Hld    http.Handler
}

func (h Handler) IsFunc() bool {
	return h.Func != nil
}

// Rsp is the struct of http api response
type Rsp struct {
	Code    uint        `json:"code,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type Option func(*Options)

var (
	modules []Module
)

func Registry(m Module) {
	modules = append(modules, m)
}

// Modules returns all of the registered modules
func Modules() []Module {
	return modules
}
