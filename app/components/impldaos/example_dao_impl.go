package impldaos

import (
	"fmt"

	examples "github.com/bitwormhole/passport-server/app/data/example"
	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// ExampleDaoImpl ...
type ExampleDaoImpl struct {
	//starter:component
	_as func(examples.DAO) //starter:as("#")

	Agent       passportdb.LocalAgent //starter:inject("#")
	UUIDService random.UUIDService    //starter:inject("#")
}

func (inst *ExampleDaoImpl) _impl() examples.DAO {
	return inst
}

// Find ...
func (inst *ExampleDaoImpl) Find(db *gorm.DB, id examples.ID) (*examples.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *ExampleDaoImpl) List(db *gorm.DB, q *examples.Query) ([]*examples.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *ExampleDaoImpl) Insert(db *gorm.DB, o *examples.Entity) (*examples.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *ExampleDaoImpl) Update(db *gorm.DB, id examples.ID, callback func(*examples.Entity) error) (*examples.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *ExampleDaoImpl) Remove(db *gorm.DB, id examples.ID) error {
	return fmt.Errorf("no impl")
}
