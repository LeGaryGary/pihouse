package api

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Jordank321/pihouse/data"
	"github.com/Jordank321/pihouse/pihouseclient/speech"

	"github.com/gorilla/websocket"
)

func ConnectToServerWebsocket(actions *[]data.Action) <-chan data.Action {
	if websocketConnection != nil {
		return nil
	}

	conn, _, err := websocket.DefaultDialer.Dial("ws://"+APIAddress+"/v1/api/websocket", nil)
	if err != nil {
		log.Panicf(err.Error())
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
	actionChan := make(chan data.Action)
	go ReceiveCommands(actionChan)
	return actionChan
}

func DisconnectFromWebsocket() {
	if websocketConnection == nil {
		return
	}

	websocketConnection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(time.Second)
	websocketConnection.Close()
}

func ReceiveCommands(actionChan chan<- data.Action) {
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

		actionChan <- action
		log.Println("Oh look, the server told me to do " + action.String())
		speech.Say("Oh look, the server told me to do " + action.String())
	}
}
