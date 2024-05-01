package webgl

import (
	"fmt"
	"syscall/js"
)

type OpenGLRenderer struct {
	gl js.Value
}

var (
	VERTEX_SHADER        js.Value
	FRAGMENT_SHADER      js.Value
	ARRAY_BUFFER         js.Value
	STATIC_DRAW          js.Value
	FLOAT                js.Value
	ELEMENT_ARRAY_BUFFER js.Value
	TEXTURE0             js.Value
	TEXTURE_2D           js.Value
	TEXTURE_MIN_FILTER   js.Value
	TEXTURE_MAG_FILTER   js.Value
	LINEAR               js.Value
	TEXTURE_WRAP_S       js.Value
	TEXTURE_WRAP_T       js.Value
	CLAMP_TO_EDGE        js.Value
	COLOR_BUFFER_BIT     js.Value
	DEPTH_BUFFER_BIT     js.Value
	CULL_FACE            js.Value
	DEPTH_TEST           js.Value
	LESS                 js.Value
	TRIANGLES            js.Value
	UNSIGNED_INT         js.Value
	UNSIGNED_SHORT       js.Value
	UNSIGNED_BYTE        js.Value
	RGBA                 js.Value
	COMPILE_STATUS       js.Value
	LINK_STATUS          js.Value
)

func CreateOpenGLRenderer(glContext any) *OpenGLRenderer {
	gl, _ := glContext.(js.Value)

	VERTEX_SHADER = gl.Get("VERTEX_SHADER")
	FRAGMENT_SHADER = gl.Get("FRAGMENT_SHADER")
	ARRAY_BUFFER = gl.Get("ARRAY_BUFFER")
	STATIC_DRAW = gl.Get("STATIC_DRAW")
	FLOAT = gl.Get("FLOAT")
	ELEMENT_ARRAY_BUFFER = gl.Get("ELEMENT_ARRAY_BUFFER")
	TEXTURE0 = gl.Get("TEXTURE0")
	TEXTURE_2D = gl.Get("TEXTURE_2D")
	TEXTURE_MIN_FILTER = gl.Get("TEXTURE_MIN_FILTER")
	TEXTURE_MAG_FILTER = gl.Get("TEXTURE_MAG_FILTER")
	LINEAR = gl.Get("LINEAR")
	TEXTURE_WRAP_S = gl.Get("TEXTURE_WRAP_S")
	TEXTURE_WRAP_T = gl.Get("TEXTURE_WRAP_T")
	CLAMP_TO_EDGE = gl.Get("CLAMP_TO_EDGE")
	COLOR_BUFFER_BIT = gl.Get("COLOR_BUFFER_BIT")
	DEPTH_BUFFER_BIT = gl.Get("DEPTH_BUFFER_BIT")
	CULL_FACE = gl.Get("CULL_FACE")
	DEPTH_TEST = gl.Get("DEPTH_TEST")
	LESS = gl.Get("LESS")
	TRIANGLES = gl.Get("TRIANGLES")
	UNSIGNED_INT = gl.Get("UNSIGNED_INT")
	UNSIGNED_SHORT = gl.Get("UNSIGNED_SHORT")
	UNSIGNED_BYTE = gl.Get("UNSIGNED_BYTE")
	RGBA = gl.Get("RGBA")
	COMPILE_STATUS = gl.Get("COMPILE_STATUS")
	LINK_STATUS = gl.Get("LINK_STATUS")

	return &OpenGLRenderer{
		gl: gl,
	}
}

func (r *OpenGLRenderer) SetViewport(width, height int32) {
	r.gl.Call("viewport", 0, 0, width, height)
}

func (r *OpenGLRenderer) checkShaderCompileError(vertexShader js.Value) {
	if !r.gl.Call("getShaderParameter", vertexShader, COMPILE_STATUS).Bool() {
		infoLog := r.gl.Call("getShaderInfoLog", vertexShader).String()
		fmt.Println("Vertex shader compile error: ", infoLog)
	}
}

func (r *OpenGLRenderer) checkProgramLinkError(vertexShader js.Value) {
	if !r.gl.Call("getProgramParameter", vertexShader, LINK_STATUS).Bool() {
		infoLog := r.gl.Call("getProgramInfoLog", vertexShader).String()
		fmt.Println("Program link error: ", infoLog)
	}
}

func (r *OpenGLRenderer) CreateProgram(vertex, fragment string) any {

	vertexShader := r.gl.Call("createShader", VERTEX_SHADER)
	r.gl.Call("shaderSource", vertexShader, vertex)
	r.gl.Call("compileShader", vertexShader)
	r.checkShaderCompileError(vertexShader)

	fragmentShader := r.gl.Call("createShader", FRAGMENT_SHADER)
	r.gl.Call("shaderSource", fragmentShader, fragment)
	r.gl.Call("compileShader", fragmentShader)
	r.checkShaderCompileError(fragmentShader)

	// todo: check compile/link errors

	program := r.gl.Call("createProgram")

	r.gl.Call("attachShader", program, vertexShader)
	r.gl.Call("attachShader", program, fragmentShader)
	r.gl.Call("linkProgram", program)
	r.checkProgramLinkError(program)

	r.gl.Call("deleteShader", vertexShader)
	r.gl.Call("deleteShader", fragmentShader)

	return program
}

