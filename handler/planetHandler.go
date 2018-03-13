package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/DanielFrag/starwar-rest-api/dto"
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
	if p.Name == "" || p.Climate == "" || p.Terrain == "" {
		http.Error(w, "Incomplete planet data (missing name, climate, terrain)", http.StatusBadRequest)
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
	w.Header().Set("Location", "/api/planets?id="+planetID)
	w.WriteHeader(201)
	w.Write([]byte("added!"))
}

//GetPlanets query planets data
func GetPlanets(w http.ResponseWriter, r *http.Request) {
	planetRepository, planetAuxRepository, planetRepositoryError := extractPlanetRepository(r)
	if planetRepositoryError != nil {
		http.Error(w, planetRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	var resp interface{}
	var respError error
	if len(r.URL.Query()) > 0 {
		resp, respError = getSinglePlanetData(r.URL.Query(), planetRepository, planetAuxRepository)
	} else {
		resp, respError = getAllPlanetsData(planetRepository, planetAuxRepository)
	}
	if respError != nil {
		http.Error(w, respError.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(resp))
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

func getAllPlanetsData(pRepo repository.PlanetRepository, pAuxRepo repository.PlanetAuxRepository) ([]model.Planet, error) {
	planets, planetsError := pRepo.GetPlanets()
	if planetsError != nil {
		return planets, errors.New("Can't find the requested planets: " + planetsError.Error())
	}
	if len(planets) < 9 {
		populateNumOfFilmsWithID(planets, pAuxRepo)
	} else {
		populateNumOfFilmsWithMap(planets, pAuxRepo)
	}
	return planets, nil
}

func getSinglePlanetData(p url.Values, pRepo repository.PlanetRepository, pAuxRepo repository.PlanetAuxRepository) (model.Planet, error) {
	var planet model.Planet
	var planetError error
	if p.Get("id") != "" {
		planet, planetError = pRepo.FindPlanetByID(p.Get("id"))
	} else if p.Get("name") != "" {
		planet, planetError = pRepo.FindPlanetByName(p.Get("name"))
	} else {
		planetError = errors.New("Can't find any planet with the provided identifier")
	}
	if planetError != nil {
		return planet, errors.New("Can't find the requested planet: " + planetError.Error())
	}
	planet.NumberOfFilms = getSinglePlanetNumberOfFilms(planet, pAuxRepo)
	return planet, nil
}

func getSinglePlanetNumberOfFilms(planet model.Planet, r repository.PlanetAuxRepository) int32 {
	var planetData dto.PlanetDTO
	var planetDataError error
	if planet.ForeignID > 0 {
		planetData, planetDataError = r.GetSinglePlanet(planet.ForeignID)
	} else {
		planetData, planetDataError = r.SearchPlanetByName(planet.Name)
	}
	if planetDataError != nil {
		return 0
	}
	return int32(len(planetData.Films))
}

func populateNumOfFilmsWithID(planets []model.Planet, pAuxRepo repository.PlanetAuxRepository) {
	worker, semaphore := make(chan int, 3), make(chan int, len(planets))
	for i := range planets {
		planetIndex := i
		worker <- 1
		go func(index int) {
			planets[index].NumberOfFilms = getSinglePlanetNumberOfFilms(planets[index], pAuxRepo)
			<-worker
			semaphore <- 1
		}(planetIndex)
	}
	for i := 0; i < len(planets); i++ {
		<-semaphore
	}
}

func populateNumOfFilmsWithMap(planets []model.Planet, pAuxRepo repository.PlanetAuxRepository) {
	planetsDTO, planetsDTOError := pAuxRepo.GetAllPlanets()
	if planetsDTOError != nil {
		return
	}
	m := make(map[string]*model.Planet)
	for _, p := range planets {
		m[p.Name] = &p
	}
	for _, p := range planetsDTO {
		if m[p.Name] != nil {
			m[p.Name].NumberOfFilms = int32(len(p.Films))
		}
	}
}
