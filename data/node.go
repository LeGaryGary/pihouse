package data

import "github.com/jinzhu/gorm"

type Node struct {
	gorm.Model
	TemperatureReadings []TemperatureReading
	Name                string
}
