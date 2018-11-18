package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Jordank321/pihouse/data"
	"github.com/shopspring/decimal"

	"github.com/d2r2/go-dht"
)

func getTemperatureAndHumidity() (float32, float32) {
	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, 2, true, 10)

	if err != nil {
		log.Println(err.Error())
	}

	return temperature, humidity
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
	postTemperature(decimal.NewFromFloat32(temp), nodeID)
	postHumidity(decimal.NewFromFloat32(humid), nodeID)
}
