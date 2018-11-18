package db

import (
	"strconv"

	"github.com/Jordank321/pihouse/data"

	"github.com/jinzhu/gorm"
)

// TemperatureRepository is the data repository for temperature readings
type TemperatureRepository interface {
	GetReadingByID(ID int) *data.TemperatureReading
	GetLatestForNode(nodeID int) *data.TemperatureReading
	GetAllReadings() []*data.TemperatureReading
	AddReading(reading *data.TemperatureReading)
}

// SQLTemperatureRepository is the MSSQL implementation of TemperatureRepository
type SQLTemperatureRepository struct {
	Connection *gorm.DB
}

func (repository *SQLTemperatureRepository) Close() {
	repository.Connection.Close()
}

func (repository *SQLTemperatureRepository) GetReadingByID(ID int) *data.TemperatureReading {
	var reading data.TemperatureReading
	if err := repository.Connection.Preload("Node").First(&reading, ID).Error; err != nil {
		panic(err)
	}
	return &reading
}

func (repository *SQLTemperatureRepository) GetLatestForNode(nodeID int) *data.TemperatureReading {
	var reading data.TemperatureReading
	if err := repository.Connection.Preload("Node").Order("CreatedAt DESC").Where("NodeID = " + strconv.Itoa(nodeID)).First(&reading).Error; err != nil {
		panic(err)
	}
	return &reading
}

func (repository *SQLTemperatureRepository) GetAllReadings() []*data.TemperatureReading {
	var reading []*data.TemperatureReading
	if err := repository.Connection.Find(&reading).Error; err != nil {
		panic(err)
	}
	return reading
}

func (repository *SQLTemperatureRepository) AddReading(reading *data.TemperatureReading) {
	repository.Connection.Create(reading)
}
