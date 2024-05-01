package main

import (
	"fmt"

	"github.com/laykku/example/ecs"
	"github.com/laykku/example/systems"
	"github.com/laykku/tower/engine"
	"github.com/laykku/tower/webgl"
)

func main() {
	window := webgl.CreateWindow()
	renderer := webgl.CreateOpenGLRenderer(window.GetGlContext())

	engine.CreateEngine(window, renderer, func(towerWorld *engine.World) {

		eng := engine.Resolve[engine.Engine](towerWorld)

		exampleWorld := ecs.CreateWorld()
		exampleWorld.Inject(towerWorld)
		exampleWorld.Inject(eng)

		exampleWorld.RegisterInit(systems.InitScene)

		towerWorld.Register(func(world *engine.World, deltaTime float32) { exampleWorld.Tick(deltaTime) })

		//registerStat(towerWorld, exampleWorld)
	})
}

func registerStat(towerWorld *engine.World, exampleWorld *ecs.World) {
	towerWorld.Register(func(world *engine.World, deltaTime float32) {
		e, c := world.Stat()
		fmt.Println("tower:", e, c)
	})

	exampleWorld.Register(func(world *ecs.World, deltaTime float32) {
		e, c := world.Stat()
		fmt.Println("atc:", e, c)
	})
}
