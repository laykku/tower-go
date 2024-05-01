package engine

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/laykku/tower/utils"
)

type Engine struct {
	window   Window
	renderer Renderer
	Width    int32
	Height   int32
	world    *World
}

func CreateEngine(window Window, renderer Renderer, onInit func(world *World)) {

	world := CreateWorld()
	world.RegisterInit(initSystem)
	world.Register(updateCameraSystem)
	world.Register(batchSystem)
	world.Register(renderSystem)

	eng := &Engine{
		window:   window,
		renderer: renderer,
		world:    world,
	}

	world.Inject(eng)

	window.SetOnResizeCallback(func(width, height int32) {
		eng.renderer.SetViewport(width, height)
		eng.SetResolution(width, height)
	})

	window.SetOnTickCallback(eng.tick)

	onInit(world)

	window.Run()
}

func (engine *Engine) GetWindow() Window {
	return engine.window
}

func (engine *Engine) GetRenderer() Renderer {
	return engine.renderer
}

func (engine *Engine) tick(deltaTime float32) {
	engine.world.Tick(deltaTime)
}

func (engine *Engine) SetResolution(width, height int32) {
	engine.Width = width
	engine.Height = height
}

func initSystem(world *World) {
	//r.Inject(BatchList{make([]Batch, 0)})

	engine := Resolve[Engine](world)

	// todo: set these options in shader
	engine.renderer.SetDepthTestMode(true)
	engine.renderer.SetCullingMode(true)
}

func batchSystem(world *World, deltaTime float32) {
	batchList := BatchList{
		Entity:  world.CreateEntity(),
		Batches: make(map[*Material][]BatchFrame),
	}
	q := world.Query().MeshComponent().TransformComponent().Fetch()
	for _, e := range q {
		mesh := world.GetMeshComponent(e)
		transform := world.GetTransformComponent(e)
		if _, ok := batchList.Batches[mesh.Material]; !ok {
			batchList.Batches = make(map[*Material][]BatchFrame)
		}
		batchList.Batches[mesh.Material] = append(batchList.Batches[mesh.Material], BatchFrame{
			Handle:    mesh.Handle,
			Transform: transform,
		})
	}

	batchList.Add(world)
}

func renderSystem(world *World, deltaTime float32) {
	engine := Resolve[Engine](world)

	engine.renderer.Clear(0.0, 0.5, 0.0, 1.0)

	for _, e := range world.Query().CameraComponent().Fetch() {
		camera := world.GetCameraComponent(e)

		projection := mgl32.Perspective(mgl32.DegToRad(camera.Fov), float32(engine.Width)/float32(engine.Height), 0.1, 100.0)
		view := camera.ViewMatrix

		for _, e := range world.Query().BatchList().Fetch() {
			batchList := world.GetBatchList(e)
			for material, frames := range batchList.Batches {
				engine.renderer.UseTexture(material.Texture0) // todo: set texture samplers in loop
				engine.renderer.UseProgram(material.Program)
				engine.renderer.SetGlobalInt(material.Program, "texture1", 0) // todo: set texture samplers in loop

				engine.renderer.SetGlobalMatrix(material.Program, "projection", projection[:])
				engine.renderer.SetGlobalMatrix(material.Program, "view", view[:])

				for _, frame := range frames {
					model := frame.Transform.GetMatrix()
					engine.renderer.SetGlobalMatrix(material.Program, "model", model[:])
					engine.renderer.UseMesh(frame.Handle.Id)
					engine.renderer.Draw(frame.Handle.IndexCount)
				}
			}
			batchList.Remove(world)
			// todo: check all components are removed
			// log batches (ingame console)
		}
	}
}

func updateCameraSystem(world *World, deltaTime float32) {
	//engine := Resolve[Engine](r)

	q := world.Query().CameraComponent().TransformComponent().Fetch()

	for _, e := range q {
		camera := world.GetCameraComponent(e)
		transform := world.GetTransformComponent(e)

		transform.Front = mgl32.Vec3{
			utils.Cos(mgl32.DegToRad(camera.Yaw)) * utils.Cos(mgl32.DegToRad(camera.Pitch)),
			utils.Sin(mgl32.DegToRad(camera.Pitch)),
			utils.Sin(mgl32.DegToRad(camera.Yaw)) * utils.Cos(mgl32.DegToRad(camera.Pitch)),
		}.Normalize()
		transform.Right = transform.Front.Cross(mgl32.Vec3{0, 1, 0}).Normalize()
		transform.Up = transform.Right.Cross(transform.Front).Normalize()
		camera.ViewMatrix = mgl32.LookAtV(transform.Position, transform.Position.Add(transform.Front), transform.Up)
	}
}
