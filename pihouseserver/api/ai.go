package api

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"

	"github.com/Jordank321/pihouse/data"

	"github.com/jsgoecke/go-wit"

	"github.com/Jordank321/pihouse/pihouseserver/control"
	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/go-chi/chi"
)

type AIController struct {
	aiRepo           db.AIRepository
	clientController control.ClientController
	nodeRepo         db.NodeRepository
}

type AIControllerMethod func(controller *AIController, w http.ResponseWriter, r *http.Request)

func AIRoutes(aiRepo func() db.AIRepository, clientController func() control.ClientController, nodeRepo func() db.NodeRepository) (string, *chi.Mux) {
	router := chi.NewRouter()
	doMethod := func(method AIControllerMethod) func(w http.ResponseWriter, r *http.Request) {
		controller := &AIController{
			aiRepo:           aiRepo(),
			clientController: clientController(),
			nodeRepo:         nodeRepo(),
		}
		return func(w http.ResponseWriter, r *http.Request) {
			method(controller, w, r)
		}
	}
	router.Post("/outcomes/{NodeName}", doMethod((*AIController).NewWitAIOutcome))
	return "/ai", router
}

// GetReadingByID retrieves a single temperature reading by its ID
func (controller *AIController) NewWitAIOutcome(w http.ResponseWriter, r *http.Request) {
	outcomes := []wit.Outcome{}
	if err := json.NewDecoder(r.Body).Decode(&outcomes); err != nil {
		panic(err.Error())
	}

	nodeName := chi.URLParam(r, "NodeName")
	node := controller.nodeRepo.GetNodeByName(nodeName)

	for _, outcome := range outcomes {
		request := &data.AIRequest{
			Text: outcome.Text,
			Node: node,
		}

		request.Intents = []data.Intent{}

		for entityName, entityValues := range outcome.Entities {
			for _, entityValue := range entityValues {
				value := (*entityValue.Value).(string)
				request.Intents = append(request.Intents, data.Intent{
					Value:      (entityName + ":" + value),
					Confidence: decimal.NewFromFloat32(outcome.Confidence),
				})
			}
		}

		controller.aiRepo.NewWitAIOutcome(request)
		controller.clientController.ProcessRequest(request)
	}
}
