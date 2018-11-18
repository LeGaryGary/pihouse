package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/Jordank321/pihouse/data"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

var (
	GetNodeRepo func() db.NodeRepository
)

func NodeRoutes(getNodeRepo func() db.NodeRepository) *chi.Mux {
	GetNodeRepo = getNodeRepo
	router := chi.NewRouter()
	router.Get("/", GetAllNodes)
	router.Post("/", CreateNode)
	router.Get("/{NodeName}", GetNodeByName)
	return router
}

// GetNodeByName retrieves a single node reading by its name
func GetNodeByName(w http.ResponseWriter, r *http.Request) {
	nodeName := chi.URLParam(r, "NodeName")
	repo := GetNodeRepo()
	node := repo.GetNodeByName(nodeName)
	if node == nil {
		http.NotFound(w, r)
		return
	}
	render.JSON(w, r, node)
}

// GetAllNodes retrieves all nodes
func GetAllNodes(w http.ResponseWriter, r *http.Request) {
	repo := GetNodeRepo()
	nodes := repo.GetAllNodes()
	render.JSON(w, r, nodes)
}

// CreateNode creates a new temperature reading
func CreateNode(w http.ResponseWriter, r *http.Request) {
	node := &data.Node{}
	if err := json.NewDecoder(r.Body).Decode(node); err != nil {
		panic(err.Error())
	}

	repo := GetNodeRepo()
	repo.AddNode(node)

	response := make(map[string]string)
	response["message"] = "Success: the node has been added"
	render.JSON(w, r, response)
}
