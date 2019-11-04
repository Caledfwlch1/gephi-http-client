package gephi_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type gephiClient struct {
	w   *io.PipeWriter
	enc *json.Encoder
}

func NewGephiClient(client *http.Client, host, workspace string) (GephiClient, error) {
	switch "" {
	case host:
		return nil, fmt.Errorf("host is empty")
	case workspace:
		return nil, fmt.Errorf("workspace is empty")
	}

	if client == nil {
		client = http.DefaultClient
	}

	url := fmt.Sprintf("http://%s/%s?operation=updateGraph", host, workspace)
	emptyBuf := bytes.NewBuffer([]byte{})

	_, err := client.Post(url, "application/json", emptyBuf)
	if err != nil {
		return nil, err
	}

	r, w := io.Pipe()
	go func() {
		defer r.Close()
		_, _ = client.Post(url, "application/json", r)
	}()

	return &gephiClient{
		w:   w,
		enc: json.NewEncoder(w),
	}, nil
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
