package api

import (
	"github.com/gorilla/websocket"
)

var (
	// APIAddress is the static address the package uses to contact the PiHouse API
	APIAddress          string
	websocketConnection *websocket.Conn
)
