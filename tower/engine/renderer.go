package engine

type Renderer interface {
	SetViewport(width, height int32)

	CreateProgram(vertex, fragment string) any
	UseProgram(program any)

	SetGlobalMatrix(program any, name string, value []float32)
	SetGlobalInt(program any, name string, value int32)

	CreateMesh(vertices []float32, triangles []uint32) any
	UseMesh(vao any)

	CreateTexture(data []uint8, width int32, height int32) any
	UseTexture(texture any)

	SetDepthTestMode(value bool)
	SetCullingMode(value bool)

	Clear(r, g, b, a float32)
	Draw(count int32)
}
