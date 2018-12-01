package control

import (
	"encoding/json"
	"log"
	"sort"
	"strings"

	"github.com/Jordank321/pihouse/data"
	"github.com/Jordank321/pihouse/pihouseserver/db"
)

type ClientController interface {
	ProcessRequest(*data.AIRequest)
	AddClient(Client)
}

type WebSocketClientController struct {
	clients      *[]Client
	aiRepository db.AIRepository
}

func NewWebSocketClientController(aiRepository db.AIRepository) ClientController {
	return &WebSocketClientController{
		aiRepository: aiRepository,
		clients:      &[]Client{},
	}
}

func (controller *WebSocketClientController) ProcessRequest(request *data.AIRequest) {
	log.Println("Processing AI request!")

	sort.Slice(request.Intents, func(i, j int) bool {
		return request.Intents[j].Confidence.LessThan(request.Intents[i].Confidence)
	})

	actionSent := false
	for _, intent := range request.Intents {

		actions := controller.aiRepository.FindActions(intent.Value)
		for _, action := range actions {
			requiredintents := strings.Split(action.IntentValue, ",")

			actionValid := true
			for _, requiredIntent := range requiredintents {
				intentNotFound := true
				for _, presentIntent := range request.Intents {
					if requiredIntent == presentIntent.Value {
						intentNotFound = false
						break
					}
				}
				if intentNotFound {
					actionValid = false
					break
				}
			}

			if actionValid {
				log.Printf("Acting on action %s!", action.Action)
				controller.actOnActionRequest(action.Action)
				actionSent = true
				break
			}
		}
		if actionSent {
			break
		}
	}
}

func (controller *WebSocketClientController) AddClient(client Client) {
	newClients := append(*controller.clients, client)
	controller.clients = &newClients
}

func (controller *WebSocketClientController) findClients(action data.Action) []Client {
	clients := []Client{}
	log.Printf("controller has %d clients", len(*controller.clients))
	for _, client := range *controller.clients {
		actionsString, err := json.Marshal(client.GetApplicableActions())
		if err != nil {
			log.Panicln(err)
		}
		log.Println(string(actionsString))
		for _, applicableAction := range client.GetApplicableActions() {
			if applicableAction == action {
				clients = append(clients, client)
				break
			}
		}
	}
	return clients
}

func (controller *WebSocketClientController) actOnActionRequest(action data.Action) {
	clients := controller.findClients(action)
	for _, client := range clients {
		log.Printf("Sending action %s to client at %s", action, client)
		client.SendAction(action)
	}
}
