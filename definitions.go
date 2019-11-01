package gephi_http_client

import (
	"bytes"
	"fmt"
	"net/http"
)

type GephiClient struct {
	Node
	Edge
}

func NewGephiClient(client *http.Client, host, workspace string) (*GephiClient, error) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	url := fmt.Sprintf("http://%s/%s?operation=updateGraph", host, workspace)
	_, err := client.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}

	return &GephiClient{
		Node: newNode(buf),
		Edge: newEdge(buf),
	}, nil
}

type Node interface {
	NodeAdd(node ...interface{}) error
	NodeChange(node ...interface{}) error
	NodeDelete(node ...interface{}) error
}

type Edge interface {
	EdgeAdd(edge ...interface{}) error
	EdgeChange(edge ...interface{}) error
	EdgeDelete(edge ...interface{}) error
}
