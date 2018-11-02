package db

import "PiHouse/data"

type TemperatureRepository interface {
	GetReadingByID(ID int) *data.TemperatureReading
	GetAllReadings() []*data.TemperatureReading
	AddReading(reading *data.TemperatureReading)
}

type SqlTemperatureRepository struct {
}
