package db

import (
	"strconv"

	"github.com/Jordank321/pihouse/data"

	"github.com/jinzhu/gorm"
)

// HumidityRepository is the data repository for Humidity readings
type HumidityRepository interface {
	GetReadingByID(ID int) *data.HumidityReading
	GetLatestForNode(nodeID int) *data.HumidityReading
	GetAllReadings() []*data.HumidityReading
	AddReading(reading *data.HumidityReading)
}

// SQLHumidityRepository is the MSSQL implementation of HumidityRepository
type SQLHumidityRepository struct {
	Connection *gorm.DB
}

func (repository *SQLHumidityRepository) Close() {
	repository.Connection.Close()
}

func (repository *SQLHumidityRepository) GetReadingByID(ID int) *data.HumidityReading {
	var reading data.HumidityReading
	if err := repository.Connection.Preload("Node").First(&reading, ID).Error; err != nil {
		panic(err)
	}
	return &reading
}

func (repository *SQLHumidityRepository) GetLatestForNode(nodeID int) *data.HumidityReading {
	var reading data.HumidityReading
	if err := repository.Connection.Preload("Node").Order("created_at DESC").Where("node_id = " + strconv.Itoa(nodeID)).First(&reading).Error; err != nil {
		panic(err)
	}
	return &reading
}

func (repository *SQLHumidityRepository) GetAllReadings() []*data.HumidityReading {
	var reading []*data.HumidityReading
	if err := repository.Connection.Find(&reading).Error; err != nil {
		panic(err)
	}
	return reading
}

func (repository *SQLHumidityRepository) AddReading(reading *data.HumidityReading) {
	repository.Connection.Create(reading)
}
