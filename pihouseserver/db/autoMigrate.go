package db

import (
	"pihouse/data"

	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&data.Node{}, &data.TemperatureReading{})
}
