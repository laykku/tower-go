package engine

type Window interface {
	GetGlContext() any
	SetOnResizeCallback(callback func(width, height int32))
	SetOnTickCallback(callback func(float32))
	Run()
	IsKeyPressed(key string) bool
	GetMouseState() (int, int)
}
