package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/DanielFrag/starwar-rest-api/dto"
	"github.com/DanielFrag/starwar-rest-api/utils"
)

var swapiURL string

//GetSinglePlanet search the planet by its id
func GetSinglePlanet(path string) (dto.Planet, error) {
	var p dto.PlanetDTO
	res, resError := utils.PerformRequest(path, "GET", nil, nil)
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
func SearchPlanetByName(planetName string) (dto.Planet, error) {
	var planetList dto.PlanetListDTO
	var planet dto.PlanetDTO
	planetList, planetListError := getPlanets(fmt.Sprintf("%v/planets", getAPIUrl()))
	if planetListError != nil {
		return planet, planetListError
	}
	planet = filterPlanetsByName(planetList.Results, planetName)
	for i := 0; planet.Name == "" && planetList.Next != ""; i++ {
		log.Println(i)
		log.Println(planetList.Next)
		planetList, planetListError = getPlanets(planetList.Next)
		if planetListError != nil {
			return planet, planetListError
		}
		planet = filterPlanetsByName(planetList.Results, planetName)
	}
	return planet, nil
}

func getPlanets(url string) (dto.PlanetListDTO, error) {
	var planetList dto.PlanetListDTO
	res, resError := utils.PerformRequest(url, "GET", nil, nil)
	if resError != nil {
		return planetList, resError
	}
	jsonError := json.Unmarshal(res, &planetList)
	return planetList, jsonError
}

func filterPlanetsByName(planets []dto.Planet, target string) dto.Planet {
	for _, planet := range planets {
		if planet.Name == target {
			return planet
		}
	}
	return dto.Planet{}
}

func getAPIUrl() string {
	if swapiURL == "" {
		swapiURL = os.Getenv("SWAPI_URL")
		if swapiURL == "" {
			swapiURL = "https://swapi.co/api"
		}
	}
	return swapiURL
}
