package messageprocessing

import (
	"encoding/json"
	"log"

	"github.com/Jordank321/pihouse/pihouseclient/api"

	wit "github.com/jsgoecke/go-wit"
)

var client *wit.Client

func SetToken(token string) {
	client = wit.NewClient(token)
}

func GetIntent(msg <-chan string, intent chan []wit.Outcome) {
	for {
		select {
		case msgIn := <-msg:
			// Process a text message
			request := &wit.MessageRequest{}
			request.Query = msgIn
			result, err := client.Message(request)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
			intent <- result.Outcomes
		}
	}
}

func ProcessIntent(intent <-chan []wit.Outcome) {
	for {
		select {
		case outcomes := <-intent:
			data, _ := json.MarshalIndent(outcomes, "", "    ")
			log.Println(string(data[:]))
			api.PostAIIntent(outcomes)
		}
	}
}
