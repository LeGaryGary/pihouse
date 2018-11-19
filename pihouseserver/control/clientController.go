package control

import (
	"github.com/Jordank321/pihouse/data"
)

type ClientController interface {
	ProcessRequest(*data.AIRequest)
}

type WebSocketClientController struct {
}

func (controller WebSocketClientController) ProcessRequest(*data.AIRequest) {

}
