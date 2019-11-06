package gephi_http_client

import (
	"encoding/json"
	"fmt"
)

func (g *gephiClient) EdgeAdd(edge ...Edge) error {
	return g.marshal("ae", edge)
}

func (g *gephiClient) EdgeChange(edge ...Edge) error {
	return g.marshal("ce", edge)
}

func (g *gephiClient) EdgeDelete(edge ...Edge) error {
	return g.marshal("de", edge)
}

func (g *gephiClient) EdgeGet(edge ...Edge) (interface{}, error) {
	panic("implement me")
}

func (e *Edge) validate() error {
	//if e.Id == "" {
	//	return fmt.Errorf("edge %s has empty Id", e)
	//}
	//if e.Source == "" {
	//	return fmt.Errorf("edge %s has empty source", e)
	//}
	//if e.Target == "" {
	//	return fmt.Errorf("edge %s has empty target", e)
	//}
	return nil
}

func (e Edge) String() string {
	return fmt.Sprintf("%#v", e)
}

func (e *Edge) MarshalJSON() ([]byte, error) {
	m := map[string]map[string]string(*e)
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return append(b, '\r'), nil
}
