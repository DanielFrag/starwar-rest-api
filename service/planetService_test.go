package service

import (
	"encoding/json"
	"testing"

	"github.com/DanielFrag/starwar-rest-api/dto"
	"github.com/DanielFrag/starwar-rest-api/mock"
	"github.com/DanielFrag/starwar-rest-api/utils"
)

func TestServiceWithMock(t *testing.T) {
	requestUtils := mock.RequestMock{}
	requestUtils.AddRequest("https://swapi.co/api/planets/-1", "GET", nil)
	requestUtils.AddRequest("https://swapi.co/api/planets/1", "GET", []byte(mock.Planet1))
	requestUtils.AddRequest("https://swapi.co/api/planets", "GET", []byte(mock.PlanetPage1))
	requestUtils.AddRequest("https://swapi.co/api/planets/?page=2", "GET", []byte(mock.PlanetPage2))
	requestUtils.AddRequest("https://swapi.co/api/planets/?page=3", "GET", []byte(mock.PlanetPage3))
	startTest(t, &requestUtils)
}

func TestServiceWithSWAPI(t *testing.T) {
	requestUtils := utils.RequestUtils{}
	startTest(t, &requestUtils)
}

func startTest(t *testing.T, requestWrapper utils.RequestWrapper) {
	swapiService := SWAPIService{
		swapiURL:       "https://swapi.co/api",
		requestWrapper: requestWrapper,
	}
	t.Run("RequestInvalidPlanet", func(t *testing.T) {
		result, resultError := swapiService.GetSinglePlanet(-1)
		if resultError == nil {
			t.Error("Should return a not found error")
			return
		}
		if result.Name != "" {
			t.Error("This planet should not exist")
			return
		}
	})
	t.Run("RequestValidPlanet", func(t *testing.T) {
		result, resultError := swapiService.GetSinglePlanet(1)
		if resultError != nil {
			t.Error("Get planet by id error: " + resultError.Error())
		}
		if result.Name == "" {
			t.Error("This planet should exist")
			return
		}
	})
	t.Run("RequestPlanetByName", func(t *testing.T) {
		result, resultError := swapiService.SearchPlanetByName("Chandrila")
		if resultError != nil {
			t.Error("Get planet by name error: " + resultError.Error())
		}
		if result.Name == "" {
			t.Error("This planet should exist")
			return
		}
	})
	t.Run("RequestInvalidPlanetByName", func(t *testing.T) {
		result, resultError := swapiService.SearchPlanetByName("unknowPlanet")
		if resultError == nil {
			t.Error("Should return a not found error")
			return
		}
		if result.Name != "" {
			t.Error("This planet should not exist")
			return
		}
	})
	t.Run("RequestAllPlanets", func(t *testing.T) {
		result, resultError := swapiService.GetAllPlanets()
		if resultError != nil {
			t.Error("Should not return an error. Detail: " + resultError.Error())
			return
		}
		if len(result) <= 0 {
			t.Error("Should not return a valid list of plantes")
			return
		}
		res, resError := requestWrapper.PerformRequest("https://swapi.co/api/planets", "GET", nil, nil)
		if resError != nil {
			t.Error("Request error: " + resError.Error())
			return
		}
		var planetList dto.PlanetListDTO
		jsonError := json.Unmarshal(res, &planetList)
		if jsonError != nil {
			t.Error("Unmarshal json error: " + jsonError.Error())
			return
		}
		if planetList.Count != int32(len(result)) {
			t.Error("Invalid planets array length")
			return
		}
	})
}
