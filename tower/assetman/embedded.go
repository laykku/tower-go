package assetman

import (
	_ "embed"
)

//go:embed assets/shaders/vertex.shader
var vertexShader []byte

//go:embed assets/shaders/fragment.shader
var fragmentShader []byte

//go:embed assets/textures/checker.png
var checker []byte

//go:embed assets/meshes/cube.tmf
var cube []byte

func GetEmbeddedResource(path string) []byte {
	switch path {
	case "vertex_shader":
		return vertexShader
	case "fragment_shader":
		return fragmentShader
	case "checker_texture":
		return checker
	case "cube_mesh":
		return cube
	}
	return nil
}
