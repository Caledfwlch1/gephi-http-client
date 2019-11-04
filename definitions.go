package gephi_http_client

import (
	"net/http"
)

type GephiClient interface {
	SetClientProp(client *http.Client)

	NodeAdd(node ...Node) error
	NodeChange(node ...Node) error
	NodeDelete(node ...Node) error

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
	Directed bool
	Weight   int
}
