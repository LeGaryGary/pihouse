package main

import (
	"fmt"

	"github.com/Jordank321/pihouse/pihouseclient/temphumidreader/dht"
)

func main() {
	temperature, humidity, _, _ := dht.ReadDHTxxWithRetry(dht.DHT11, 2, true, 10)
	fmt.Printf("%f %f", temperature, humidity)
}
