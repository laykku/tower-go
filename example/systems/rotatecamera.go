package systems

import (
	"github.com/laykku/example/ecs"
	"github.com/laykku/tower/engine"
)

func RotateCameraSystem(world *ecs.World, deltaTime float32) {

	towerWorld := ecs.Resolve[engine.World](world)

	for _, e := range world.Query().MouseMove().Fetch() {
		mouseMove := world.GetMouseMove(e)
		deltaX := float32(mouseMove.XRel) * 0.1
		deltaY := float32(mouseMove.YRel) * 0.1

		for _, e := range towerWorld.Query().CameraComponent().Fetch() {
			camera := towerWorld.GetCameraComponent(e)

			camera.Yaw += deltaX
			camera.Pitch -= deltaY

			if camera.Pitch > 89.0 {
				camera.Pitch = 89.0
			}
			if camera.Pitch < -89.0 {
				camera.Pitch = -89.0
			}
		}
	}
}
