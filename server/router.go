package server

import (
	"bookstore/internal/controller"
	"bookstore/server/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	*controller.BookController
}

func (r *Router) Register() http.Handler {
	r.Use(middleware.Logging, middleware.Validating)
	r.HandleFunc("/book", r.CreateBookHandler).Methods("POST")
	r.HandleFunc("/book/{id}", r.GetBookHandler).Methods("GET")
	return r
}
