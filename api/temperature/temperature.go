package temperature

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	"github.com/shopspring/decimal"
)

type Reading struct {
	TemperatureReadingId int
	Value                decimal.Decimal
	NodeId               int
	Timestamp            time.Time
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllReadings)
	router.Post("/", CreateReading)
	router.Get("/{TemperatureReadingId}", GetReadingByID)
	return router
}

// GetReadingByID retrieves a single temperature reading by its ID
func GetReadingByID(w http.ResponseWriter, r *http.Request) {
	readingID, _ := strconv.Atoi(chi.URLParam(r, "TemperatureReadingId"))
	val, _ := decimal.NewFromString("47.6")
	reading := Reading{
		TemperatureReadingId: readingID,
		Value:                val,
		NodeId:               1,
		Timestamp:            time.Date(2018, time.October, 30, 21, 50, 33, 1, time.Local),
	}
	render.JSON(w, r, reading)
}

// GetAllReadings retrieves all temperature readings
func GetAllReadings(w http.ResponseWriter, r *http.Request) {
	val, _ := decimal.NewFromString("47.6")
	readings := []Reading{
		{
			TemperatureReadingId: 1,
			Value:                val,
			NodeId:               1,
			Timestamp:            time.Date(2018, time.October, 30, 21, 50, 33, 1, time.Local),
		},
	}
	render.JSON(w, r, readings)
}

// CreateReading creates a new temperature reading
func CreateReading(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Success: the reading has been added"
	render.JSON(w, r, response)
}
