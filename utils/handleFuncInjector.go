package utils

import (
	"net/http"
)

type HandlerFuncInjector struct {
	Dependencies []func(http.HandlerFunc) http.HandlerFunc
	Handler     http.HandlerFunc
}

func (hfi *HandlerFuncInjector) InjectDependencies() {
	for i := len(hfi.Dependencies) - 1; i >=0; i-- {
		hfi.Handler = hfi.Dependencies[i](hfi.Handler)
	}
}
