package implconvertors

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/books"
)

// BookConvertorImpl ...
type BookConvertorImpl struct {
	//starter:component
	_as func(books.Convertor) //starter:as("#")
}

func (inst *BookConvertorImpl) _impl() books.Convertor {
	return inst
}

// ConvertListE2D ...
func (inst *BookConvertorImpl) ConvertListE2D(c context.Context, src []*books.Entity) ([]*books.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertE2D ...
func (inst *BookConvertorImpl) ConvertE2D(c context.Context, src *books.Entity) (*books.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *BookConvertorImpl) ConvertD2E(c context.Context, src *books.DTO) (*books.Entity, error) {
	return nil, fmt.Errorf("no impl")
}
