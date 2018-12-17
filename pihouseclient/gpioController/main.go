package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Jordank321/pihouse/data"
	"github.com/Jordank321/pihouse/pihouseclient/api"
)

import . "github.com/ahmetb/go-linq"

func main() {
	b, err := ioutil.ReadFile("appsetting.json")
	if err != nil {
		log.Panicln(err)
	}
	settings := gpioSettings{}
	err = json.Unmarshal(b, settings)
	if err != nil {
		log.Panicln(err)
	}

	api.APIAddress = settings.ApiAddress

	actions := &[]data.Action{}
	From(settings.Mappings).Select(func(mapping interface{}) interface{} {
		return mapping.(data.ActionToGpioMapping).Action
	}).ToSlice(actions)

	actionChan := api.ConnectToServerWebsocket(actions)

	for {
		select {
		case action := <-actionChan:
			mapping := From(settings.Mappings).Where(func(mapping interface{}) bool {
				return mapping.(data.ActionToGpioMapping).Action == action
			}).First().(data.ActionToGpioMapping)
			toggleGpioPin(mapping.GpioPin)
		}
	}
}

func toggleGpioPin(pin int) {
	log.Printf("The pin %d was toggled", pin)
}
