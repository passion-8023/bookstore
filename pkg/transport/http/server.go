package http

import (
	"bytes"
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Serve   *http.Server
	lis     net.Listener
	err     error
	network string
	address string
	timeout time.Duration
}

func NewServer(opts ...serverOptions) *Server {
	srv := &Server{
		Serve:   new(http.Server),
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
	}

	for _, opt := range opts {
		opt(srv)
	}

	if srv.lis == nil {
		srv.lis, srv.err = net.Listen(srv.network, srv.address)
	}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return nil
	}

	s.Serve.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	err := s.Serve.Serve(s.lis)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Serve.Shutdown(ctx)
}

func (s *Server) Addr(b *bytes.Buffer) {
	if b.Len() > 0 {
		b.WriteString(",")
	}
	b.WriteString(s.lis.Addr().String())
}
