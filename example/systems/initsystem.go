package systems

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/laykku/example/ecs"
	"github.com/laykku/tower/assetman"
	"github.com/laykku/tower/engine"
)

func InitScene(world *ecs.World) {

	towerWorld := ecs.Resolve[engine.World](world)

	eng := ecs.Resolve[engine.Engine](world)

	renderer := eng.GetRenderer()

	vs := assetman.LoadShader("vertex_shader")
	fs := assetman.LoadShader("fragment_shader")

	shader := renderer.CreateProgram(vs, fs)

	meshData := assetman.LoadMesh("cube_mesh")

	camera := towerWorld.CreateEntity()

	engine.TransformComponent{
		Entity:   camera,
		Position: mgl32.Vec3{0, 2, 2},
	}.Add(towerWorld)

	engine.CameraComponent{
		Entity: camera,
		Yaw:    -90.0,
		Pitch:  -45,
		Fov:    60.0,
	}.Add(towerWorld)

	data, w, h := assetman.LoadTexture("checker_texture")
	texture := renderer.CreateTexture(data, w, h)

	// meshes

	mat := &engine.Material{
		Program:  shader,
		Texture0: texture,
	}

	mesh := towerWorld.CreateEntity()

	engine.TransformComponent{
		Entity:   mesh,
		Position: mgl32.Vec3{0, 0, 0},
		Rotation: mgl32.QuatRotate(45, mgl32.Vec3{0, 1, 0}),
		Scale:    mgl32.Vec3{1, 1, 1},
	}.Add(towerWorld)

	engine.MeshComponent{
		Entity:   mesh,
		Data:     meshData,
		Material: mat,
		Handle:   meshData.CreateHandle(renderer),
	}.Add(towerWorld)

	world.Register(func(world *ecs.World, deltaTime float32) {
		for _, e := range towerWorld.Query().TransformComponent().MeshComponent().Fetch() {
			t := towerWorld.GetTransformComponent(e)
			t.Rotation = t.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(180.0*deltaTime), mgl32.Vec3{0, 1, 0}))
		}
	})

	world.RegisterInit(InitInputSystem)
	world.Register(ProcessMouseInputSystem)
	world.Register(ProcessKeyboardInputSystem)
	world.Register(RotateCameraSystem)
	world.Register(MoveCameraSystem)
}
