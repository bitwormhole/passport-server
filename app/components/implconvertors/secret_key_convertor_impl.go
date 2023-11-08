package implconvertors

import (
	"context"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/secretkeys"
)

// SecretKeyConvertorImpl ...
type SecretKeyConvertorImpl struct {
	//starter:component
	_as func(secretkeys.Convertor) //starter:as("#")
}

func (inst *SecretKeyConvertorImpl) _impl() secretkeys.Convertor {
	return inst
}

// ConvertListE2D ...
func (inst *SecretKeyConvertorImpl) ConvertListE2D(c context.Context, src []*secretkeys.Entity) ([]*secretkeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertE2D ...
func (inst *SecretKeyConvertorImpl) ConvertE2D(c context.Context, src *secretkeys.Entity) (*secretkeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// ConvertD2E ...
func (inst *SecretKeyConvertorImpl) ConvertD2E(c context.Context, src *secretkeys.DTO) (*secretkeys.Entity, error) {
	return nil, fmt.Errorf("no impl")
}
