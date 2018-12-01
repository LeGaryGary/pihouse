package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/Jordank321/pihouse/data"
	"github.com/shopspring/decimal"
)

func getTemperatureAndHumidity() (decimal.Decimal, decimal.Decimal) {
	process := exec.Command("temphumidreader")
	output, err := process.Output()
	if err != nil {
		panic(err)
	}
	log.Println("test")
	log.Println(len(output))
	outString := string(output)
	log.Println(outString)
	outputs := strings.Split(outString, " ")
	temp, err := decimal.NewFromString(outputs[0])
	if err != nil {
		panic(err)
	}
	humidity, err := decimal.NewFromString(outputs[1])
	if err != nil {
		panic(err)
	}
	return temp, humidity
}

func postTemperature(val decimal.Decimal, nodeID uint) {
	read := &data.TemperatureReading{
		NodeID: nodeID,
		Value:  val,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(read)
	res, _ := http.Post("http://"+APIAddress+"/v1/api/temperature", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}

func postHumidity(val decimal.Decimal, nodeID uint) {
	read := &data.HumidityReading{
		NodeID: nodeID,
		Value:  val,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(read)
	res, _ := http.Post("http://"+APIAddress+"/v1/api/humidity", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}

// PostCurrentTemperature does what is says on the tin you twats
func PostCurrentTemperatureAndHumidity(nodeID uint) {
	temp, humid := getTemperatureAndHumidity()
	postTemperature(temp, nodeID)
	postHumidity(humid, nodeID)
}
