package implservices

import (
	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/starter-go/libgorm"
	"gorm.io/gorm"
)

// LocalAgentImpl ...
type LocalAgentImpl struct {

	//starter:component
	_as func(passportdb.LocalAgent) //starter:as("#")

	DataSourceMan   libgorm.DataSourceManager //starter:inject("#")
	DataSourceAlias string                    //starter:inject("${datagroup.passport.datasource}")

	cache libgorm.DataSourceAgent
}

func (inst *LocalAgentImpl) _impl() {
	inst._as(inst)
}

// DB ...
func (inst *LocalAgentImpl) DB(db *gorm.DB) *gorm.DB {
	c := &inst.cache
	if !c.Ready() {
		c.Init(inst.DataSourceMan, inst.DataSourceAlias)
	}
	return c.DB(db)
}
