package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Jordank321/pihouse/data"
)

// GetNodeID retrives the node ID for the hostname from the API
func GetNodeID() uint {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}
	resp, err := http.Get("http://" + APIAddress + "/v1/api/node/" + hostname)
	if err != nil {
		panic(err.Error())
	}
	if resp.StatusCode == 404 {
		return newNode(hostname)
	}
	node := &data.Node{}
	if err := json.NewDecoder(resp.Body).Decode(node); err != nil {
		panic(err.Error())
	}
	return node.ID
}

func newNode(hostname string) uint {
	node := &data.Node{
		Name: hostname,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(node)
	res, _ := http.Post("http://"+APIAddress+"/v1/api/node", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
	respNode, err := http.Get("http://" + APIAddress + "/v1/api/node/" + hostname)
	if err != nil {
		panic(err.Error())
	}
	if respNode.StatusCode == 404 {
		panic("wtf?")
	}
	node = &data.Node{}
	if err := json.NewDecoder(respNode.Body).Decode(node); err != nil {
		panic(err.Error())
	}
	return node.ID
}
