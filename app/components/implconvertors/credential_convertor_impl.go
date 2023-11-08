package implconvertors

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/credentials"
)

// CredentialConvertorImpl ...
type CredentialConvertorImpl struct {
	//starter:component
	_as func(credentials.Convertor) //starter:as("#")
}

func (inst *CredentialConvertorImpl) _impl() credentials.Convertor {
	return inst
}

// ConvertListE2D ...
func (inst *CredentialConvertorImpl) ConvertListE2D(c context.Context, src []*credentials.Entity) ([]*credentials.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertE2D ...
func (inst *CredentialConvertorImpl) ConvertE2D(c context.Context, src *credentials.Entity) (*credentials.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *CredentialConvertorImpl) ConvertD2E(c context.Context, src *credentials.DTO) (*credentials.Entity, error) {
	return nil, fmt.Errorf("no impl")
}
