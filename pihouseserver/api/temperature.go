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

type TemperatureController struct {
	temperatureRepo db.TemperatureRepository
}

type TemperatureControllerMethod func(controller *TemperatureController, w http.ResponseWriter, r *http.Request)

func TemperatureRoutes(getTempRepo func() db.TemperatureRepository) (string, *chi.Mux) {
	router := chi.NewRouter()
	doMethod := func(method TemperatureControllerMethod) func(w http.ResponseWriter, r *http.Request) {
		controller := &TemperatureController{
			temperatureRepo: getTempRepo(),
		}
		return func(w http.ResponseWriter, r *http.Request) {
			method(controller, w, r)
		}
	}
	router.Get("/", doMethod((*TemperatureController).GetAllReadings))
	router.Get("/latest/{nodeID}", doMethod((*TemperatureController).GetLatestForNode))
	router.Post("/", doMethod((*TemperatureController).CreateReading))
	router.Get("/{TemperatureReadingId}", doMethod((*TemperatureController).GetReadingByID))
	return "/temperature", router
}

// GetReadingByID retrieves a single temperature reading by its ID
func (controller *TemperatureController) GetReadingByID(w http.ResponseWriter, r *http.Request) {
	readingID, err := strconv.Atoi(chi.URLParam(r, "TemperatureReadingId"))
	if err != nil {
		panic(err.Error())
	}
	reading := controller.temperatureRepo.GetReadingByID(readingID)
	render.JSON(w, r, reading)
}

// GetReadingByID retrieves a single temperature reading by its ID
func (controller *TemperatureController) GetLatestForNode(w http.ResponseWriter, r *http.Request) {
	nodeID, err := strconv.Atoi(chi.URLParam(r, "nodeID"))
	if err != nil {
		panic(err.Error())
	}
	reading := controller.temperatureRepo.GetLatestForNode(nodeID)
	render.JSON(w, r, reading)
}

// GetAllReadings retrieves all temperature readings
func (controller *TemperatureController) GetAllReadings(w http.ResponseWriter, r *http.Request) {
	readings := controller.temperatureRepo.GetAllReadings()
	render.JSON(w, r, readings)
}

// CreateReading creates a new temperature reading
func (controller *TemperatureController) CreateReading(w http.ResponseWriter, r *http.Request) {
	read := &data.TemperatureReading{}
	if err := json.NewDecoder(r.Body).Decode(read); err != nil {
		panic(err.Error())
	}

	controller.temperatureRepo.AddReading(read)

	response := make(map[string]string)
	response["message"] = "Success: the reading has been added"
	render.JSON(w, r, response)
}
