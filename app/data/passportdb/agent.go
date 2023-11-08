package passportdb

import "github.com/starter-go/libgorm"

// LocalAgent 数据库本地代理
type LocalAgent interface {
	libgorm.Agent
}
