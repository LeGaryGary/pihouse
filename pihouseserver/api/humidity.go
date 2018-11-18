package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jordank321/pihouse/data"

	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

var (
	GetHumidityRepo func() db.HumidityRepository
)

func HumidityRoutes(getHumidityRepo func() db.HumidityRepository) *chi.Mux {
	GetHumidityRepo = getHumidityRepo
	router := chi.NewRouter()
	router.Get("/", GetAllHumidityReadings)
	router.Get("/latest/{nodeID}", GetLatestHumidityForNode)
	router.Post("/", CreateHumidityReading)
	router.Get("/{HumidityReadingId}", GetHumidityReadingByID)
	return router
}

// GetHumidityReadingByID retrieves a single Humidity reading by its ID
func GetHumidityReadingByID(w http.ResponseWriter, r *http.Request) {
	readingID, err := strconv.Atoi(chi.URLParam(r, "HumidityReadingId"))
	if err != nil {
		panic(err.Error())
	}
	repo := GetHumidityRepo()
	reading := repo.GetReadingByID(readingID)
	render.JSON(w, r, reading)
}

// GetLatestHumidityForNode retrieves a single Humidity reading by its ID
func GetLatestHumidityForNode(w http.ResponseWriter, r *http.Request) {
	nodeID, err := strconv.Atoi(chi.URLParam(r, "nodeID"))
	if err != nil {
		panic(err.Error())
	}
	repo := GetHumidityRepo()
	reading := repo.GetLatestForNode(nodeID)
	render.JSON(w, r, reading)
}

// GetAllHumidityReadings retrieves all Humidity readings
func GetAllHumidityReadings(w http.ResponseWriter, r *http.Request) {
	repo := GetHumidityRepo()
	readings := repo.GetAllReadings()
	render.JSON(w, r, readings)
}

// CreateHumidityReading creates a new Humidity reading
func CreateHumidityReading(w http.ResponseWriter, r *http.Request) {
	read := &data.HumidityReading{}
	if err := json.NewDecoder(r.Body).Decode(read); err != nil {
		panic(err.Error())
	}

	repo := GetHumidityRepo()
	repo.AddReading(read)

	response := make(map[string]string)
	response["message"] = "Success: the reading has been added"
	render.JSON(w, r, response)
}
