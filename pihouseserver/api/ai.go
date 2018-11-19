package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jordank321/pihouse/data"

	"github.com/jsgoecke/go-wit"

	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/go-chi/chi"
)

var (
	GetAIRepo func() db.AIRepository
)

func AIRoutes(getAIRepo func() db.AIRepository) *chi.Mux {
	GetAIRepo = getAIRepo
	router := chi.NewRouter()
	router.Post("/outcomes", NewWitAIOutcome)
	return router
}

// GetReadingByID retrieves a single temperature reading by its ID
func NewWitAIOutcome(w http.ResponseWriter, r *http.Request) {
	outcomes := []wit.Outcome{}
	if err := json.NewDecoder(r.Body).Decode(&outcomes); err != nil {
		panic(err.Error())
	}

	repo := GetAIRepo()
	for _, outcome := range outcomes {
		request := &data.AIRequest{
			Text: outcome.Text,
		}

		request.Intents = []data.Intent{}

		for entityName, entityValues := range outcome.Entities {
			for _, entityValue := range entityValues {
				value := (*entityValue.Value).(string)
				request.Intents = append(request.Intents, data.Intent{
					Value: (entityName + ":" + value),
				})
			}
		}

		repo.NewWitAIOutcome(request)
	}
}
