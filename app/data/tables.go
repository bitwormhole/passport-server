package data

import (
	"github.com/bitwormhole/passport-server/app/data/books"
	"github.com/bitwormhole/passport-server/app/data/credentials"
	"github.com/bitwormhole/passport-server/app/data/publickeys"
	"github.com/bitwormhole/passport-server/app/data/secretkeys"
)

func listTables() []any {
	list := make([]any, 0)

	list = append(list, &books.Entity{})
	list = append(list, &credentials.Entity{})
	list = append(list, &publickeys.Entity{})
	list = append(list, &secretkeys.Entity{})

	return list
}
