package main

import (
	"github.com/Jordank321/pihouse/data"
)

type gpioSettings struct {
	ApiAddress string                     `json:"API_ADDRESS"`
	Mappings   []data.ActionToGpioMapping `json:"actionToGpios"`
}
