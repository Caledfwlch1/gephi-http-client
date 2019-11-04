package gephi_http_client

import (
	"bytes"
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
}

func NewGephiClient(client *http.Client, host, workspace string) (GephiClient, error) {
	r, w := io.Pipe()
	out := gephiClient{
		httpClient: client,
		url:        fmt.Sprintf("http://%s/%s?operation=updateGraph", host, workspace),
		r:          r,
		w:          w,
		enc:        json.NewEncoder(w),
	}

	err := out.runPost()
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (g *gephiClient) SetClientProp(client *http.Client) {
	g.httpClient = client
}

func (g *gephiClient) runPost() error {
	emptyBuf := bytes.NewBuffer([]byte{})

	_, err := g.httpClient.Post(g.url, "application/json", emptyBuf)
	if err != nil {
		return err
	}

	go func() {
		defer g.r.Close()
		_, _ = g.httpClient.Post(g.url, "application/json", g.r)
	}()

	return nil
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
