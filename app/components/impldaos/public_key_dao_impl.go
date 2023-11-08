package impldaos

import (
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/bitwormhole/passport-server/app/data/publickeys"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// PublicKeyDaoImpl ...
type PublicKeyDaoImpl struct {
	//starter:component
	_as func(publickeys.DAO) //starter:as("#")

	Agent       passportdb.LocalAgent //starter:inject("#")
	UUIDService random.UUIDService    //starter:inject("#")
}

func (inst *PublicKeyDaoImpl) _impl() publickeys.DAO {
	return inst
}

func (inst *PublicKeyDaoImpl) model() *publickeys.Entity {
	return &publickeys.Entity{}
}

func (inst *PublicKeyDaoImpl) modelList() []*publickeys.Entity {
	return make([]*publickeys.Entity, 0)
}

func (inst *PublicKeyDaoImpl) makeResult(ent *publickeys.Entity, res *gorm.DB) (*publickeys.Entity, error) {
	if ent == nil {
		return nil, fmt.Errorf("publickeys.Entity is nil")
	}
	err := res.Error
	if err != nil {
		return nil, err
	}
	return ent, nil
}

// Find ...
func (inst *PublicKeyDaoImpl) Find(db *gorm.DB, id publickeys.ID) (*publickeys.Entity, error) {
	db = inst.Agent.DB(db)
	o1 := inst.model()
	res := db.First(o1, id)
	return inst.makeResult(o1, res)
}

// FindBySum ...
func (inst *PublicKeyDaoImpl) FindBySum(db *gorm.DB, sum lang.Hex) (*publickeys.Entity, error) {
	db = inst.Agent.DB(db)
	o1 := inst.model()
	res := db.Where("sha256 = ?", sum).First(o1)
	return inst.makeResult(o1, res)
}

// ExistSum ...
func (inst *PublicKeyDaoImpl) ExistSum(db *gorm.DB, sum lang.Hex) (bool, error) {
	db = inst.Agent.DB(db)
	list := inst.modelList()
	res := db.Where("sha256 = ?", sum).Find(&list)
	err := res.Error
	if err != nil {
		return false, err
	}
	count := res.RowsAffected
	return count > 0, nil
}

// List ...
func (inst *PublicKeyDaoImpl) List(db *gorm.DB, q *publickeys.Query) ([]*publickeys.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *PublicKeyDaoImpl) Insert(db *gorm.DB, o1 *publickeys.Entity) (*publickeys.Entity, error) {

	uuid := inst.UUIDService.Build().Class("publickeys.Entity").Generate()

	o1.UUID = uuid
	o1.ID = 0

	db = inst.Agent.DB(db)
	res := db.Create(o1)
	return inst.makeResult(o1, res)
}

// Update ...
func (inst *PublicKeyDaoImpl) Update(db *gorm.DB, id publickeys.ID, callback func(*publickeys.Entity) error) (*publickeys.Entity, error) {

	db = inst.Agent.DB(db)
	o1 := inst.model()
	err := db.First(o1, id).Error
	if err != nil {
		return nil, err
	}

	err = callback(o1)
	if err != nil {
		return nil, err
	}

	res := db.Save(o1)
	return inst.makeResult(o1, res)
}

// Remove ...
func (inst *PublicKeyDaoImpl) Remove(db *gorm.DB, id publickeys.ID) error {
	db = inst.Agent.DB(db)
	o1 := inst.model()
	res := db.Delete(o1, id)
	return res.Error
}
