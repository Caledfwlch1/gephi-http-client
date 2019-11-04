package gephi_http_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type gephiClient struct {
	httpClient *http.Client
	url        string
	r          *io.PipeReader
	w          *io.PipeWriter
	enc        *json.Encoder
	postErr    error
}

func NewGephiClient(client *http.Client, host, workspace string) (GephiClient, error) {
	r, w := io.Pipe()
	return &gephiClient{
		httpClient: client,
		url:        fmt.Sprintf("http://%s/%s?operation=updateGraph", host, workspace),
		r:          r,
		w:          w,
		enc:        json.NewEncoder(w),
	}, nil
}

func (g *gephiClient) SetClientProp(client *http.Client) {
	g.httpClient = client
}

func (g *gephiClient) RunPost() {
	go func() {
		defer g.r.Close()
		_, g.postErr = g.httpClient.Post(g.url, "application/json", g.r)
	}()
}

func (g *gephiClient) Close() error {
	return g.w.Close()
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
