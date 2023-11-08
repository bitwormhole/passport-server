package main

import (
	"os"

	passportserver "github.com/bitwormhole/passport-server"
	"github.com/starter-go/starter"
)

func main() {
	a := os.Args
	m := passportserver.Module()
	i := starter.Init(a)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
