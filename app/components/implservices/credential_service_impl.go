package implservices

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/credentials"
)

// CredentialServiceImpl ...
type CredentialServiceImpl struct {

	//starter:component
	_as func(credentials.Service) //starter:as("#")

	Convertor credentials.Convertor //starter:inject("#")
	Dao       credentials.DAO       //starter:inject("#")
}

func (inst *CredentialServiceImpl) _impl() credentials.Service {
	return inst
}

// Find ...
func (inst *CredentialServiceImpl) Find(c context.Context, id credentials.ID) (*credentials.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *CredentialServiceImpl) List(c context.Context, q *credentials.Query) ([]*credentials.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *CredentialServiceImpl) Insert(c context.Context, o *credentials.DTO) (*credentials.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *CredentialServiceImpl) Update(c context.Context, id credentials.ID, src *credentials.DTO) (*credentials.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *CredentialServiceImpl) Remove(c context.Context, id credentials.ID) error {
	return fmt.Errorf("no impl")
}
