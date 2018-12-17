package control

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jordank321/pihouse/data"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	clientController ClientController
}

type WebSocketControllerMethod func(controller *WebSocketController, w http.ResponseWriter, r *http.Request)

func WebSocketRoutes(getClientController func() ClientController) (string, *chi.Mux) {
	router := chi.NewRouter()
	doMethod := func(method WebSocketControllerMethod) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			controller := &WebSocketController{
				clientController: getClientController(),
			}
			method(controller, w, r)
		}
	}
	router.Get("/", doMethod((*WebSocketController).InitiateWebsocket))
	return "/websocket", router
}

var upgrader = websocket.Upgrader{} // use default options

func (controller *WebSocketController) InitiateWebsocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	stop := make(chan bool)
	client := &WebSocketClient{
		connection: c,
		stop:       stop,
	}
	c.SetCloseHandler(func(mt int, message string) error {
		client.SetAsClosed()
		return nil
	})

	controller.clientController.AddClient(client)
	for {
		// select {
		// case <-stop:
		// 	client.SetAsClosed()
		// 	return
		// default:
		mt, message, _ := c.ReadMessage()
		if err != nil {
			log.Println("websocket read:", err)
			return
		}
		log.Printf("websocket recv %d from %s: %s", mt, c.RemoteAddr().String(), message)
		if mt == websocket.BinaryMessage {

			actions := []data.Action{}
			err := json.Unmarshal(message, &actions)
			if err != nil {
				log.Panicln(err)
			}
			client.applicableActions = actions
		}
		if mt == -1 {
			break
		}
		//}

	}
}
