package engine

import "github.com/go-gl/mathgl/mgl32"

type TransformComponent struct {
	Entity   int
	OneFrame bool
	Position mgl32.Vec3
	Rotation mgl32.Quat
	Scale    mgl32.Vec3
	Front    mgl32.Vec3
	Right    mgl32.Vec3
	Up       mgl32.Vec3
}

func (t *TransformComponent) GetMatrix() mgl32.Mat4 {
	translation := mgl32.Translate3D(t.Position.X(), t.Position.Y(), t.Position.Z())
	rotation := t.Rotation
	scale := mgl32.Scale3D(t.Scale.X(), t.Scale.Y(), t.Scale.Z())
	return translation.Mul4(rotation.Mat4().Mul4(scale))
}
