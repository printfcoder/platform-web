package modules

// Module is the biggest unit of web component.
// a component is a webapp which will be registered on root path by Name function
type Module interface {

	// Name of module
	Name() string
	// Path returns the root path of module
	Path() string
	// Handlers returns all handlers of module
	Handlers()
}

// Rsp is the struct of http api response
type Rsp struct {
	Code    uint        `json:"code,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
