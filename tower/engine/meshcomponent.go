package engine

type MeshComponent struct {
	Entity   int
	OneFrame bool
	Data     *MeshData
	Material *Material
	Handle   MeshHandle
}
