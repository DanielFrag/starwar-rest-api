package repository

import (
	"github.com/DanielFrag/starwar-rest-api/dto"
	"github.com/DanielFrag/starwar-rest-api/service"
	"github.com/DanielFrag/starwar-rest-api/utils"
)

//PlanetAuxRepository interface to access the planet external data
type PlanetAuxRepository interface {
	GetAllPlanets() ([]dto.PlanetDTO, error)
	GetSinglePlanet(externalID int32) (dto.PlanetDTO, error)
	SearchPlanetByName(planetName string) (dto.PlanetDTO, error)
}

//GetPlanetAuxRepository return the entity to access the external planet data
func GetPlanetAuxRepository(r utils.RequestWrapper) PlanetAuxRepository {
	return service.GetPlanetService(r)
}
