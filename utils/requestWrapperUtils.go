package utils

import "io"

//RequestWrapper exposes the method to perform a web request
type RequestWrapper interface {
	PerformRequest(url string, method string, headers map[string]string, body io.Reader) ([]byte, error)
}

//GetRequestWrapper return the entity to perform web requests
func GetRequestWrapper() RequestWrapper {
	return &RequestUtils{}
}
