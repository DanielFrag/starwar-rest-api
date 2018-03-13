package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//RequestUtils wrap the web request for external resources
type RequestUtils struct {
	client *http.Client
}

//PerformRequest perform a web request
func (r *RequestUtils) PerformRequest(url, method string, headers map[string]string, body io.Reader) ([]byte, error) {
	request, requestError := http.NewRequest(method, url, body)
	if requestError != nil {
		return nil, requestError
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}
	if r.client == nil {
		r.initializeClient()
	}
	response, requisitionError := r.client.Do(request)
	if requisitionError != nil {
		return nil, requisitionError
	}
	defer response.Body.Close()
	if response.StatusCode/100 > 2 {
		return nil, errors.New(response.Status)
	}
	bodyByte, readBodyError := ioutil.ReadAll(response.Body)
	if readBodyError != nil {
		return nil, readBodyError
	}
	return bodyByte, nil
}

func (r *RequestUtils) initializeClient() {
	r.client = &http.Client{
		Timeout: time.Second * time.Duration(getTimeoutNumber()),
	}
}

func getTimeoutNumber() int64 {
	timeoutStr := os.Getenv("TIMEOUT_SECONDS")
	if timeoutStr == "" {
		timeoutStr = "30"
	}
	timeoutNum, timeoutNumError := strconv.ParseInt(timeoutStr, 10, 0)
	if timeoutNumError != nil {
		timeoutNum = 30
	}
	return timeoutNum
}
