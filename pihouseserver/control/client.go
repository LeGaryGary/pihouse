package control

import (
	"github.com/gorilla/websocket"
)

type Client interface {
}

type WebSocketClient struct {
	stop       chan bool
	connection *websocket.Conn
}
