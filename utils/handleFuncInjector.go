package utils

import (
	"net/http"
)

//HandlerFuncInjector store the http.HandlerFunc and inject its middlewares
type HandlerFuncInjector struct {
	Dependencies []func(http.HandlerFunc) http.HandlerFunc
	Handler      http.HandlerFunc
}

//InjectDependencies inject the middlewares
func (hfi *HandlerFuncInjector) InjectDependencies() {
	for i := len(hfi.Dependencies) - 1; i >= 0; i-- {
		hfi.Handler = hfi.Dependencies[i](hfi.Handler)
	}
}
