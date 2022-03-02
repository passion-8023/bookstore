package server

import (
	"bookstore/pkg/transport"
	"bookstore/pkg/transport/http"
)

func NewHttpServer(router *Router) []transport.ServerInterface {
	httpServer := http.NewServer(
		http.Address(":5000"),
		http.Handler(router.Register()),
	)
	return []transport.ServerInterface{httpServer}
}
