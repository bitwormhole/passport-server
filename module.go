package passportserver

import (
	"embed"

	"github.com/bitwormhole/passport-server/gen/gen4server"
	"github.com/starter-go/application"
	"github.com/starter-go/security-gin-gorm/modules/securitygingorm"
)

const (
	theModuleName     = "github.com/bitwormhole/passport-server"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块 "github.com/bitwormhole/passport-server"
func Module() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4server.ExportComForPassportServer)

	mb.Depend(securitygingorm.Module())
	return mb.Create()
}
