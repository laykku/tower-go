package ecs

type MouseMove struct {
	Entity       int
	OneFrame     bool
	XPrev, YPrev int
	XRel, YRel   int
}

type MoveInput struct {
	Entity   int
	OneFrame bool
	Forward  int
	Right    int
}
