package control

import (
	"sort"

	"github.com/Jordank321/pihouse/data"
	"github.com/Jordank321/pihouse/pihouseserver/db"
)

type ClientController interface {
	ProcessRequest(*data.AIRequest)
	AddClient(Client)
}

type WebSocketClientController struct {
	clients      []Client
	aiRepository db.AIRepository
}

func NewWebSocketClientController(aiRepository db.AIRepository) ClientController {
	return &WebSocketClientController{
		aiRepository: aiRepository,
	}
}

func (controller WebSocketClientController) ProcessRequest(request *data.AIRequest) {
	sort.Slice(request.Intents, func(i, j int) bool {
		return request.Intents[j].Confidence.LessThan(request.Intents[i].Confidence)
	})

	for _, intent := range request.Intents {
		action := controller.aiRepository.FindAction(intent.Value)
		if action != nil {
			controller.actOnActionRequest(*action)
			break
		}
	}
}

func (controller WebSocketClientController) AddClient(client Client) {
	controller.clients = append(controller.clients, client)
}

func (controller WebSocketClientController) findClients(action data.Action) []Client {
	clients := []Client{}
	for _, client := range controller.clients {
		for _, applicableAction := range client.GetApplicableActions() {
			if applicableAction == action {
				clients = append(clients, client)
				break
			}
		}
	}
	return clients
}

func (controller WebSocketClientController) actOnActionRequest(action data.Action) {
	clients := controller.findClients(action)
	for _, client := range clients {
		client.SendAction(action)
	}
}
