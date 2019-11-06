package gephi_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type gephiClient struct {
	w   io.WriteCloser
	enc json.Encoder
}

func NewGephiClient(client *http.Client, host, workspace string, r io.ReadCloser, w io.WriteCloser) (GephiClient, error) {
	//func NewGephiClient(client *http.Client, host, workspace string, r *io.PipeReader, w *io.PipeWriter) (GephiClient, error) {
	switch "" {
	case host:
		return nil, fmt.Errorf("host is empty")
	case workspace:
		return nil, fmt.Errorf("workspace is empty")
	}

	if r == nil || w == nil {
		return nil, fmt.Errorf("reader or writer is nil")
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

	//r, w := io.Pipe()
	go func() {
		defer r.Close()
		_, err = client.Post(url, "application/json", r)
		if err != nil {
			log.Println("PUT error:", err)
		}
	}()

	return &gephiClient{
		w:   w,
		enc: *json.NewEncoder(w),
	}, nil
}

func (g *gephiClient) Close() error {
	return g.w.Close()
}

func (g *gephiClient) marshal(operation string, obj interface{}) error {
	var err error
	m := make(map[string]interface{})

	if nodes, ok := obj.([]Node); ok {
		for _, n := range nodes {
			err = n.validate()
			if err != nil {
				return err
			}
			m[operation] = n

			b, ee := json.Marshal(m) // TODO: remove
			fmt.Println("******** ", b, ee)

			err = g.enc.Encode(m)
			if err != nil {
				return err
			}
			_, err = g.w.Write([]byte{'\r'})
			if err != nil {
				return err
			}
			delete(m, operation)

		}
		return nil

	} else if edges, ok := obj.([]Edge); ok {
		for _, e := range edges {
			err = e.validate()
			if err != nil {
				return err
			}
			m[operation] = e

			b, ee := json.Marshal(&m) // TODO: remove
			fmt.Println("******** ", string(b), ee)

			err = g.enc.Encode(m)
			if err != nil {
				return err
			}
			_, err = g.w.Write([]byte{'\r'})
			if err != nil {
				return err
			}
			delete(m, operation)
		}
		return nil

	}
	return fmt.Errorf("unknown type of object")

	//for _, o := range obj {
	//	if n, ok := o.(Node); ok {
	//		err = n.validate()
	//		if err != nil {
	//			return err
	//		}
	//		m[operation] = n
	//	} else if e, ok := o.(Edge); ok {
	//		err = e.validate()
	//		if err != nil {
	//			return err
	//		}
	//		m[operation] = e
	//	} else {
	//		return fmt.Errorf("unknown type of object")
	//	}
	//
	//	err = g.enc.Encode(m)
	//	if err != nil {
	//		return err
	//	}
	//	delete(m, operation)
	//}

}
