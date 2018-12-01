package control

import (
	"reflect"
	"testing"

	"github.com/Jordank321/pihouse/data"
)

import . "github.com/ahmetb/go-linq"

type TestAIRepository struct {
	mappings []data.ActionMaping
}

func (testRepo *TestAIRepository) NewWitAIOutcome(request *data.AIRequest) {

}
func (testRepo *TestAIRepository) FindActions(intentValue string) []data.ActionMaping {
	return testRepo.mappings
}

type TestClient struct {
	applicableActions []data.Action
	actionsSent       []data.Action
}

func (client *TestClient) GetApplicableActions() []data.Action {
	return client.applicableActions
}

func (client *TestClient) SendAction(action data.Action) {
	client.actionsSent = append(client.actionsSent, action)
}

func (client *TestClient) SetAsClosed() {
}

func (client *TestClient) String() string {
	return "This is a test!"
}

func TestWebSocketClientController_ProcessRequest(t *testing.T) {
	type args struct {
		request *data.AIRequest
	}
	type result struct {
		sentActions []data.Action
	}

	lightsClent := &TestClient{
		applicableActions: []data.Action{
			data.LivingRoomLightsOn,
		},
	}
	heatingClient := &TestClient{
		applicableActions: []data.Action{
			data.HeatingOn,
		},
	}

	tests := []struct {
		name       string
		client     *TestClient
		controller *WebSocketClientController
		args       args
		result     result
	}{
		{
			"AI request for living room lights on with single mapping",
			lightsClent,
			&WebSocketClientController{
				aiRepository: &TestAIRepository{
					mappings: []data.ActionMaping{
						data.ActionMaping{
							Action:      data.LivingRoomLightsOn,
							IntentValue: "intent:lights_living_room,on_off:on",
						},
					},
				},
				clients: &[]Client{
					lightsClent,
				},
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
			result{
				sentActions: []data.Action{
					data.LivingRoomLightsOn,
				},
			},
		},
		{
			"AI request for heating on with multiple mappings",
			heatingClient,
			&WebSocketClientController{
				aiRepository: &TestAIRepository{
					mappings: []data.ActionMaping{
						data.ActionMaping{
							Action:      data.HeatingOn,
							IntentValue: "intent:heating,on_off:on",
						},
						data.ActionMaping{
							Action:      data.HeatingOff,
							IntentValue: "intent:heating,on_off:off",
						},
					},
				},
				clients: &[]Client{
					heatingClient,
				},
			},
			args{
				request: &data.AIRequest{
					Intents: []data.Intent{
						data.Intent{
							Value: "intent:heating",
						},
						data.Intent{
							Value: "on_off:on",
						},
					},
				},
			},
			result{
				sentActions: []data.Action{
					data.HeatingOn,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.ProcessRequest(tt.args.request)
			missingSentActions := []data.Action{}
			From(tt.result.sentActions).Where(func(sa interface{}) bool {
				return !From(tt.client.actionsSent).Contains(sa.(data.Action))
			}).ToSlice(&missingSentActions)
			for _, missingAction := range missingSentActions {
				t.Errorf("Expected sent actions to contain %s", missingAction)
			}
			if len(tt.result.sentActions) != len(tt.client.actionsSent) {
				t.Error("Mismatch in sent actions!")
			}
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
