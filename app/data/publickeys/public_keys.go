package publickeys

import (
	"context"

	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
	"github.com/starter-go/security/rbac"
	"gorm.io/gorm"
)

// ID ...
type ID int64

// Entity ...
type Entity struct {
	ID ID

	rbacdb.BaseEntity

	Algorithm string
	Format    string
	Data      lang.Base64
	SHA256    lang.Hex `gorm:"unique"` // aka: fingerprint
}

// DTO ...
type DTO struct {
	ID ID `json:"id"`

	rbac.BaseDTO

	Algorithm string      `json:"algorithm"`
	Format    string      `json:"format"`
	SHA256    lang.Hex    `json:"fingerprint"`
	Data      lang.Base64 `json:"data"`
}

// VO ...
type VO struct {
	rbac.BaseVO

	Items []*DTO `json:"publickeys"`
}

// Query ...
type Query struct {
	Conditions rbac.Conditions
	Pagination rbac.Pagination
	All        bool
}

////////////////////////////////////////////////////////////////////////////////

// Convertor ...
type Convertor interface {
	ConvertListE2D(c context.Context, src []*Entity) ([]*DTO, error)
	ConvertE2D(c context.Context, src *Entity) (*DTO, error)
	ConvertD2E(c context.Context, src *DTO) (*Entity, error)
}

// Service ...
type Service interface {
	Find(c context.Context, id ID) (*DTO, error)
	List(c context.Context, q *Query) ([]*DTO, error)
	FindBySum(c context.Context, sum lang.Hex) (*DTO, error)
	ExistSum(c context.Context, sum lang.Hex) (bool, error)

	Insert(c context.Context, o *DTO) (*DTO, error)
	Update(c context.Context, id ID, src *DTO) (*DTO, error)
	Remove(c context.Context, id ID) error
}

// DAO ...
type DAO interface {
	Find(db *gorm.DB, id ID) (*Entity, error)
	List(db *gorm.DB, q *Query) ([]*Entity, error)
	FindBySum(db *gorm.DB, sum lang.Hex) (*Entity, error)
	ExistSum(db *gorm.DB, sum lang.Hex) (bool, error)

	Insert(db *gorm.DB, o *Entity) (*Entity, error)
	Update(db *gorm.DB, id ID, callback func(*Entity) error) (*Entity, error)
	Remove(db *gorm.DB, id ID) error
}

////////////////////////////////////////////////////////////////////////////////

// TableName ...
func (Entity) TableName() string {
	prefix := passportdb.Prefix()
	return prefix + "public_keys"
}

////////////////////////////////////////////////////////////////////////////////
