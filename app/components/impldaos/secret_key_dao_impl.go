package impldaos

import (
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/bitwormhole/passport-server/app/data/secretkeys"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// SecretKeyDaoImpl ...
type SecretKeyDaoImpl struct {
	//starter:component
	_as func(secretkeys.DAO) //starter:as("#")

	Agent       passportdb.LocalAgent //starter:inject("#")
	UUIDService random.UUIDService    //starter:inject("#")
}

func (inst *SecretKeyDaoImpl) _impl() secretkeys.DAO {
	return inst
}

// Find ...
func (inst *SecretKeyDaoImpl) Find(db *gorm.DB, id secretkeys.ID) (*secretkeys.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *SecretKeyDaoImpl) List(db *gorm.DB, q *secretkeys.Query) ([]*secretkeys.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *SecretKeyDaoImpl) Insert(db *gorm.DB, o *secretkeys.Entity) (*secretkeys.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *SecretKeyDaoImpl) Update(db *gorm.DB, id secretkeys.ID, callback func(*secretkeys.Entity) error) (*secretkeys.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *SecretKeyDaoImpl) Remove(db *gorm.DB, id secretkeys.ID) error {
	return fmt.Errorf("no impl")
}
