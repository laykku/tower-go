//go:generate go run config.go
//go:build generate

package main

import (
	"github.com/laykku/genecs/gen"
	"github.com/laykku/tower/engine"
)

func main() {
	gen.RegisterComponent[engine.TransformComponent]()
	gen.RegisterComponent[engine.CameraComponent]()
	gen.RegisterComponent[engine.MeshComponent]()
	gen.RegisterComponent[engine.BatchList]()
	gen.Generate("engine")
}
