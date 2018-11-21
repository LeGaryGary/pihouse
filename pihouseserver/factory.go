package main

import (
	"log"

	"github.com/Jordank321/pihouse/pihouseserver/control"
	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/jinzhu/gorm"
)

var dbConnection *gorm.DB

func provideDB() *gorm.DB {
	if dbConnection != nil {
		return dbConnection
	}
	dbConnection, err := gorm.Open("mssql", GetSqlConnectionString())
	if err != nil {
		log.Panicf(err.Error())
	}
	return dbConnection
}

func ProvideTemperaureRepository() db.TemperatureRepository {
	dbret := provideDB()
	return &db.SQLTemperatureRepository{Connection: dbret}
}

func ProvideNodeRepository() db.NodeRepository {
	dbret := provideDB()
	return &db.SQLNodeRepository{Connection: dbret}
}

func ProvideHumidityRepository() db.HumidityRepository {
	dbret := provideDB()
	return &db.SQLHumidityRepository{Connection: dbret}
}

func ProvideAIRepository() db.AIRepository {
	dbret := provideDB()
	return &db.SQLAIRespository{Connection: dbret}
}

var clientController control.ClientController

func ProvideClientController() control.ClientController {
	if clientController == nil {
		clientController = &control.WebSocketClientController{}
	}
	return clientController
}
