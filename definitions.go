package gephi_http_client

import (
	"bytes"
	"fmt"
	"net/http"
)

type GephiClient struct {
	NodeOperations
	EdgeOperations
}

func NewGephiClient(client *http.Client, host, workspace string) (*GephiClient, error) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	url := fmt.Sprintf("http://%s/%s?operation=updateGraph", host, workspace)
	_, err := client.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}

	return &GephiClient{
		NodeOperations: newNode(buf),
		EdgeOperations: newEdge(buf),
	}, nil
}

type NodeOperations interface {
	NodeAdd(node ...Node) error
	NodeChange(node ...Node) error
	NodeDelete(node ...Node) error
}

type EdgeOperations interface {
	EdgeAdd(edge ...Edge) error
	EdgeChange(edge ...Edge) error
	EdgeDelete(edge ...Edge) error
}

type Node struct {
	Id    string
	Lable string
	X     float64
	Y     float64
	Size  int
	Prop  map[string]string
}

type Edge struct {
	Id       string
	Source   string
	Target   string
	Directed string
	Weight   string
}
