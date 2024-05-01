package engine

import "github.com/go-gl/mathgl/mgl32"

type CameraComponent struct {
	Entity     int
	OneFrame   bool
	Pitch      float32
	Yaw        float32
	Fov        float32
	ViewMatrix mgl32.Mat4
}
