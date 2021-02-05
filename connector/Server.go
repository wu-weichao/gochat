package connector

type Server struct {
	C map[interface{}]*Connect
}

func NewServer() *Server {
	return &Server{}
}
