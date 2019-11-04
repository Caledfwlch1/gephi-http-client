package gephi_http_client

type GephiClient interface {
	NodeAdd(node ...Node) error
	NodeChange(node ...Node) error
	NodeDelete(node ...Node) error

	EdgeAdd(edge ...Edge) error
	EdgeChange(edge ...Edge) error
	EdgeDelete(edge ...Edge) error

	Close() error
}

type Node map[string]map[string]string

//struct {
//	Id    string
//	Lable string
//	X     float64
//	Y     float64
//	Size  int
//	Prop  map[string]string
//}

type Edge map[string]map[string]string

//struct {
//	Id       string
//	Source   string
//	Target   string
//	Directed bool
//	Weight   int
//}
