package mock

import (
	"errors"

	"github.com/DanielFrag/starwar-rest-api/model"
	"gopkg.in/mgo.v2/bson"
)

//PlanetMock implements the PlanetRepository interface for tests
type PlanetMock struct {
	planetList []model.Planet
}

//AddPlanet push planet into 'planetList'
func (p *PlanetMock) AddPlanet(planet model.Planet) (string, error) {
	planet.ID = bson.NewObjectId()
	p.planetList = append(p.planetList, planet)
	return planet.ID.Hex(), nil
}

//GetPlanets return the 'planetList'
func (p *PlanetMock) GetPlanets() ([]model.Planet, error) {
	if p.planetList == nil {
		p.planetList = []model.Planet{}
	}
	return p.planetList, nil
}

//FindPlanetByID find planet by its id
func (p *PlanetMock) FindPlanetByID(planetID string) (model.Planet, error) {
	for _, planet := range p.planetList {
		if planet.ID.Hex() == planetID {
			return planet, nil
		}
	}
	return model.Planet{}, errors.New("not found")
}

//FindPlanetByName find planet by its name
func (p *PlanetMock) FindPlanetByName(planetName string) (model.Planet, error) {
	for _, planet := range p.planetList {
		if planet.Name == planetName {
			return planet, nil
		}
	}
	return model.Planet{}, errors.New("not found")
}

//RemovePlanet find and remove a planet from 'planetList', based on its id
func (p *PlanetMock) RemovePlanet(planetID string) error {
	index := -1
	for i := 0; i < len(p.planetList) && index == -1; i++ {
		if p.planetList[i].ID.Hex() == planetID {
			index = i
		}
	}
	if index == -1 {
		return errors.New("not found")
	}
	p.planetList = append(p.planetList[:index], p.planetList[index+1:]...)
	return nil
}

//InitializeDBMock populate the 'planetList'
func (p *PlanetMock) InitializeDBMock() {
	p.planetList = []model.Planet{
		model.Planet{
			ID:        bson.NewObjectId(),
			Name:      "Tatooine",
			Climate:   "hot",
			Terrain:   "stone",
			ForeignID: 1,
		},
		model.Planet{
			ID:        bson.NewObjectId(),
			Name:      "foo",
			Climate:   "cold",
			Terrain:   "wather",
			ForeignID: 0,
		},
		model.Planet{
			ID:        bson.NewObjectId(),
			Name:      "bar",
			Climate:   "ok",
			Terrain:   "grass",
			ForeignID: 0,
		},
		model.Planet{
			ID:        bson.NewObjectId(),
			Name:      "sunda",
			Climate:   "hot",
			Terrain:   "stone",
			ForeignID: 0,
		},
	}
}
