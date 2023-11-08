package impldaos

import (
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/books"
	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// BookDaoImpl ...
type BookDaoImpl struct {
	//starter:component
	_as func(books.DAO) //starter:as("#")

	Agent       passportdb.LocalAgent //starter:inject("#")
	UUIDService random.UUIDService    //starter:inject("#")
}

func (inst *BookDaoImpl) _impl() books.DAO {
	return inst
}

// Find ...
func (inst *BookDaoImpl) Find(db *gorm.DB, id books.ID) (*books.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *BookDaoImpl) List(db *gorm.DB, q *books.Query) ([]*books.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *BookDaoImpl) Insert(db *gorm.DB, o *books.Entity) (*books.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *BookDaoImpl) Update(db *gorm.DB, id books.ID, callback func(*books.Entity) error) (*books.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *BookDaoImpl) Remove(db *gorm.DB, id books.ID) error {
	return fmt.Errorf("no impl")
}
