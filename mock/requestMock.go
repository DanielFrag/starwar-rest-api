package mock

import (
	"errors"
	"io"
)

type requestData struct {
	pattern string
	method  string
	data    []byte
}

//RequestMock simulte a http request
type RequestMock struct {
	requests []requestData
}

//AddRequest add a new pattern for responses
func (r *RequestMock) AddRequest(url, method string, responseData []byte) {
	r.requests = append(r.requests, requestData{
		pattern: url,
		method:  method,
		data:    responseData,
	})
}

//PerformRequest simulate swapi request
func (r *RequestMock) PerformRequest(url string, method string, headers map[string]string, body io.Reader) ([]byte, error) {
	for _, req := range r.requests {
		if req.pattern == url && req.method == method {
			return req.data, nil
		}
	}
	return nil, errors.New("not found")
}
