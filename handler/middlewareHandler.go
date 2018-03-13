package handler

import (
	"net/http"
	"strings"

	"github.com/DanielFrag/starwar-rest-api/repository"
	"github.com/DanielFrag/starwar-rest-api/utils"
	"github.com/gorilla/context"
)

//JSONContentTypeChecker check the content-type header of request allowing only application/json
func JSONContentTypeChecker(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("content-type")
		if contentType == "" || !strings.Contains(contentType, "application/json") {
			http.Error(w, "Only json is supported", http.StatusUnsupportedMediaType)
			return
		}
		next(w, r)
		return
	})
}

//PlanetRepositoryInjector inject the interface to access the planet data
func PlanetRepositoryInjector(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestWrapper := utils.GetRequestWrapper()
		context.Set(r, "PlanetRepository", repository.GetPlanetRepository())
		context.Set(r, "PlanetAuxRepository", repository.GetPlanetAuxRepository(requestWrapper))
		next(w, r)
		return
	})
}
