package implconvertors

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/terminals"
)

// TerminalConvertorImpl ...
type TerminalConvertorImpl struct {
	//starter:component
	_as func(terminals.Convertor) //starter:as("#")
}

func (inst *TerminalConvertorImpl) _impl() terminals.Convertor {
	return inst
}

// ConvertListE2D ...
func (inst *TerminalConvertorImpl) ConvertListE2D(c context.Context, src []*terminals.Entity) ([]*terminals.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertE2D ...
func (inst *TerminalConvertorImpl) ConvertE2D(c context.Context, src *terminals.Entity) (*terminals.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *TerminalConvertorImpl) ConvertD2E(c context.Context, src *terminals.DTO) (*terminals.Entity, error) {
	return nil, fmt.Errorf("no impl")
}
