package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DanielFrag/starwar-rest-api/dto"
	"github.com/DanielFrag/starwar-rest-api/utils"
)

//PlanetService wrap the access to planet aux data
type PlanetService interface {
	GetSinglePlanet(externalID int32) (dto.PlanetDTO, error)
	SearchPlanetByName(planetName string) (dto.PlanetDTO, error)
}

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

//SearchPlanetByName search the planet by its name
func (sw *SWAPIService) SearchPlanetByName(planetName string) (dto.PlanetDTO, error) {
	var planetList dto.PlanetListDTO
	var planet dto.PlanetDTO
	planetList, planetListError := sw.getPlanets(fmt.Sprintf("%v/planets", sw.swapiURL))
	if planetListError != nil {
		return planet, planetListError
	}
	planet = filterPlanetsByName(planetList.Results, planetName)
	for i := 0; planet.Name == "" && planetList.Next != ""; i++ {
		planetList, planetListError = sw.getPlanets(planetList.Next)
		if planetListError != nil {
			return planet, planetListError
		}
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

func (sw *SWAPIService) setAPIUrl(url string) {
	sw.swapiURL = url
}

func filterPlanetsByName(planets []dto.PlanetDTO, target string) dto.PlanetDTO {
	for _, planet := range planets {
		if planet.Name == target {
			return planet
		}
	}
	return dto.PlanetDTO{}
}

/*
os.Getenv("SWAPI_URL")
if swapiURL == "" {
	swapiURL = "https://swapi.co/api"
}
*/
