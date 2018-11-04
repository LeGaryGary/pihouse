package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"pihouse/data"

	"github.com/spf13/viper"

	"github.com/jasonlvhit/gocron"
	"github.com/shopspring/decimal"
)

var (
	apiAddress string
)

func getTemperature() decimal.Decimal {
	val, err := decimal.NewFromString("45.5")
	if err != nil {
		panic(err.Error())
	}
	return val
}

func postTemperature(val decimal.Decimal, nodeID uint) {
	read := &data.TemperatureReading{
		NodeID: nodeID,
		Value:  val,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(read)
	res, _ := http.Post("http://"+apiAddress+"/v1/api/temperature", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}

func postCurrentTemperature(nodeID uint) {
	postTemperature(getTemperature(), nodeID)
}

func getNodeID() uint {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}
	resp, err := http.Get("http://" + apiAddress + "/v1/api/node/" + hostname)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode == 404 {
		return newNode(hostname)
	}
	node := &data.Node{}
	if err := json.NewDecoder(resp.Body).Decode(node); err != nil {
		panic(err.Error())
	}
	return node.ID
}

func newNode(hostname string) uint {
	node := &data.Node{
		Name: hostname,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(node)
	res, _ := http.Post("http://"+apiAddress+"/v1/api/node", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
	respNode, err := http.Get("http://" + apiAddress + "/v1/api/node/" + hostname)
	if err != nil {
		panic(err.Error())
	}
	if respNode.StatusCode == 404 {
		panic("wtf?")
	}
	node = &data.Node{}
	if err := json.NewDecoder(respNode.Body).Decode(node); err != nil {
		panic(err.Error())
	}
	return node.ID
}

func main() {
	viper.AutomaticEnv()
	apiAddress = viper.GetString("ApiAddress")

	nodeID := getNodeID()

	gocron.Every(1).Second().Do(func() { postCurrentTemperature(nodeID) })
	<-gocron.Start()
}
