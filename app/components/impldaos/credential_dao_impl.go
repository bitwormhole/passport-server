package impldaos

import (
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/credentials"
	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// CredentialDaoImpl ...
type CredentialDaoImpl struct {
	//starter:component
	_as func(credentials.DAO) //starter:as("#")

	Agent       passportdb.LocalAgent //starter:inject("#")
	UUIDService random.UUIDService    //starter:inject("#")
}

func (inst *CredentialDaoImpl) _impl() credentials.DAO {
	return inst
}

// Find ...
func (inst *CredentialDaoImpl) Find(db *gorm.DB, id credentials.ID) (*credentials.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *CredentialDaoImpl) List(db *gorm.DB, q *credentials.Query) ([]*credentials.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *CredentialDaoImpl) Insert(db *gorm.DB, o *credentials.Entity) (*credentials.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *CredentialDaoImpl) Update(db *gorm.DB, id credentials.ID, callback func(*credentials.Entity) error) (*credentials.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *CredentialDaoImpl) Remove(db *gorm.DB, id credentials.ID) error {
	return fmt.Errorf("no impl")
}
