package control

import (
	"log"
	"net/http"

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
		controller := &WebSocketController{
			clientController: getClientController(),
		}
		return func(w http.ResponseWriter, r *http.Request) {
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
	controller.clientController.AddClient(&WebSocketClient{
		connection: c,
		stop:       stop,
	})
	for {
		select {
		case <-stop:
			return
		default:
			mt, message, _ := c.ReadMessage()
			if err != nil {
				log.Println("websocket read:", err)
				return
			}
			log.Printf("websocket recv from %s: %s", c.RemoteAddr().String(), message)
			c.WriteMessage(mt, message)
			if err != nil {
				log.Println("websocket write:", err)
				return
			}
		}

	}
}
