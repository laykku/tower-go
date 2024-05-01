package engine

type MeshData struct {
	Positions []float32
	Uv0       []float32
	Triangles []uint32
}

type MeshHandle struct {
	Id         any
	IndexCount int32
}

func (meshData *MeshData) CreateHandle(renderer Renderer) MeshHandle {
	const vertexDataSize = 5

	vertexData := make([]float32, len(meshData.Triangles)*vertexDataSize)
	indices := make([]uint32, len(meshData.Triangles))

	for i, index := range meshData.Triangles {
		vertexData[i*vertexDataSize] = meshData.Positions[index*3]
		vertexData[i*vertexDataSize+1] = meshData.Positions[index*3+1]
		vertexData[i*vertexDataSize+2] = meshData.Positions[index*3+2]
		vertexData[i*vertexDataSize+3] = meshData.Uv0[i*2]
		vertexData[i*vertexDataSize+4] = meshData.Uv0[i*2+1]
		indices[i] = uint32(i)
	}

	meshId := renderer.CreateMesh(vertexData, indices)

	return MeshHandle{
		Id:         meshId,
		IndexCount: int32(len(meshData.Triangles))}
}
