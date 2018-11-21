package control

import (
	"github.com/Jordank321/pihouse/data"
)

type ClientController interface {
	ProcessRequest(*data.AIRequest)
	AddClient(Client)
}

type WebSocketClientController struct {
	clients []Client
}

func (controller WebSocketClientController) ProcessRequest(request *data.AIRequest) {
	// fuck all
}

func (controller WebSocketClientController) AddClient(client Client) {
	controller.clients = append(controller.clients, client)
}
