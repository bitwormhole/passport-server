package gen4server

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForPassportServer 导出组件配置
func ExportComForPassportServer(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
