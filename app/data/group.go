package data

import (
	"github.com/bitwormhole/passport-server/app/data/passportdb"
	"github.com/starter-go/libgorm"
)

// ThisGroup ...
type ThisGroup struct {

	//starter:component
	_as func(libgorm.GroupRegistry) //starter:as(".")

	Enabled bool   //starter:inject("${datagroup.passport.enabled}")
	Alias   string //starter:inject("${datagroup.passport.alias}")
	URI     string //starter:inject("${datagroup.passport.uri}")
	Prefix  string //starter:inject("${datagroup.passport.table-name-prefix}")
	Source  string //starter:inject("${datagroup.passport.datasource}")
}

func (inst *ThisGroup) _impl() (libgorm.GroupRegistry, libgorm.Group) {
	return inst, inst
}

// Groups ...
func (inst *ThisGroup) Groups() []*libgorm.GroupRegistration {

	passportdb.SetPrefix().Set(inst.Prefix)

	r1 := &libgorm.GroupRegistration{
		Enabled: inst.Enabled,
		Alias:   inst.Alias,
		URI:     inst.URI,
		Prefix:  inst.Prefix,
		Source:  inst.Source,
		Group:   inst,
	}
	return []*libgorm.GroupRegistration{r1}
}

// Prototypes ...
func (inst *ThisGroup) Prototypes() []any {
	return listTables()
}
