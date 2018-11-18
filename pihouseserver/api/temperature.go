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
	GetTempRepo func() db.TemperatureRepository
)

func TemperatureRoutes(getTempRepo func() db.TemperatureRepository) *chi.Mux {
	GetTempRepo = getTempRepo
	router := chi.NewRouter()
	router.Get("/", GetAllReadings)
	router.Get("/latest/{nodeID}", GetLatestForNode)
	router.Post("/", CreateReading)
	router.Get("/{TemperatureReadingId}", GetReadingByID)
	return router
}

// GetReadingByID retrieves a single temperature reading by its ID
func GetReadingByID(w http.ResponseWriter, r *http.Request) {
	readingID, err := strconv.Atoi(chi.URLParam(r, "TemperatureReadingId"))
	if err != nil {
		panic(err.Error())
	}
	repo := GetTempRepo()
	reading := repo.GetReadingByID(readingID)
	render.JSON(w, r, reading)
}

// GetReadingByID retrieves a single temperature reading by its ID
func GetLatestForNode(w http.ResponseWriter, r *http.Request) {
	nodeID, err := strconv.Atoi(chi.URLParam(r, "nodeID"))
	if err != nil {
		panic(err.Error())
	}
	repo := GetTempRepo()
	reading := repo.GetLatestForNode(nodeID)
	render.JSON(w, r, reading)
}

// GetAllReadings retrieves all temperature readings
func GetAllReadings(w http.ResponseWriter, r *http.Request) {
	repo := GetTempRepo()
	readings := repo.GetAllReadings()
	render.JSON(w, r, readings)
}

// CreateReading creates a new temperature reading
func CreateReading(w http.ResponseWriter, r *http.Request) {
	read := &data.TemperatureReading{}
	if err := json.NewDecoder(r.Body).Decode(read); err != nil {
		panic(err.Error())
	}

	repo := GetTempRepo()
	repo.AddReading(read)

	response := make(map[string]string)
	response["message"] = "Success: the reading has been added"
	render.JSON(w, r, response)
}
