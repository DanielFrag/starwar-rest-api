package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/DanielFrag/starwar-rest-api/model"
	"github.com/DanielFrag/starwar-rest-api/repository"
	"github.com/DanielFrag/starwar-rest-api/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//AddPlanet save a new planet
func AddPlanet(w http.ResponseWriter, r *http.Request) {
	planetRepository, planetAuxRepository, planetRepositoryError := extractPlanetRepository(r)
	if planetRepositoryError != nil {
		http.Error(w, planetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body request: "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	var p model.Planet
	jsonError := json.Unmarshal(body, &p)
	if jsonError != nil {
		http.Error(w, "Json error: "+jsonError.Error(), http.StatusInternalServerError)
		return
	}
	planetExternalData, planetExternalDataError := planetAuxRepository.SearchPlanetByName(p.Name)
	if planetExternalDataError == nil {
		path := strings.Split(planetExternalData.URL, "/")
		for i := len(path) - 1; p.ForeignID == 0 && i >= 0; i-- {
			if id, idError := strconv.ParseInt(path[i], 10, 32); idError == nil {
				p.ForeignID = int32(id)
			}
		}
	}
	planetID, addError := planetRepository.AddPlanet(p)
	if addError != nil {
		http.Error(w, "Repository error: "+addError.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Header().Set("Location", "/planets/id/"+planetID)
	w.Write([]byte("added!"))
}

//GetPlanets return all the planets
func GetPlanets(w http.ResponseWriter, r *http.Request) {
	planetRepository, planetAuxRepository, planetRepositoryError := extractPlanetRepository(r)
	if planetRepositoryError != nil {
		http.Error(w, planetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	planets, planetsError := planetRepository.GetPlanets()
	if planetsError != nil {
		http.Error(w, "Can't find the requested planets: "+planetsError.Error(), http.StatusNotFound)
		return
	}
	worker, semaphore := make(chan int, 3), make(chan int, len(planets))
	for i := range planets {
		planetIndex := i
		worker <- 1
		go func(index int) {
			if planets[index].ForeignID > 0 {
				planets[index].NumberOfFilms = getPlanetNumberOfFilms(planets[index].ForeignID, planetAuxRepository)
			}
			<-worker
			semaphore <- 1
		}(planetIndex)
	}
	for i := 0; i < len(planets); i++ {
		<-semaphore
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(planets))
	return
}

//FindPlanetByID find planet by id
func FindPlanetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "Error id not provided", http.StatusBadRequest)
		return
	}
	planetRepository, planetAuxRepository, planetRepositoryError := extractPlanetRepository(r)
	if planetRepositoryError != nil {
		http.Error(w, planetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	planet, planetError := planetRepository.FindPlanetByID(vars["id"])
	if planetError != nil {
		http.Error(w, "Can't find the requested planet: "+planetError.Error(), http.StatusNotFound)
		return
	}
	if planet.ForeignID > 0 {
		planet.NumberOfFilms = getPlanetNumberOfFilms(planet.ForeignID, planetAuxRepository)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(planet))
	return
}

//FindPlanetByName find planet by name
func FindPlanetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["name"] == "" {
		http.Error(w, "Error name not provided", http.StatusBadRequest)
		return
	}
	planetRepository, planetAuxRepository, planetRepositoryError := extractPlanetRepository(r)
	if planetRepositoryError != nil {
		http.Error(w, planetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	planet, planetError := planetRepository.FindPlanetByName(vars["name"])
	if planetError != nil {
		http.Error(w, "Can't find the requested planet: "+planetError.Error(), http.StatusNotFound)
		return
	}
	if planet.ForeignID > 0 {
		planet.NumberOfFilms = getPlanetNumberOfFilms(planet.ForeignID, planetAuxRepository)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(planet))
	return
}

//RemovePlanet find and remove a planet based on its id
func RemovePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "Error id not provided", http.StatusBadRequest)
		return
	}
	planetRepository, _, planetRepositoryError := extractPlanetRepository(r)
	if planetRepositoryError != nil {
		http.Error(w, planetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	planetError := planetRepository.RemovePlanet(vars["id"])
	if planetError != nil {
		message := planetError.Error()
		var code int
		if strings.Contains(message, "not found") {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}
		http.Error(w, message, code)
		return
	}
	w.WriteHeader(204)
}

func extractPlanetRepository(r *http.Request) (repository.PlanetRepository, repository.PlanetAuxRepository, error) {
	contextPlanetRepository := context.Get(r, "PlanetRepository")
	contextPlanetAuxRepository := context.Get(r, "PlanetAuxRepository")
	if contextPlanetRepository == nil {
		return nil, nil, errors.New("Can't access the context 'PlanetRepository'")
	}
	if contextPlanetAuxRepository == nil {
		return nil, nil, errors.New("Can't access the context 'PlanetAuxRepository'")
	}
	planetRepository, planetRepositoryOk := contextPlanetRepository.(repository.PlanetRepository)
	planetAuxRepository, planetAuxRepositoryOk := contextPlanetAuxRepository.(repository.PlanetAuxRepository)
	if !planetRepositoryOk || !planetAuxRepositoryOk {
		return nil, nil, errors.New("Can't access the planet repository")
	}
	return planetRepository, planetAuxRepository, nil
}

func getPlanetNumberOfFilms(planetForeignID int32, r repository.PlanetAuxRepository) int32 {
	planetData, planetDataError := r.GetSinglePlanet(planetForeignID)
	if planetDataError != nil {
		return 0
	}
	return int32(len(planetData.Films))
}
