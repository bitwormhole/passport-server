package implconvertors

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/example"
)

// ExampleConvertorImpl ...
type ExampleConvertorImpl struct {
	//starter:component
	_as func(example.Convertor) //starter:as("#")
}

func (inst *ExampleConvertorImpl) _impl() example.Convertor {
	return inst
}

// ConvertListE2D ...
func (inst *ExampleConvertorImpl) ConvertListE2D(c context.Context, src []*example.Entity) ([]*example.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertE2D ...
func (inst *ExampleConvertorImpl) ConvertE2D(c context.Context, src *example.Entity) (*example.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *ExampleConvertorImpl) ConvertD2E(c context.Context, src *example.DTO) (*example.Entity, error) {
	return nil, fmt.Errorf("no impl")
}
