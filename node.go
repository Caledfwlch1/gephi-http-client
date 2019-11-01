package gephi_http_client

type node struct {
	c *gclient
}

func newNode(c *gclient) Node {
	return &node{c}
}

func (n *node) NodeAdd(node interface{}) error {
	panic("implement me")
}

func (n *node) NodesAdd(node []interface{}) error {
	panic("implement me")
}

func (n *node) NodeChange(node interface{}) error {
	panic("implement me")
}

func (n *node) NodesChange(node []interface{}) error {
	panic("implement me")
}

func (n *node) NodeDelete(node interface{}) error {
	panic("implement me")
}

func (n *node) NodesDelete(node []interface{}) error {
	panic("implement me")
}

func (n *node) NodeGet(node interface{}) (interface{}, error) {
	panic("implement me")
}

func (n *node) NodesGet(node []interface{}) ([]interface{}, error) {
	panic("implement me")
}
