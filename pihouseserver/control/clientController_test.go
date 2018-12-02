package control

import (
	"reflect"
	"testing"

	"github.com/Jordank321/pihouse/data"
)

type TestAIRepository struct {
}

func (testRepo *TestAIRepository) NewWitAIOutcome(request *data.AIRequest) {

}
func (testRepo *TestAIRepository) FindActions(intentValue string) []data.ActionMaping {
	return []data.ActionMaping{}
}

func TestWebSocketClientController_ProcessRequest(t *testing.T) {
	type args struct {
		request *data.AIRequest
	}
	tests := []struct {
		name       string
		controller *WebSocketClientController
		args       args
	}{
		{
			"AI request for living room lights on with single mapping",
			&WebSocketClientController{
				aiRepository: &TestAIRepository{},
			},
			args{
				request: &data.AIRequest{
					Intents: []data.Intent{
						data.Intent{
							Value: "intent:lights_living_room",
						},
						data.Intent{
							Value: "on_off:on",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.ProcessRequest(tt.args.request)
		})
	}
}

func TestWebSocketClientController_AddClient(t *testing.T) {
	type args struct {
		client Client
	}
	tests := []struct {
		name       string
		controller *WebSocketClientController
		args       args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.AddClient(tt.args.client)
		})
	}
}

func TestWebSocketClientController_findClients(t *testing.T) {
	type args struct {
		action data.Action
	}
	tests := []struct {
		name       string
		controller *WebSocketClientController
		args       args
		want       []Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.controller.findClients(tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebSocketClientController.findClients() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebSocketClientController_actOnActionRequest(t *testing.T) {
	type args struct {
		action data.Action
	}
	tests := []struct {
		name       string
		controller *WebSocketClientController
		args       args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.actOnActionRequest(tt.args.action)
		})
	}
}
