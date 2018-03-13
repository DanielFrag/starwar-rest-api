package router

import (
	"net/http"

	"github.com/DanielFrag/starwar-rest-api/handler"
	"github.com/DanielFrag/starwar-rest-api/utils"
	"github.com/gorilla/mux"
)

//Route store the route data
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc utils.HandlerFuncInjector
}

var routes = []Route{
	Route{
		Method:  "POST",
		Pattern: "/api/planets",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.JSONContentTypeChecker,
				handler.PlanetRepositoryInjector,
			},
			Handler: handler.AddPlanet,
		},
	},
	Route{
		Method:  "GET",
		Pattern: "/api/planets",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.PlanetRepositoryInjector,
			},
			Handler: handler.GetPlanets,
		},
	},
	Route{
		Method:  "DELETE",
		Pattern: "/api/planets/{id}",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.PlanetRepositoryInjector,
			},
			Handler: handler.RemovePlanet,
		},
	},
}

//NewRouter return the application router with its handlers
func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		route.HandlerFunc.InjectDependencies()
		router.
			HandleFunc(route.Pattern, route.HandlerFunc.Handler).
			Methods(route.Method)
	}
	return handler.CorsSetup(router)
}
