package implconvertors

import (
	"context"

	"github.com/bitwormhole/passport-server/app/data/publickeys"
	"github.com/starter-go/security-gorm/rbacdb"
)

// PublicKeyConvertorImpl ...
type PublicKeyConvertorImpl struct {
	//starter:component
	_as func(publickeys.Convertor) //starter:as("#")
}

func (inst *PublicKeyConvertorImpl) _impl() publickeys.Convertor {
	return inst
}

// ConvertListE2D ...
func (inst *PublicKeyConvertorImpl) ConvertListE2D(c context.Context, src []*publickeys.Entity) ([]*publickeys.DTO, error) {
	dst := make([]*publickeys.DTO, 0)
	for _, o1 := range src {
		o2, err := inst.ConvertE2D(c, o1)
		if err != nil {
			return nil, err
		}
		dst = append(dst, o2)
	}
	return dst, nil
}

// ConvertE2D ...
func (inst *PublicKeyConvertorImpl) ConvertE2D(c context.Context, src *publickeys.Entity) (*publickeys.DTO, error) {
	dst := &publickeys.DTO{}
	dst.ID = src.ID
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.Algorithm = src.Algorithm
	dst.Format = src.Format
	dst.Data = src.Data
	dst.SHA256 = src.SHA256
	return dst, nil
}

// ConvertD2E ...
func (inst *PublicKeyConvertorImpl) ConvertD2E(c context.Context, src *publickeys.DTO) (*publickeys.Entity, error) {
	dst := &publickeys.Entity{}
	dst.ID = src.ID
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.Algorithm = src.Algorithm
	dst.Format = src.Format
	dst.Data = src.Data
	dst.SHA256 = src.SHA256
	return dst, nil
}
