package main

import (
	"os"
	"os/signal"

	"github.com/Jordank321/pihouse/pihouseclient/api"
	"github.com/Jordank321/pihouse/pihouseclient/messageprocessing"
	"github.com/Jordank321/pihouse/pihouseclient/voice"

	"github.com/jasonlvhit/gocron"
	"github.com/jsgoecke/go-wit"
	"github.com/spf13/viper"
)

func main() {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)

	listen(shutdownChan)

	viper.AutomaticEnv()
	api.APIAddress = viper.GetString("APIAddress")

	nodeID := api.GetNodeID()

	gocron.Every(1).Minute().Do(func() { api.PostCurrentTemperature(nodeID) })
	<-gocron.Start()
}

func listen(shutdownChan <-chan os.Signal) {
	messageChan := make(chan string)
	go voice.Listen(shutdownChan, messageChan)
	intentChan := make(chan []wit.Outcome)
	go messageprocessing.GetIntent(shutdownChan, messageChan, intentChan)
}
