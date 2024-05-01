//go:generate go run config.go
//go:build generate

package main

import (
	"github.com/laykku/example/ecs"
	"github.com/laykku/genecs/gen"
)

func main() {
	gen.RegisterComponent[ecs.MouseMove]()
	gen.RegisterComponent[ecs.MoveInput]()
	gen.Generate("ecs")
}
