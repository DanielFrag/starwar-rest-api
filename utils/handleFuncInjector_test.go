package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
)

func TestHandlerFuncInjector(t *testing.T) {
	fA := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "fA", "fA")
			next(w, r)
		}
	}
	fB := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "fB", []rune{'f', 'B'})
			next(w, r)
		}
	}
	fCollection := []func(http.HandlerFunc) http.HandlerFunc{
		fA,
		fB,
	}
	f := HandlerFuncInjector{
		Dependencies: fCollection,
		Handler:      func(w http.ResponseWriter, r *http.Request) {},
	}
	f.InjectDependencies()
	req, reqError := http.NewRequest("GET", "/", nil)
	if reqError != nil {
		t.Error("Error to create the request: " + reqError.Error())
	}
	reqRecorder := httptest.NewRecorder()
	f.Handler.ServeHTTP(reqRecorder, req)
	fAContext := context.Get(req, "fA")
	fBContext := context.Get(req, "fB")
	handlerContextData := context.GetAll(req)
	if fAContext == nil || fBContext == nil || len(handlerContextData) != 2 {
		t.Error("Context not seted")
	}
	if fAContent, ok := fAContext.(string); !ok || fAContent != "fA" {
		t.Error("fA context not seted correctly")
	}
	if fBContent, ok := fBContext.([]rune); !ok || len(fBContent) != 2 || fBContent[0] != 'f' || fBContent[1] != 'B' {
		t.Error("fA context not seted correctly")
	}
}
