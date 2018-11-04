package data

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type TemperatureReading struct {
	gorm.Model
	Value  decimal.Decimal `sql:"type:decimal(4,2)"`
	Node   Node
	NodeID uint `gorm:"index"`
}
