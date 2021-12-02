package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"errors"
	"fmt"
	"sync"
)

func init() {
	memStore := MemStore{books: make(map[string]*mystore.Book)}
	factory.Register("mem", &memStore)
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}


func (m *MemStore) Create(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()
	if book == nil {
		return errors.New("request is nil")
	}

	if _, ok := m.books[book.Id]; ok {
		return fmt.Errorf("The book already exists")
	}

	m.books[book.Id] = book
	return nil
}

func (m *MemStore) Update(book *mystore.Book) error {
	panic("implement me")
}

func (m *MemStore) Get(s string) (mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()
	book, ok := m.books[s]
	if !ok {
		return mystore.Book{}, fmt.Errorf("The book was not found, id:%s", s)
	}
	return *book, nil
}

func (m *MemStore) GetAll() ([]mystore.Book, error) {
	panic("implement me")
}

func (m *MemStore) Delete(s string) error {
	panic("implement me")
}
