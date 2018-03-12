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

//PerformRequest wrap request for external resources
func PerformRequest(url, method string, headers map[string]string, body io.Reader) ([]byte, error) {
	request, requestError := http.NewRequest(method, url, body)
	if requestError != nil {
		return nil, requestError
	}
	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}
	customClient := &http.Client{
		Timeout: time.Second * time.Duration(getTimeoutNumber()),
	}
	response, requisitionError := customClient.Do(request)
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
