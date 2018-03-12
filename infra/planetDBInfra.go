package infra

import (
	"github.com/DanielFrag/starwar-rest-api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PlanetMGO wrap the session to access planet data
type PlanetMGO struct {
	session *mgo.Session
}

//AddPlanet save a new planet into db
func (p *PlanetMGO) AddPlanet(planet model.Planet) (string, error) {
	p.session = getSession()
	defer p.session.Close()
	planetCollection := p.session.DB(getDbName()).C("Planets")
	planet.ID = bson.NewObjectId()
	err := planetCollection.Insert(planet)
	return planet.ID.Hex(), err
}

//GetPlanets return all the planets in db
func (p *PlanetMGO) GetPlanets() ([]model.Planet, error) {
	p.session = getSession()
	defer p.session.Close()
	planetCollection := p.session.DB(getDbName()).C("Planets")
	var planets []model.Planet
	err := planetCollection.
		Find(bson.M{}).
		All(&planets)
	return planets, err
}

//FindPlanetByID find planet by its id
func (p *PlanetMGO) FindPlanetByID(planetID string) (model.Planet, error) {
	p.session = getSession()
	defer p.session.Close()
	planetCollection := p.session.DB(getDbName()).C("Planets")
	var planet model.Planet
	err := planetCollection.
		Find(bson.M{
			"_id": bson.ObjectIdHex(planetID),
		}).
		One(&planet)
	return planet, err
}

//FindPlanetByName find planet by its name
func (p *PlanetMGO) FindPlanetByName(planetName string) (model.Planet, error) {
	p.session = getSession()
	defer p.session.Close()
	planetCollection := p.session.DB(getDbName()).C("Planets")
	var planet model.Planet
	err := planetCollection.
		Find(bson.M{
			"name": planetName,
		}).
		One(&planet)
	return planet, err
}

//RemovePlanet find and remove a planet from db, based on its id
func (p *PlanetMGO) RemovePlanet(planetID string) error {
	p.session = getSession()
	defer p.session.Close()
	planetCollection := p.session.DB(getDbName()).C("Planets")
	return planetCollection.RemoveId(bson.ObjectIdHex(planetID))
}
