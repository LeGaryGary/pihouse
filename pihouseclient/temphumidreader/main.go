package main

import (
	"fmt"

	dht "github.com/d2r2/go-dht"
)

func main() {
	temperature, humidity, _, _ := dht.ReadDHTxxWithRetry(dht.DHT11, 2, false, 10)
	fmt.Printf("%f %f", temperature, humidity)
}
