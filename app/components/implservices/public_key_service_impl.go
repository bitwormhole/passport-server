package implservices

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/publickeys"
	"github.com/starter-go/base/lang"
)

// PublicKeyServiceImpl ...
type PublicKeyServiceImpl struct {

	//starter:component
	_as func(publickeys.Service) //starter:as("#")

	Convertor publickeys.Convertor //starter:inject("#")
	Dao       publickeys.DAO       //starter:inject("#")
}

func (inst *PublicKeyServiceImpl) _impl() publickeys.Service {
	return inst
}

// Find ...
func (inst *PublicKeyServiceImpl) Find(c context.Context, id publickeys.ID) (*publickeys.DTO, error) {
	o1, err := inst.Dao.Find(nil, id)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertE2D(c, o1)
}

// FindBySum ...
func (inst *PublicKeyServiceImpl) FindBySum(c context.Context, sum lang.Hex) (*publickeys.DTO, error) {
	o1, err := inst.Dao.FindBySum(nil, sum)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertE2D(c, o1)
}

// ExistSum ...
func (inst *PublicKeyServiceImpl) ExistSum(c context.Context, sum lang.Hex) (bool, error) {
	return inst.Dao.ExistSum(nil, sum)
}

// List ...
func (inst *PublicKeyServiceImpl) List(c context.Context, q *publickeys.Query) ([]*publickeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *PublicKeyServiceImpl) Insert(c context.Context, o *publickeys.DTO) (*publickeys.DTO, error) {
	err := inst.checkFingerprint(o)
	if err != nil {
		return nil, err
	}
	o2, err := inst.Convertor.ConvertD2E(c, o)
	if err != nil {
		return nil, err
	}
	o3, err := inst.Dao.Insert(nil, o2)
	if err != nil {
		return nil, err
	}
	return inst.Convertor.ConvertE2D(c, o3)
}

func (inst *PublicKeyServiceImpl) checkFingerprint(o *publickeys.DTO) error {

	data := o.Data.Bytes()
	sum := sha256.Sum256(data)
	have := sum[:]

	if o.SHA256 == "" {
		o.SHA256 = lang.HexFromBytes(have)
		return nil
	}

	want := o.SHA256.Bytes()
	if !bytes.Equal(want, have) {
		s1 := lang.HexFromBytes(want).String()
		s2 := lang.HexFromBytes(have).String()
		return fmt.Errorf("bad public-key fingerprint, want:%s, have:%s", s1, s2)
	}
	return nil
}

// Update ...
func (inst *PublicKeyServiceImpl) Update(c context.Context, id publickeys.ID, src *publickeys.DTO) (*publickeys.DTO, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *PublicKeyServiceImpl) Remove(c context.Context, id publickeys.ID) error {
	return fmt.Errorf("no impl")
}
