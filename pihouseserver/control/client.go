package control

import (
	"encoding/json"
	"log"

	"github.com/Jordank321/pihouse/data"
	"github.com/gorilla/websocket"
)

type Client interface {
	GetApplicableActions() []data.Action
	SendAction(action data.Action)
	String() string
	SetAsClosed()
}

type WebSocketClient struct {
	stop              chan bool
	connection        *websocket.Conn
	applicableActions []data.Action
	closed            bool
}

func (client *WebSocketClient) GetApplicableActions() []data.Action {
	return client.applicableActions
}

func (client *WebSocketClient) SendAction(action data.Action) {
	b, err := json.Marshal(action)
	if err != nil {
		log.Panicln(err.Error())
	}
	err = client.connection.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		log.Print(err.Error())
	}
}

func (client *WebSocketClient) String() string {
	return client.connection.RemoteAddr().String()
}

func (client *WebSocketClient) SetAsClosed() {
	client.closed = true
	client.connection.Close()
}

//<3
