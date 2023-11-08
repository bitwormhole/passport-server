package implservices

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/books"
)

// BookServiceImpl ...
type BookServiceImpl struct {

	//starter:component
	_as func(books.Service) //starter:as("#")

	Convertor books.Convertor //starter:inject("#")
	Dao       books.DAO       //starter:inject("#")
}

func (inst *BookServiceImpl) _impl() books.Service {
	return inst
}

// Find ...
func (inst *BookServiceImpl) Find(c context.Context, id books.ID) (*books.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *BookServiceImpl) List(c context.Context, q *books.Query) ([]*books.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *BookServiceImpl) Insert(c context.Context, o *books.DTO) (*books.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *BookServiceImpl) Update(c context.Context, id books.ID, src *books.DTO) (*books.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *BookServiceImpl) Remove(c context.Context, id books.ID) error {
	return fmt.Errorf("no impl")
}
