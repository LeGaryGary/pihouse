package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Jordank321/pihouse/data"
	"github.com/shopspring/decimal"

	"github.com/d2r2/go-dht"
)

func getTemperature() float32 {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)

	// cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// if err := cmd.Run(); err != nil {
	// 	panic(err.Error())
	// }
	// outString := out.String()
	// tempString := outString[strings.IndexByte(outString, '=')+1 : strings.IndexByte(outString, '\'')]
	// val, err := decimal.NewFromString(tempString)
	// if err != nil {
	// 	panic(err.Error())
	// }
	return temperature
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

// PostCurrentTemperature does what is says on the tin you twats
func PostCurrentTemperature(nodeID uint) {
	postTemperature(decimal.NewFromFloat32(getTemperature()), nodeID)
}
