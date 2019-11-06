package gephi_http_client

import (
	"encoding/json"
	"fmt"
)

func (g *gephiClient) NodeAdd(node ...Node) error {
	return g.marshal("an", node)
}

func (g *gephiClient) NodeChange(node ...Node) error {
	return g.marshal("cn", node)
}

func (g *gephiClient) NodeDelete(node ...Node) error {
	return g.marshal("dn", node)
}

func (g *gephiClient) NodeGet(node ...Node) (interface{}, error) {
	panic("implement me")
}

func (n *Node) validate() error {
	//if n.Id == "" {
	//	return fmt.Errorf("node %s has empty Id", n)
	//}
	//if n.Lable == "" {
	//	n.Lable = n.Id
	//}
	//if n.Size <= 0 {
	//	n.Size = 1
	//}
	return nil
}

func (n Node) String() string {
	return fmt.Sprintf("%#v", n)
}

func (n *Node) MarshalJSON() ([]byte, error) {
	m := map[string]map[string]string(*n)
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return append(b, '\r'), nil
}
