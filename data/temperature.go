package data

import (
	"time"

	"github.com/shopspring/decimal"
)

type TemperatureReading struct {
	TemperatureReadingID int
	Value                decimal.Decimal
	NodeID               int
	Timestamp            time.Time
}
