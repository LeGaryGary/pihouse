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

type HumidityController struct {
	humidityRepo db.HumidityRepository
}

type HumidityControllerMethod func(controller *HumidityController, w http.ResponseWriter, r *http.Request)

func HumidityRoutes(getHumidityRepo func() db.HumidityRepository) (string, *chi.Mux) {
	router := chi.NewRouter()

	doMethod := func(method HumidityControllerMethod) func(w http.ResponseWriter, r *http.Request) {
		controller := &HumidityController{
			humidityRepo: getHumidityRepo(),
		}
		return func(w http.ResponseWriter, r *http.Request) {
			method(controller, w, r)
		}
	}

	router.Get("/", doMethod((*HumidityController).GetAllHumidityReadings))
	router.Get("/latest/{nodeID}", doMethod((*HumidityController).GetLatestHumidityForNode))
	router.Post("/", doMethod((*HumidityController).CreateHumidityReading))
	router.Get("/{HumidityReadingId}", doMethod((*HumidityController).GetHumidityReadingByID))
	return "/humidity", router
}

// GetHumidityReadingByID retrieves a single Humidity reading by its ID
func (controller *HumidityController) GetHumidityReadingByID(w http.ResponseWriter, r *http.Request) {
	readingID, err := strconv.Atoi(chi.URLParam(r, "HumidityReadingId"))
	if err != nil {
		panic(err.Error())
	}

	reading := controller.humidityRepo.GetReadingByID(readingID)
	render.JSON(w, r, reading)
}

// GetLatestHumidityForNode retrieves a single Humidity reading by its ID
func (controller *HumidityController) GetLatestHumidityForNode(w http.ResponseWriter, r *http.Request) {
	nodeID, err := strconv.Atoi(chi.URLParam(r, "nodeID"))
	if err != nil {
		panic(err.Error())
	}
	reading := controller.humidityRepo.GetLatestForNode(nodeID)
	render.JSON(w, r, reading)
}

// GetAllHumidityReadings retrieves all Humidity readings
func (controller *HumidityController) GetAllHumidityReadings(w http.ResponseWriter, r *http.Request) {
	readings := controller.humidityRepo.GetAllReadings()
	render.JSON(w, r, readings)
}

// CreateHumidityReading creates a new Humidity reading
func (controller *HumidityController) CreateHumidityReading(w http.ResponseWriter, r *http.Request) {
	read := &data.HumidityReading{}
	if err := json.NewDecoder(r.Body).Decode(read); err != nil {
		panic(err.Error())
	}

	controller.humidityRepo.AddReading(read)

	response := make(map[string]string)
	response["message"] = "Success: the reading has been added"
	render.JSON(w, r, response)
}
