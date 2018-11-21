package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/Jordank321/pihouse/data"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

type NodeController struct {
	nodeRepo db.NodeRepository
}

type NodeControllerMethod func(controller *NodeController, w http.ResponseWriter, r *http.Request)

func NodeRoutes(getNodeRepo func() db.NodeRepository) (string, *chi.Mux) {
	router := chi.NewRouter()
	doMethod := func(method NodeControllerMethod) func(w http.ResponseWriter, r *http.Request) {
		controller := &NodeController{
			nodeRepo: getNodeRepo(),
		}
		return func(w http.ResponseWriter, r *http.Request) {
			method(controller, w, r)
		}
	}
	router.Get("/", doMethod((*NodeController).GetAllNodes))
	router.Post("/", doMethod((*NodeController).CreateNode))
	router.Get("/{NodeName}", doMethod((*NodeController).GetNodeByName))
	return "/node", router
}

// GetNodeByName retrieves a single node reading by its name
func (controller *NodeController) GetNodeByName(w http.ResponseWriter, r *http.Request) {
	nodeName := chi.URLParam(r, "NodeName")
	node := controller.nodeRepo.GetNodeByName(nodeName)
	if node == nil {
		http.NotFound(w, r)
		return
	}
	render.JSON(w, r, node)
}

// GetAllNodes retrieves all nodes
func (controller *NodeController) GetAllNodes(w http.ResponseWriter, r *http.Request) {
	nodes := controller.nodeRepo.GetAllNodes()
	render.JSON(w, r, nodes)
}

// CreateNode creates a new temperature reading
func (controller *NodeController) CreateNode(w http.ResponseWriter, r *http.Request) {
	node := &data.Node{}
	if err := json.NewDecoder(r.Body).Decode(node); err != nil {
		panic(err.Error())
	}

	controller.nodeRepo.AddNode(node)

	response := make(map[string]string)
	response["message"] = "Success: the node has been added"
	render.JSON(w, r, response)
}
