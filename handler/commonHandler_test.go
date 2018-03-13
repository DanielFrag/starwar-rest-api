package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCorsSetup(t *testing.T) {
	fA := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	fB := CorsSetup(fA)
	req, reqError := http.NewRequest("GET", "/", nil)
	if reqError != nil {
		t.Error("Error to create the request: " + reqError.Error())
	}
	req.Header["Origin"] = []string{"anyware"}
	reqRecorder := httptest.NewRecorder()
	fB.ServeHTTP(reqRecorder, req)
	fBHeader := reqRecorder.Header()
	if fBHeader == nil || len(fBHeader) < 0 {
		t.Error("Headers not seted")
	}
	if fBHeader["Access-Control-Allow-Origin"] == nil || len(fBHeader["Access-Control-Allow-Origin"]) != 1 || fBHeader["Access-Control-Allow-Origin"][0] != "anyware" {
		t.Error("Header 'Access-Control-Allow-Origin' not seted")
	}
	allowMethods := strings.Split(fBHeader["Access-Control-Allow-Methods"][0], ", ")
	methodsOk, missingMethodsDesc := checkHeaders(allowMethods, []string{
		"GET",
		"POST",
		"PUT",
		"OPTIONS",
	})
	if !methodsOk {
		t.Error(fmt.Sprintf("Method %v not seted in 'Access-Control-Allow-Methods'", missingMethodsDesc))
	}
	allowHeaders := strings.Split(fBHeader["Access-Control-Allow-Headers"][0], ", ")
	headersOk, missingHeadersDesc := checkHeaders(allowHeaders, []string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"Content-Length",
		"Content-Type",
	})
	if !headersOk {
		t.Error(fmt.Sprintf("Header %v not seted in 'Access-Control-Allow-Headers'", missingHeadersDesc))
	}
}

func checkHeaders(allowedItens []string, requiredItens []string) (bool, string) {
	if len(allowedItens) != len(requiredItens) {
		return false, ""
	}
	m := make(map[string]bool)
	for i := range requiredItens {
		m[requiredItens[i]] = false
	}
	for i := range allowedItens {
		m[allowedItens[i]] = true
	}
	for key, value := range m {
		if !value {
			return false, key
		}
	}
	return true, ""
}
