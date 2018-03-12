package service

import (
	"fmt"
	"log"
	"testing"
)

func TestGetPlanetByID(t *testing.T) {
	t.Run("GetPlanetByID", func(t *testing.T) {
		result, resultError := GetSinglePlanet(fmt.Sprintf("%v/planets/%v", getAPIUrl(), 2))
		if resultError != nil {
			t.Error("Get planet by id error: " + resultError.Error())
		}
		log.Println(result)
	})
	t.Run("GetPlanetByName", func(t *testing.T) {
		result, resultError := SearchPlanetByName("Chandrila")
		if resultError != nil {
			t.Error("Get planet by id error: " + resultError.Error())
		}
		log.Println(result)
	})
}
