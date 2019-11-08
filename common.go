package gephi_http_client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type gephiClient struct {
	w   io.WriteCloser
	enc json.Encoder
	wg  sync.WaitGroup
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

	gc := gephiClient{
		w:   w,
		enc: *json.NewEncoder(w),
	}
	gc.wg.Add(1)

	//r, w := io.Pipe()
	go func() {
		defer gc.wg.Done()
		defer r.Close()

		buf := bufio.NewReader(r)
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("read line error:", err)
				continue
			}

			resp, err := client.Post(url, "application/json", bytes.NewReader(line))
			if err != nil {
				log.Println("Post error:", err)
			}
			aaa, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("******* resp.Body %s, %v\n", aaa, err)
			err = resp.Body.Close()
			if err != nil {
				log.Println("Post Body.Close error:", err)
			}
		}
	}()

	return &gc, nil
}

func (g *gephiClient) Close() error {
	return g.w.Close()
}

func (g *gephiClient) Wait() {
	g.wg.Wait()
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

			b, err := json.Marshal(m)

			//err = g.enc.Encode(m)
			if err != nil {
				log.Println("lllll ", err, m)
				return err
			}
			aaa := append([]byte("\\r\\n"), '\n')
			_, err = g.w.Write(append(b, aaa...))
			if err != nil {
				log.Println("lllll node", err, string(aaa))
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

			b, err := json.Marshal(m)
			//err = g.enc.Encode(m)
			if err != nil {
				log.Println("lllll ", err, m)
				return err
			}
			aaa := append([]byte("\\r\\n"), '\n')
			_, err = g.w.Write(append(b, aaa...))
			if err != nil {
				log.Println("lllll edge", err, string(aaa))
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
