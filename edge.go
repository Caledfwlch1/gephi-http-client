package gephi_http_client

type edge struct {
	c *gclient
}

func newEdge(c *gclient) Edge {
	return &edge{c}
}

func (e *edge) EdgeAdd(edge interface{}) error {
	panic("implement me")
}

func (e *edge) EdgesAdd(edge []interface{}) error {
	panic("implement me")
}

func (e *edge) EdgeChange(edge interface{}) error {
	panic("implement me")
}

func (e *edge) EdgesChange(edge []interface{}) error {
	panic("implement me")
}

func (e *edge) EdgeDelete(edge interface{}) error {
	panic("implement me")
}

func (e *edge) EdgesDelete(edge []interface{}) error {
	panic("implement me")
}

func (e *edge) EdgeGet(edge interface{}) (interface{}, error) {
	panic("implement me")
}

func (e *edge) EdgesGet(edge []interface{}) ([]interface{}, error) {
	panic("implement me")
}
