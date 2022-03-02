package controller

import (
	"bookstore/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type BookController struct {
	s store.Store
}

func NewBookController(s store.Store) *BookController {
	return &BookController{s: s}
}

func (bs *BookController) CreateBookHandler(w http.ResponseWriter, request *http.Request)  {
	dec := json.NewDecoder(request.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bs.s.Create(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs *BookController) GetBookHandler(w http.ResponseWriter, request *http.Request)  {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}

	book, err := bs.s.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response(w, book)
}

func response(w http.ResponseWriter, v interface{})  {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}