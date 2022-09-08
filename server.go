package goswa

import "net/http"

type Server struct {
	*http.ServeMux
}

func NewServer() *Server {
	s := &Server{}
	s.ServeMux = http.NewServeMux()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}
