package implservices

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/secretkeys"
)

// SecretKeyServiceImpl ...
type SecretKeyServiceImpl struct {

	//starter:component
	_as func(secretkeys.Service) //starter:as("#")

	Convertor secretkeys.Convertor //starter:inject("#")
	Dao       secretkeys.DAO       //starter:inject("#")
}

func (inst *SecretKeyServiceImpl) _impl() secretkeys.Service {
	return inst
}

// Find ...
func (inst *SecretKeyServiceImpl) Find(c context.Context, id secretkeys.ID) (*secretkeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *SecretKeyServiceImpl) List(c context.Context, q *secretkeys.Query) ([]*secretkeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *SecretKeyServiceImpl) Insert(c context.Context, o *secretkeys.DTO) (*secretkeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *SecretKeyServiceImpl) Update(c context.Context, id secretkeys.ID, src *secretkeys.DTO) (*secretkeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *SecretKeyServiceImpl) Remove(c context.Context, id secretkeys.ID) error {
	return fmt.Errorf("no impl")
}
