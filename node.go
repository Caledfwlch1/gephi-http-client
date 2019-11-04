package gephi_http_client

import (
	"bytes"
	"encoding/json"
)

type node struct {
	enc *json.Encoder
}

func newNode(buf *bytes.Buffer) NodeOperations {
	return &node{json.NewEncoder(buf)}
}

func (n *node) NodeAdd(node ...Node) error {
	return n.marshal("an", node)
}

func (n *node) NodeChange(node ...Node) error {
	return n.marshal("cn", node)
}

func (n *node) NodeDelete(node ...Node) error {
	return n.marshal("dn", node)
}

func (n *node) NodeGet(node ...Node) (interface{}, error) {
	panic("implement me")
}

func (n *node) marshal(operation string, o ...interface{}) error {
	return marshal(n.enc, operation, o)
}
