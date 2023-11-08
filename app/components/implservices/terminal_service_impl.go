package implservices

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/terminals"
)

// TerminalServiceImpl ...
type TerminalServiceImpl struct {

	//starter:component
	_as func(terminals.Service) //starter:as("#")

	Convertor terminals.Convertor //starter:inject("#")
	Dao       terminals.DAO       //starter:inject("#")
}

func (inst *TerminalServiceImpl) _impl() terminals.Service {
	return inst
}

// Find ...
func (inst *TerminalServiceImpl) Find(c context.Context, id terminals.ID) (*terminals.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *TerminalServiceImpl) List(c context.Context, q *terminals.Query) ([]*terminals.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *TerminalServiceImpl) Insert(c context.Context, o *terminals.DTO) (*terminals.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *TerminalServiceImpl) Update(c context.Context, id terminals.ID, src *terminals.DTO) (*terminals.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *TerminalServiceImpl) Remove(c context.Context, id terminals.ID) error {
	return fmt.Errorf("no impl")
}
