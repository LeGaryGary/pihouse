package main

import (
	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/jinzhu/gorm"
)

func ProvideDB() (*gorm.DB, error) {
	return gorm.Open("mssql",
		GetSqlConnectionString())
}

func ProvideTemperaureRepository() db.TemperatureRepository {
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	return &db.SQLTemperatureRepository{Connection: dbret}
}

func ProvideNodeRepository() db.NodeRepository {
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	return &db.SQLNodeRepository{Connection: dbret}
}
