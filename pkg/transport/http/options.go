package http

import (
	"net/http"
	"time"
)

type serverOptions func(*Server)

func Address(addr string) serverOptions {
	return func(s *Server) {
		s.address = addr
	}
}

func Timeout(timeout time.Duration) serverOptions {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func Handler(handler http.Handler) serverOptions {
	return func(s *Server) {
		s.Serve.Handler = handler
	}
}

