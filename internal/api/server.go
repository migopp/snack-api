package api

type Server struct {
	Address string
	Port    uint16
}

func (s *Server) Run() error {
	return nil
}
