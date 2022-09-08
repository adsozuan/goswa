package goswa

import (
	"fmt"
	"net/http"
)

func (s *Server) routes() {
	s.Handle("/", s.handleOla())
}

func (s *Server) handleOla() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s")
	}

}
