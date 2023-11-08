package implservices

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/example"
)

// ExampleServiceImpl ...
type ExampleServiceImpl struct {

	//starter:component
	_as func(example.Service) //starter:as("#")

	Convertor example.Convertor //starter:inject("#")
	Dao       example.DAO       //starter:inject("#")
}

func (inst *ExampleServiceImpl) _impl() example.Service {
	return inst
}

// Find ...
func (inst *ExampleServiceImpl) Find(c context.Context, id example.ID) (*example.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ExampleServiceImpl) List(c context.Context, q *example.Query) ([]*example.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *ExampleServiceImpl) Insert(c context.Context, o *example.DTO) (*example.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ExampleServiceImpl) Update(c context.Context, id example.ID, src *example.DTO) (*example.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ExampleServiceImpl) Remove(c context.Context, id example.ID) error {
	return fmt.Errorf("no impl")
}
