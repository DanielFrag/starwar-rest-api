package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/DanielFrag/starwar-rest-api/mock"
	"github.com/DanielFrag/starwar-rest-api/model"
	"github.com/DanielFrag/starwar-rest-api/repository"
	"github.com/DanielFrag/starwar-rest-api/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func TestAddPlanet(t *testing.T) {
	hfi, planetRepository, planetAuxRepository := prepareHandler(AddPlanet, false)
	t.Run("AddValidPlanet", func(t *testing.T) {
		planetName := "Chandrila"
		planet := model.Planet{
			Name:    planetName,
			Terrain: "stone",
			Climate: "hot",
		}
		body := bytes.NewReader(utils.FormatJSON(planet))
		req, reqError := http.NewRequest("POST", "/", body)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 201 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 201, returned %v", result.StatusCode))
			return
		}
		planet, planetError := planetRepository.FindPlanetByName(planetName)
		if planetError != nil {
			t.Error("DB error: " + planetError.Error())
			return
		}
		if planet.Name != planetName && planet.ForeignID == 0 {
			t.Error("Planet not added")
			return
		}
	})
	t.Run("AddInvalidPlanet", func(t *testing.T) {
		planetName := "unknowPlanet"
		_, planetExternalDataError := planetAuxRepository.SearchPlanetByName(planetName)
		if planetExternalDataError == nil {
			t.Error("The planet should not exists in aux repository")
			return
		}
		planet := model.Planet{
			Name:    planetName,
			Terrain: "stone",
			Climate: "hot",
		}
		body := bytes.NewReader(utils.FormatJSON(planet))
		req, reqError := http.NewRequest("POST", "/", body)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 201 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 201, returned %v", result.StatusCode))
			return
		}
		planet, planetError := planetRepository.FindPlanetByName(planetName)
		if planetError != nil {
			t.Error("DB error: " + planetError.Error())
			return
		}
		if planet.Name != planetName {
			t.Error("Planet not added")
			return
		}
		if planet.ForeignID != 0 {
			t.Error(fmt.Sprintf("Inconsistent data. Planet foreign id should be 0, but got %v", planet.ForeignID))
			return
		}
	})
}

func TestGetPlanets(t *testing.T) {
	t.Run("EmptyDB", func(t *testing.T) {
		hfi, _, _ := prepareHandler(GetPlanets, false)
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		var planets []model.Planet
		json.Unmarshal(reqRecorder.Body.Bytes(), &planets)
		if len(planets) != 0 {
			t.Error("The DB should be empty")
			return
		}
	})
	t.Run("PopulatedDB", func(t *testing.T) {
		hfi, planetRepository, planetAuxRepository := prepareHandler(GetPlanets, true)
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		var planets []model.Planet
		json.Unmarshal(reqRecorder.Body.Bytes(), &planets)
		if len(planets) == 0 {
			t.Error("The DB should be populated")
			return
		}
		for _, p := range planets {
			_, internalPlanetDataError := planetRepository.FindPlanetByID(p.ID.Hex())
			externalPlanetData, externalPlanetDataError := planetAuxRepository.GetSinglePlanet(p.ForeignID)
			if internalPlanetDataError != nil {
				t.Error("Inconsistent data: " + internalPlanetDataError.Error())
				return
			}
			if externalPlanetDataError != nil && p.ForeignID != 0 && p.NumberOfFilms != int32(len(externalPlanetData.Films)) {
				t.Error(fmt.Sprintf("Inconsistent data the number of films should be %v, but returned %v", p.NumberOfFilms, len(externalPlanetData.Films)))
				return
			}
		}
	})
}

