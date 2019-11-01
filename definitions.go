package gephi_http_client

import "net/http"

type GephiClient struct {
	Client *http.Client
	Node
	Edge
}

type Node interface {
	NodeAdd(node interface{}) error
	NodesAdd(node []interface{}) error
	NodeChange(node interface{}) error
	NodesChange(node []interface{}) error
	NodeDelete(node interface{}) error
	NodesDelete(node []interface{}) error
	NodeGet(node interface{}) (interface{}, error)
	NodesGet(node []interface{}) ([]interface{}, error)
}

type Edge interface {
	EdgeAdd(edge interface{}) error
	EdgesAdd(edge []interface{}) error
	EdgeChange(edge interface{}) error
	EdgesChange(edge []interface{}) error
	EdgeDelete(edge interface{}) error
	EdgesDelete(edge []interface{}) error
	EdgeGet(edge interface{}) (interface{}, error)
	EdgesGet(edge []interface{}) ([]interface{}, error)
}
