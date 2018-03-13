package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DanielFrag/starwar-rest-api/dto"
	"github.com/DanielFrag/starwar-rest-api/utils"
)

//SWAPIService request the external planet data
type SWAPIService struct {
	swapiURL       string
	requestWrapper utils.RequestWrapper
}

//GetSinglePlanet search the planet by its path
func (sw *SWAPIService) GetSinglePlanet(externalID int32) (dto.PlanetDTO, error) {
	var p dto.PlanetDTO
	path := fmt.Sprintf("%v/planets/%v", sw.swapiURL, externalID)
	res, resError := sw.requestWrapper.PerformRequest(path, "GET", nil, nil)
	if resError != nil {
		return p, resError
	}
	jsonError := json.Unmarshal(res, &p)
	if jsonError != nil {
		return p, jsonError
	}
	return p, nil
}

//GetAllPlanets return all planets of swapi
func (sw *SWAPIService) GetAllPlanets() ([]dto.PlanetDTO, error) {
	planets := []dto.PlanetDTO{}
	nextPage := fmt.Sprintf("%v/planets", sw.swapiURL)
	for hasNext := true; hasNext; hasNext = nextPage != "" {
		planetList, planetListError := sw.getPlanets(nextPage)
		if planetListError != nil {
			return planets, planetListError
		}
		nextPage = planetList.Next
		planets = append(planets, planetList.Results...)
	}
	return planets, nil
}

//SearchPlanetByName search the planet by its name
func (sw *SWAPIService) SearchPlanetByName(planetName string) (dto.PlanetDTO, error) {
	var planetList dto.PlanetListDTO
	var planetListError error
	var planet dto.PlanetDTO
	nextPage := fmt.Sprintf("%v/planets", sw.swapiURL)
	for hasNext := true; hasNext; hasNext = (planet.Name == "" && nextPage != "") {
		planetList, planetListError = sw.getPlanets(nextPage)
		if planetListError != nil {
			return planet, planetListError
		}
		nextPage = planetList.Next
		planet = filterPlanetsByName(planetList.Results, planetName)
	}
	if planet.Name == "" {
		planetListError = errors.New("Not found")
	}
	return planet, planetListError
}

func (sw *SWAPIService) getPlanets(url string) (dto.PlanetListDTO, error) {
	var planetList dto.PlanetListDTO
	res, resError := sw.requestWrapper.PerformRequest(url, "GET", nil, nil)
	if resError != nil {
		return planetList, resError
	}
	jsonError := json.Unmarshal(res, &planetList)
	return planetList, jsonError
}

func filterPlanetsByName(planets []dto.PlanetDTO, target string) dto.PlanetDTO {
	for _, planet := range planets {
		if planet.Name == target {
			return planet
		}
	}
	return dto.PlanetDTO{}
}

//GetPlanetService return the entity responsable to access the external planet data
func GetPlanetService(r utils.RequestWrapper) *SWAPIService {
	return &SWAPIService{
		swapiURL:       "https://swapi.co/api",
		requestWrapper: r,
	}
}
