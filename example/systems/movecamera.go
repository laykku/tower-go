package systems

import (
	"github.com/laykku/example/ecs"
	"github.com/laykku/tower/engine"
)

const MovementSpeed float32 = 2.5

func MoveCameraSystem(world *ecs.World, deltaTime float32) {
	towerWorld := ecs.Resolve[engine.World](world)

	for _, e := range world.Query().MoveInput().Fetch() {
		moveInput := world.GetMoveInput(e)
		for _, e1 := range towerWorld.Query().CameraComponent().TransformComponent().Fetch() {
			transform := towerWorld.GetTransformComponent(e1)
			velocity := MovementSpeed * deltaTime

			if moveInput.Forward == 1 {
				transform.Position = transform.Position.Add(transform.Front.Mul(velocity))
			} else if moveInput.Forward == -1 {
				transform.Position = transform.Position.Add(transform.Front.Mul(-velocity))
			}

			if moveInput.Right == 1 {
				transform.Position = transform.Position.Add(transform.Right.Mul(velocity))
			} else if moveInput.Right == -1 {
				transform.Position = transform.Position.Add(transform.Right.Mul(-velocity))
			}
		}
	}
}
