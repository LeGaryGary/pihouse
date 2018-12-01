package api

import (
	"encoding/json"
	"log"

	"github.com/Jordank321/pihouse/data"
	"github.com/Jordank321/pihouse/pihouseclient/speech"

	"github.com/gorilla/websocket"
)

func ConnectToServerWebsocket() {
	if websocketConnection != nil {
		return
	}

	conn, _, err := websocket.DefaultDialer.Dial("ws://"+APIAddress+"/v1/api/websocket", nil)
	if err != nil {
		log.Panicf(err.Error())
	}

	actions := []data.Action{
		data.LivingRoomLightsOn,
		data.HeatingOn,
	}

	b, err := json.Marshal(actions)
	if err != nil {
		log.Panicln(err.Error())
	}
	err = conn.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		log.Panicln(err.Error())
	}
	websocketConnection = conn
	go ReceiveCommands()
}

func ReceiveCommands() {
	for {
		_, message, err := websocketConnection.ReadMessage()
		if err != nil {
			log.Panicln(err.Error())
		}

		var action data.Action
		err = json.Unmarshal(message, &action)
		if err != nil {
			log.Panicln(err.Error())
		}

		log.Println("Oh look, the server told me to do " + action.String())
		speech.Say("Oh look, the server told me to do " + action.String())
	}
}
