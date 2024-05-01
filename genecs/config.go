//go:generate go run config.go
//go:build generate

package main

import (
	"github.com/laykku/genecs/ecs"
	"github.com/laykku/genecs/gen"
)

func main() {
	gen.RegisterComponent[ecs.Person]()
	gen.RegisterComponent[ecs.Greeter]()
	gen.Generate("main")
}
