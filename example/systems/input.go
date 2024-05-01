package systems

import (
	"github.com/laykku/example/ecs"
	"github.com/laykku/tower/engine"
)

func InitInputSystem(world *ecs.World) {
	mouseMoveEntity := world.CreateEntity()
	ecs.MouseMove{
		Entity: mouseMoveEntity,
	}.Add(world)

	moveInputEntity := world.CreateEntity()
	ecs.MoveInput{
		Entity: moveInputEntity,
	}.Add(world)
}

func ProcessMouseInputSystem(world *ecs.World, deltaTime float32) {
	eng := ecs.Resolve[engine.Engine](world)
	window := eng.GetWindow()

	x, y := window.GetMouseState()

	for _, e := range world.Query().MouseMove().Fetch() {
		mouseMove := world.GetMouseMove(e)
		mouseMove.XRel = x - mouseMove.XPrev
		mouseMove.YRel = y - mouseMove.YPrev
	}
}

func ProcessKeyboardInputSystem(world *ecs.World, deltaTime float32) {
	eng := ecs.Resolve[engine.Engine](world)
	window := eng.GetWindow()

	for _, e := range world.Query().MoveInput().Fetch() {
		moveInput := world.GetMoveInput(e)
		moveInput.Forward = 0
		moveInput.Right = 0
		if window.IsKeyPressed("w") {
			moveInput.Forward = 1
		}
		if window.IsKeyPressed("s") {
			moveInput.Forward = -1
		}
		if window.IsKeyPressed("a") {
			moveInput.Right = -1
		}
		if window.IsKeyPressed("d") {
			moveInput.Right = 1
		}
	}
}
