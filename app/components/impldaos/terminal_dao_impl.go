package impldaos

import (
	"fmt"

	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/bitwormhole/passport-server/app/data/terminals"
	"github.com/starter-go/security/random"
	"gorm.io/gorm"
)

// TerminalDaoImpl ...
type TerminalDaoImpl struct {
	//starter:component
	_as func(terminals.DAO) //starter:as("#")

	Agent       passportdb.LocalAgent //starter:inject("#")
	UUIDService random.UUIDService    //starter:inject("#")
}

func (inst *TerminalDaoImpl) _impl() terminals.DAO {
	return inst
}

// Find ...
func (inst *TerminalDaoImpl) Find(db *gorm.DB, id terminals.ID) (*terminals.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// List ...
func (inst *TerminalDaoImpl) List(db *gorm.DB, q *terminals.Query) ([]*terminals.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Insert ...
func (inst *TerminalDaoImpl) Insert(db *gorm.DB, o *terminals.Entity) (*terminals.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Update ...
func (inst *TerminalDaoImpl) Update(db *gorm.DB, id terminals.ID, callback func(*terminals.Entity) error) (*terminals.Entity, error) {
	return nil, fmt.Errorf("no impl")
}

// Remove ...
func (inst *TerminalDaoImpl) Remove(db *gorm.DB, id terminals.ID) error {
	return fmt.Errorf("no impl")
}
