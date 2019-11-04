package gephi_http_client

import (
	"bytes"
	"encoding/json"
)

type edge struct {
	enc *json.Encoder
}

func newEdge(buf *bytes.Buffer) EdgeOperations {
	return &edge{json.NewEncoder(buf)}
}

func (e *edge) EdgeAdd(edge ...Edge) error {
	return e.marshal("ae", edge)
}

func (e *edge) EdgeChange(edge ...Edge) error {
	return e.marshal("ce", edge)
}

func (e *edge) EdgeDelete(edge ...Edge) error {
	return e.marshal("de", edge)
}

func (e *edge) EdgeGet(edge ...Edge) (interface{}, error) {
	panic("implement me")
}

func (e *edge) marshal(operation string, o ...interface{}) error {
	return marshal(e.enc, operation, o)
}
