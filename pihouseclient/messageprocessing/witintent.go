package messageprocessing

import (
	"log"
	"os"

	wit "github.com/jsgoecke/go-wit"
)

var client = wit.NewClient(os.Getenv("WIT_ACCESS_TOKEN"))

func GetIntent(shutdownChan <-chan os.Signal, msg <-chan string, intent chan []wit.Outcome) {
	for {
		select {
		case <-shutdownChan:
			return
		default:
			// Process a text message
			request := &wit.MessageRequest{}
			request.Query = <-msg
			result, err := client.Message(request)
			if err != nil {
				log.Printf(err.Error())
			}
			intent <- result.Outcomes
		}
	}
}

func ProcessIntent(shutdownChan <-chan os.Signal, intent <-chan []wit.Outcome) {
	for {
		outcomes := <-intent
		for _, outcome := range outcomes {
			log.Printf("%+v", outcome)
		}
	}
}
