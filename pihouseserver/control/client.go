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
}

type WebSocketClient struct {
	stop              chan bool
	connection        *websocket.Conn
	applicableActions []data.Action
}

func (client WebSocketClient) GetApplicableActions() []data.Action {
	return client.applicableActions
}

func (client WebSocketClient) SendAction(action data.Action) {
	b, err := json.Marshal(action)
	if err != nil {
		log.Panicln(err.Error())
	}
	client.connection.WriteMessage(websocket.BinaryMessage, b)
}

//<3
