package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DanielFrag/starwar-rest-api/repository"
	"github.com/DanielFrag/starwar-rest-api/utils"
	"github.com/gorilla/context"
)

func TestRepositoryInjection(t *testing.T) {
	fA := func(w http.ResponseWriter, r *http.Request) {}
	hfi := utils.HandlerFuncInjector{
		Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
			PlanetRepositoryInjector,
		},
		Handler: fA,
	}
	hfi.InjectDependencies()
	req, reqError := http.NewRequest("GET", "/", nil)
	if reqError != nil {
		t.Error("Error to create the request: " + reqError.Error())
		return
	}
	reqRecorder := httptest.NewRecorder()
	hfi.Handler.ServeHTTP(reqRecorder, req)
	planetRepositoryContext := context.Get(req, "PlanetRepository")
	planetAuxRepositoryContext := context.Get(req, "PlanetAuxRepository")
	if planetRepositoryContext == nil || planetAuxRepositoryContext == nil {
		t.Error("Repository not seted in requisiton context")
		return
	}
	_, planetRepositoryOk := planetRepositoryContext.(repository.PlanetRepository)
	_, planetAuxRepositoryOk := planetAuxRepositoryContext.(repository.PlanetAuxRepository)
	if !planetRepositoryOk || !planetAuxRepositoryOk {
		t.Error("Context seted with an invalid reposirories interface")
		return
	}
}

func TestJSONContentTypeChecker(t *testing.T) {
	fA := func(w http.ResponseWriter, r *http.Request) {}
	hfi := utils.HandlerFuncInjector{
		Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
			JSONContentTypeChecker,
		},
		Handler: fA,
	}
	hfi.InjectDependencies()
	t.Run("ValidRequest", func(t *testing.T) {
		req, reqError := http.NewRequest("POST", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		req.Header.Add("Content-Type", "application/json")
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 200 {
			t.Error(fmt.Sprintf("Wrong status code. Should be 200 and got %v", result.StatusCode))
		}
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		req, reqError := http.NewRequest("POST", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		req.Header.Add("Content-Type", "application/xml")
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 415 {
			t.Error(fmt.Sprintf("Wrong status code. Should be 415 and got %v", result.StatusCode))
		}
	})
}
