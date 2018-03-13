package repository

import (
	"github.com/DanielFrag/starwar-rest-api/infra"
	"github.com/DanielFrag/starwar-rest-api/model"
)

//PlanetRepository interface to access the planet database
type PlanetRepository interface {
	AddPlanet(planet model.Planet) (string, error)
	GetPlanets() ([]model.Planet, error)
	FindPlanetByID(planetID string) (model.Planet, error)
	FindPlanetByName(planetName string) (model.Planet, error)
	RemovePlanet(planetID string) error
}

//GetPlanetRepository return the entity to access the database
func GetPlanetRepository() PlanetRepository {
	return infra.GetPlanetDB()
}
