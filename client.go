package gephi_http_client

import "net/http"

type gclient struct {
	Client *http.Client

	host      string
	workspace string

	Node
	Edge
}

func New(host, workspace string) GephiClient {
	defaultClient := http.DefaultClient

	c := &gclient{
		Client:    defaultClient,
		host:      host,
		workspace: workspace,
	}

	c.Node = newNode(c)
	c.Edge = newEdge(c)
	return c
}
