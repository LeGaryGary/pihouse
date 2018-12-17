package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"

	"github.com/Jordank321/pihouse/data"
	"github.com/Jordank321/pihouse/pihouseclient/api"

	. "github.com/ahmetb/go-linq"
)

var pinStates map[int]*pinState

type pinState struct {
	initialisedForOutput bool
	isOutputing          bool
}

func main() {
	pinStates = make(map[int]*pinState)
	b, err := ioutil.ReadFile("appsettings.json")
	if err != nil {
		log.Panicln(err)
	}
	settings := gpioSettings{}
	err = json.Unmarshal(b, &settings)
	if err != nil {
		log.Panicln(err)
	}

	api.APIAddress = settings.ApiAddress

	actions := &[]data.Action{}
	From(settings.Mappings).Select(func(mapping interface{}) interface{} {
		return mapping.(data.ActionToGpioMapping).Action
	}).ToSlice(actions)

	actionChan := api.ConnectToServerWebsocket(actions)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)

	for {
		select {
		case <-shutdownChan:
			api.DisconnectFromWebsocket()
			break
		case action := <-actionChan:
			mapping := From(settings.Mappings).Where(func(mapping interface{}) bool {
				return mapping.(data.ActionToGpioMapping).Action == action
			}).First().(data.ActionToGpioMapping)
			toggleGpioPin(mapping.GpioPin)
		}
	}
}

func toggleGpioPin(pin int) {
	state := pinStates[pin]
	if state == nil {
		state = new(pinState)
		pinStates[pin] = state
	}

	if !state.initialisedForOutput {
		err := exec.Command("gpio -g mode " + strconv.Itoa(pin) + " out").Run()
		if err != nil {
			log.Panic(err)
		}
		state.initialisedForOutput = true
	}

	output := "ERR"

	if state.isOutputing {
		output = "0"
	} else {
		output = "1"
	}

	err := exec.Command("gpio -g write " + strconv.Itoa(pin) + " " + output).Run()
	if err != nil {
		log.Panic(err)
	}

	state.isOutputing = !state.isOutputing

	log.Printf("The pin %d was toggled", pin)

}
