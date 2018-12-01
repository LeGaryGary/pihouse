package main

import (
	"github.com/Jordank321/pihouse/pihouseclient/api"

	"github.com/jasonlvhit/gocron"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	api.APIAddress = viper.GetString("API_ADDRESS")

	nodeID := api.GetNodeID()

	api.ConnectToServerWebsocket()

	gocron.Every(1).Minute().Do(func() { api.PostCurrentTemperatureAndHumidity(nodeID) })
	<-gocron.Start()
}
