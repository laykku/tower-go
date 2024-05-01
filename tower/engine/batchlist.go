package engine

type BatchFrame struct {
	Handle    MeshHandle
	Transform *TransformComponent
}

type BatchList struct {
	Entity   int
	OneFrame bool
	Batches  map[*Material][]BatchFrame
}
