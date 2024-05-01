package gen

type Component interface {
	GetEntity() int
	OneFrame() bool
}
