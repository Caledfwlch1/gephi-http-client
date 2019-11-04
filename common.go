package gephi_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type gephiClient struct {
	httpClient *http.Client
	enc        *json.Encoder
}

func NewGephiClient(client *http.Client, host, workspace string) (GephiClient, error) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	url := fmt.Sprintf("http://%s/%s?operation=updateGraph", host, workspace)
	_, err := client.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}

	return &gephiClient{
		httpClient: client,
		enc:        json.NewEncoder(buf),
	}, nil
}

func (g *gephiClient) SetClientProp(client *http.Client) {
	g.httpClient = client
}

func (g *gephiClient) marshal(operation string, obj ...interface{}) error {
	var err error
	m := make(map[string]interface{})

	for _, o := range obj {
		if n, ok := o.(Node); ok {
			err = n.validate()
			if err != nil {
				return err
			}
			m[operation] = n
		} else if e, ok := o.(Edge); ok {
			err = e.validate()
			if err != nil {
				return err
			}
			m[operation] = e
		} else {
			return fmt.Errorf("unknown type of object")
		}

		err = g.enc.Encode(m)
		if err != nil {
			return err
		}
		delete(m, operation)
	}

	return nil
}
