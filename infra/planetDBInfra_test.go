package infra

import (
	"fmt"
	"testing"

	"github.com/DanielFrag/starwar-rest-api/model"
	"gopkg.in/mgo.v2/bson"
)

//TestPlanetDB test the MongoDB manipulation
func TestPlanetDB(t *testing.T) {
	testPlanetDB(t, PlanetMGO{})
}

func testPlanetDB(t *testing.T, planetRepository PlanetMGO) {
	var planetID1, planetID2 string
	defer func() {
		recoverError := recover()
		if recoverError != nil {
			t.Error(recoverError)
		}
		StopDB()
	}()
	t.Run("StartDB", func(t *testing.T) {
		startDBError := StartDB()
		if startDBError != nil {
			t.Error("Can't starts the DB")
			return
		}
		ds.dbName = ds.dbName + "_test"
		mgoSession := getSession()
		dropDatabaseError := mgoSession.DB(getDbName()).DropDatabase()
		if dropDatabaseError != nil {
			t.Error(dropDatabaseError)
		}
	})
	t.Run("CheckDB", func(t *testing.T) {
		planets, planetsError := planetRepository.GetPlanets()
		if planetsError != nil {
			t.Error("Checking DB error: " + planetsError.Error())
			return
		}
		if len(planets) != 0 {
			t.Error("The DB should be empty")
			return
		}
	})
	t.Run("AddPlanet1", func(t *testing.T) {
		planetID, planetsError := planetRepository.AddPlanet(model.Planet{
			ID:        bson.NewObjectId(),
			Name:      "sunda",
			Climate:   "hot",
			Terrain:   "stone",
			ForeignID: 1,
		})
		if planetsError != nil {
			t.Error("Checking DB error: " + planetsError.Error())
			return
		}
		if planetID == "" {
			t.Error("It should return the planet id")
			return
		}
		planetID1 = planetID
	})
	t.Run("AddPlanet2", func(t *testing.T) {
		planetID, planetsError := planetRepository.AddPlanet(model.Planet{
			ID:        bson.NewObjectId(),
			Name:      "sunda",
			Climate:   "hot",
			Terrain:   "stone",
			ForeignID: 1,
		})
		if planetsError != nil {
			t.Error("Checking DB error: " + planetsError.Error())
			return
		}
		if planetID == "" {
			t.Error("It should return the planet id")
			return
		}
		planetID2 = planetID
		if planetID1 == planetID2 {
			t.Error("The plantet ids should be unique")
			return
		}
	})
	t.Run("GetPlanets", func(t *testing.T) {
		planets, planetsError := planetRepository.GetPlanets()
		if planetsError != nil {
			t.Error("Checking DB error: " + planetsError.Error())
			return
		}
		if len(planets) != 2 {
			t.Error("The DB should contains the planet1 and the planet2")
			return
		}
		if planets[0].ID.Hex() != planetID1 {
			t.Error(fmt.Sprintf("Inconsistent data. The planet1 id should be %v, but got %v", planetID1, planets[0].ID.Hex()))
			return
		}
		if planets[1].ID.Hex() != planetID2 {
			t.Error(fmt.Sprintf("Inconsistent data. The planet2 id should be %v, but got %v", planetID2, planets[1].ID.Hex()))
			return
		}
	})
	t.Run("FindPlanetByID", func(t *testing.T) {
		planet, planetError := planetRepository.FindPlanetByID(planetID2)
		if planetError != nil {
			t.Error("Could not find the planet2: " + planetError.Error())
			return
		}
		if planet.ID.Hex() != planetID2 {
			t.Error(fmt.Sprintf("Inconsistent data. The returned planet should have the id equals to %v, but got %v", planetID2, planet.ID.Hex()))
			return
		}
	})
	t.Run("FindPlanetByName", func(t *testing.T) {
		planets, planetsError := planetRepository.GetPlanets()
		if planetsError != nil {
			t.Error("Get planets error: " + planetsError.Error())
			return
		}
		if len(planets) != 2 {
			t.Error("The DB should contains the planet1 and the planet2")
			return
		}
		planet, planetError := planetRepository.FindPlanetByName(planets[0].Name)
		if planetsError != nil {
			t.Error("Get planet by name error: " + planetError.Error())
			return
		}
		if planets[0].Name != planet.Name {
			t.Error(fmt.Sprintf("Inconsistent data. The planet name should be %v, but got %v", planets[0].Name, planet.Name))
			return
		}
	})
	t.Run("SearchInexistentPlanetByID", func(t *testing.T) {
		_, planetError := planetRepository.FindPlanetByID(bson.NewObjectId().Hex())
		if planetError == nil {
			t.Error("Should return a not found error")
			return
		}
	})
	t.Run("SearchInexistentPlanetByName", func(t *testing.T) {
		_, planetError := planetRepository.FindPlanetByName("unknowPlanet")
		if planetError == nil {
			t.Error("Should return a not found error")
			return
		}
	})
	t.Run("RemovePlanet", func(t *testing.T) {
		planetRemoveError := planetRepository.RemovePlanet(planetID1)
		if planetRemoveError != nil {
			t.Error("Remove planet error: " + planetRemoveError.Error())
			return
		}
		planets, planetsError := planetRepository.GetPlanets()
		if planetsError != nil {
			t.Error("GetPlanets error: " + planetsError.Error())
			return
		}
		if len(planets) != 1 {
			t.Error("The DB should contains only one planet")
			return
		}
		if planets[0].ID.Hex() == planetID1 {
			t.Error("Inconsistent data. The planet1 should not exists in db after the remove")
			return
		}
	})
	t.Run("RemoveInexistentPlanet", func(t *testing.T) {
		planetRemoveError := planetRepository.RemovePlanet(planetID1)
		if planetRemoveError == nil {
			t.Error("Should return a not found error")
			return
		}
	})
}
