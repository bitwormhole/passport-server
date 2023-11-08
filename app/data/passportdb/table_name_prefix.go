package passportdb

// PrefixSetter ...
type PrefixSetter interface {
	Set(prefix string)
}

////////////////////////////////////////////////////////////////////////////////

// Prefix 返回数据实体分组（passport）的表名前缀
func Prefix() string {
	return thePrefix
}

// SetPrefix ...
func SetPrefix() PrefixSetter {
	return &setter{}
}

////////////////////////////////////////////////////////////////////////////////

var thePrefix string

type setter struct {
}

func (inst *setter) Set(prefix string) {
	thePrefix = prefix
}