func (r *OpenGLRenderer) UseProgram(program any) {
	r.gl.Call("useProgram", program)
}

func (r *OpenGLRenderer) GetAttribLocation(program any, name string) any {
	return r.gl.Call("getAttribLocation", program, name)
}

func (r *OpenGLRenderer) CreateMesh(vertices []float32, triangles []uint32) any {
	vao := r.gl.Call("createVertexArray")
	r.gl.Call("bindVertexArray", vao)

	vbo := r.gl.Call("createBuffer", ARRAY_BUFFER)
	r.gl.Call("bindBuffer", ARRAY_BUFFER, vbo)
	r.gl.Call("bufferData", ARRAY_BUFFER, sliceToTypedArray(vertices), STATIC_DRAW)

	r.gl.Call("vertexAttribPointer", 0, 3, FLOAT, false, 5*4, 0)
	r.gl.Call("enableVertexAttribArray", 0)

	r.gl.Call("vertexAttribPointer", 1, 2, FLOAT, false, 5*4, 3*4)
	r.gl.Call("enableVertexAttribArray", 1)

	indexBuffer := r.gl.Call("createBuffer", ELEMENT_ARRAY_BUFFER)
	r.gl.Call("bindBuffer", ELEMENT_ARRAY_BUFFER, indexBuffer)
	r.gl.Call("bufferData", ELEMENT_ARRAY_BUFFER, sliceToTypedArray(triangles), STATIC_DRAW)

	r.gl.Call("bindVertexArray", js.Null())

	return vao
}

func (r *OpenGLRenderer) UseMesh(vao any) {
	r.gl.Call("bindVertexArray", vao)
}

func (r *OpenGLRenderer) GetUniformLocation(program any, name string) any {
	return r.gl.Call("getUniformLocation", program, name)
}

func (r *OpenGLRenderer) SetGlobalMatrix(program any, name string, value []float32) {
	location := r.GetUniformLocation(program, name)
	r.gl.Call("uniformMatrix4fv", location, false, sliceToTypedArray(value))
}

func (r *OpenGLRenderer) SetGlobalInt(program any, name string, value int32) {
	location := r.gl.Call("getUniformLocation", program, name)
	r.gl.Call("uniform1i", location, value)
}

func (r *OpenGLRenderer) CreateTexture(data []uint8, width int32, height int32) any {
	texture := r.gl.Call("createTexture")
	r.gl.Call("activeTexture", TEXTURE0)
	r.gl.Call("bindTexture", TEXTURE_2D, texture)
	r.gl.Call("texParameteri", TEXTURE_2D, TEXTURE_MIN_FILTER, LINEAR)
	r.gl.Call("texParameteri", TEXTURE_2D, TEXTURE_MAG_FILTER, LINEAR)
	r.gl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_S, CLAMP_TO_EDGE)
	r.gl.Call("texParameteri", TEXTURE_2D, TEXTURE_WRAP_T, CLAMP_TO_EDGE)
	r.gl.Call("texImage2D",
		TEXTURE_2D,
		0,
		RGBA,
		width,
		height,
		0,
		RGBA,
		UNSIGNED_BYTE,
		sliceToTypedArray(data))
	r.gl.Call("generateMipmap", TEXTURE_2D)

	return texture
}

func (r *OpenGLRenderer) UseTexture(texture any) {
	r.gl.Call("activeTexture", TEXTURE0)
	r.gl.Call("bindTexture", TEXTURE_2D, texture)
}

func (ren *OpenGLRenderer) Clear(r, g, b, a float32) {
	ren.gl.Call("clearColor", r, g, b, a)
	ren.gl.Call("clear", COLOR_BUFFER_BIT)
	ren.gl.Call("clear", DEPTH_BUFFER_BIT)
}

func (r *OpenGLRenderer) SetCullingMode(value bool) {
	if value {
		r.gl.Call("enable", CULL_FACE) // todo: move to shader
		r.gl.Call("clearDepth", 1.0)   // todo: move to shader
	} else {
		r.gl.Call("disable", CULL_FACE)
	}
}

func (r *OpenGLRenderer) SetDepthTestMode(value bool) {
	if value {
		r.gl.Call("enable", DEPTH_TEST) // todo: move to shader
		r.gl.Call("depthFunc", LESS)    // todo: move to shader
	} else {
		r.gl.Call("disable", DEPTH_TEST)
	}
}

func (r *OpenGLRenderer) Draw(count int32) {
	r.gl.Call("drawElements", TRIANGLES, count, UNSIGNED_INT, 0)
}
