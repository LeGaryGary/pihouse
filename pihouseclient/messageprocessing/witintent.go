package messageprocessing

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Jordank321/pihouse/pihouseclient/api"

	wit "github.com/jsgoecke/go-wit"
)

var client *wit.Client

func SetToken(token string) {
	client = wit.NewClient(token)
}

func GetIntent(shutdownChan <-chan os.Signal, msg <-chan string, intent chan []wit.Outcome) {
	for {
		select {
		case <-shutdownChan:
			return
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

func ProcessIntent(shutdownChan <-chan os.Signal, intent <-chan []wit.Outcome) {
	for {
		select {
		case <-shutdownChan:
			return
		case outcomes := <-intent:
			data, _ := json.MarshalIndent(outcomes, "", "    ")
			log.Println(string(data[:]))
			api.PostAIIntent(outcomes)
		}
	}
}
