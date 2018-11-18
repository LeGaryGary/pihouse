package db

import (
	"github.com/Jordank321/pihouse/data"

	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&data.Node{},
		&data.TemperatureReading{},
		&data.HumidityReading{},
		&data.AIRequest{},
		&data.Intent{})
}
