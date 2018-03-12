package repository

import "github.com/DanielFrag/starwar-rest-api/model"

type PlanetRepository interface {
	AddPlanet(planet model.Planet) (string, error)
	GetPlanets() ([]model.Planet, error)
	FindPlanetByID(planetID string) (model.Planet, error)
	FindPlanetByName(planetName string) (model.Planet, error)
	RemovePlanet(planetID string) error
}