func TestFindPlanetByName(t *testing.T) {
	hfi, planetRepository, planetAuxRepository := prepareHandler(FindPlanetByName, true)
	t.Run("WithExternalData", func(t *testing.T) {
		planetName := "Tatooine"
		req, reqError := http.NewRequest("GET", "/"+planetName, nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{name}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		var planet model.Planet
		json.Unmarshal(reqRecorder.Body.Bytes(), &planet)
		if planet.Name != planetName {
			t.Error("Invalid planet result")
			return
		}
		_, internalPlanetDataError := planetRepository.FindPlanetByID(planet.ID.Hex())
		externalPlanetData, externalPlanetDataError := planetAuxRepository.GetSinglePlanet(planet.ForeignID)
		if internalPlanetDataError != nil {
			t.Error("Inconsistent data: " + internalPlanetDataError.Error())
			return
		}
		if externalPlanetDataError != nil && planet.ForeignID != 0 && planet.NumberOfFilms != int32(len(externalPlanetData.Films)) {
			t.Error(fmt.Sprintf("Inconsistent data the number of films should be %v, but returned %v", planet.NumberOfFilms, len(externalPlanetData.Films)))
			return
		}
	})
	t.Run("WithoutExternalData", func(t *testing.T) {
		planets, _ := planetRepository.GetPlanets()
		var planetName string
		for i := 0; i < len(planets) && planetName == ""; i++ {
			if planets[i].ForeignID == 0 {
				planetName = planets[i].Name
			}
		}
		req, reqError := http.NewRequest("GET", "/"+planetName, nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{name}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		var planet model.Planet
		json.Unmarshal(reqRecorder.Body.Bytes(), &planet)
		if planet.Name != planetName {
			t.Error("Invalid planet result")
			return
		}
		if planet.NumberOfFilms != 0 {
			t.Error(fmt.Sprintf("Inconsistent data the number of films should be 0, but returned %v", planet.NumberOfFilms))
			return
		}
	})
	t.Run("Invalid", func(t *testing.T) {
		planetName := "unknowPlanet"
		req, reqError := http.NewRequest("GET", "/"+planetName, nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{name}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 404 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 404 and got: %v", result.StatusCode))
			return
		}
	})
	t.Run("NameNotProvided", func(t *testing.T) {
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{name}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 404 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 404 and got: %v", result.StatusCode))
			return
		}
	})
}

func TestFindPlanetByID(t *testing.T) {
	hfi, planetRepository, planetAuxRepository := prepareHandler(FindPlanetByID, true)
	planets, _ := planetRepository.GetPlanets()
	t.Run("WithExternalData", func(t *testing.T) {
		var validID bson.ObjectId
		for i := 0; i < len(planets) && &validID != nil; i++ {
			if planets[i].ForeignID != 0 {
				validID = planets[i].ID
			}
		}
		req, reqError := http.NewRequest("GET", "/"+validID.Hex(), nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		var planet model.Planet
		json.Unmarshal(reqRecorder.Body.Bytes(), &planet)
		if planet.ID.Hex() != validID.Hex() {
			t.Error("Invalid planet result")
			return
		}
		_, internalPlanetDataError := planetRepository.FindPlanetByID(planet.ID.Hex())
		externalPlanetData, externalPlanetDataError := planetAuxRepository.GetSinglePlanet(planet.ForeignID)
		if internalPlanetDataError != nil {
			t.Error("Inconsistent data: " + internalPlanetDataError.Error())
			return
		}
		if externalPlanetDataError != nil && planet.ForeignID != 0 && planet.NumberOfFilms != int32(len(externalPlanetData.Films)) {
			t.Error(fmt.Sprintf("Inconsistent data the number of films should be %v, but returned %v", planet.NumberOfFilms, len(externalPlanetData.Films)))
			return
		}
	})
	t.Run("WithoutExternalData", func(t *testing.T) {
		var validID bson.ObjectId
		for i := 0; i < len(planets) && &validID != nil; i++ {
			if planets[i].ForeignID == 0 {
				validID = planets[i].ID
			}
		}
		req, reqError := http.NewRequest("GET", "/"+validID.Hex(), nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		var planet model.Planet
		json.Unmarshal(reqRecorder.Body.Bytes(), &planet)
		if planet.ID.Hex() != validID.Hex() {
			t.Error("Invalid planet result")
			return
		}
		if planet.NumberOfFilms != 0 {
			t.Error(fmt.Sprintf("Inconsistent data the number of films should be 0, but returned %v", planet.NumberOfFilms))
			return
		}
	})
	t.Run("Invalid", func(t *testing.T) {
		invalidID := bson.NewObjectId()
		req, reqError := http.NewRequest("GET", "/"+invalidID.Hex(), nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 404 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 404 and got: %v", result.StatusCode))
			return
		}
	})
	t.Run("IDNotProvided", func(t *testing.T) {
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{name}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 404 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 404 and got: %v", result.StatusCode))
			return
		}
	})
}

func TestRemovePlanet(t *testing.T) {
	hfi, planetRepository, _ := prepareHandler(RemovePlanet, true)
	t.Run("ValidID", func(t *testing.T) {
		planets, _ := planetRepository.GetPlanets()
		req, reqError := http.NewRequest("GET", "/"+planets[0].ID.Hex(), nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		updatedPlanetsCollection, _ := planetRepository.GetPlanets()
		if len(updatedPlanetsCollection) >= len(planets) {
			t.Error(fmt.Sprintf("The expected length after remove was %v, but got %v", len(updatedPlanetsCollection), len(planets)))
			return
		}
	})
	t.Run("InvalidID", func(t *testing.T) {
		invalidID := bson.NewObjectId()
		req, reqError := http.NewRequest("GET", "/"+invalidID.Hex(), nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 404 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 404 and got: %v", result.StatusCode))
		}
	})
	t.Run("IDNotProvided", func(t *testing.T) {
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
			return
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{name}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 404 {
			t.Error(fmt.Sprintf("Wrong status code. Expected 404 and got: %v", result.StatusCode))
			return
		}
	})
}

func prepareHandler(h http.HandlerFunc, initializeDB bool) (utils.HandlerFuncInjector, repository.PlanetRepository, repository.PlanetAuxRepository) {
	requestMock := mock.RequestMock{}
	requestMock.AddRequest("https://swapi.co/api/planets/-1", "GET", nil)
	requestMock.AddRequest("https://swapi.co/api/planets/1", "GET", []byte(mock.Planet1))
	requestMock.AddRequest("https://swapi.co/api/planets", "GET", []byte(mock.PlanetPage1))
	requestMock.AddRequest("https://swapi.co/api/planets/?page=2", "GET", []byte(mock.PlanetPage2))
	planetAuxRepository := repository.GetPlanetAuxRepository(&requestMock)
	planetRepository := mock.PlanetMock{}
	if initializeDB {
		planetRepository.InitializeDBMock()
	}
	injectRepositoryDependencies := func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "PlanetRepository", &planetRepository)
			context.Set(r, "PlanetAuxRepository", planetAuxRepository)
			next(w, r)
			return
		})
	}
	hfi := utils.HandlerFuncInjector{
		Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
			injectRepositoryDependencies,
		},
		Handler: h,
	}
	hfi.InjectDependencies()
	return hfi, &planetRepository, planetAuxRepository
}
