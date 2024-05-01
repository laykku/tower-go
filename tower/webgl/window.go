package webgl

import (
	"fmt"
	"syscall/js"
	"time"
)

var global js.Value = js.Global()

type WebGlWindow struct {
	doc    js.Value
	canvas js.Value
	gl     js.Value

	deltaTime float32
	prevTime  time.Time

	pressedKeys map[string]bool
	mouseX      int
	mouseY      int

	onResizeCallback func(int32, int32)
	onTickCallback   func(float32)
}

func CreateWindow() *WebGlWindow {

	doc := global.Get("document")
	canvas := doc.Call("getElementById", "app-canvas")

	gl := canvas.Call("getContext", "webgl2")
	if gl.IsUndefined() {
		fmt.Println("failed to get webgl context")
	}

	window := &WebGlWindow{
		doc:         doc,
		canvas:      canvas,
		gl:          gl,
		pressedKeys: make(map[string]bool),
	}

	global.Call("addEventListener", "resize", js.FuncOf(window.onResize))
	global.Call("addEventListener", "keydown", js.FuncOf(window.onKeyPress))
	global.Call("addEventListener", "keyup", js.FuncOf(window.onKeyRelease))
	global.Call("addEventListener", "mousemove", js.FuncOf(window.onMouseMove))

	window.canvas.Call("addEventListener", "click", js.FuncOf(window.requestPointerLock))

	return window
}

func (window *WebGlWindow) Run() {
	global.Call("requestAnimationFrame", js.FuncOf(window.tick))
	select {}
}

func (window *WebGlWindow) IsKeyPressed(key string) bool {
	ok, state := window.pressedKeys[key]
	return ok && state
}

func (window *WebGlWindow) GetMouseState() (int, int) {
	return window.mouseX, window.mouseY
}

func (window *WebGlWindow) GetGlContext() any {
	return window.gl
}

func (window *WebGlWindow) SetOnResizeCallback(callback func(width, height int32)) {
	window.onResizeCallback = callback
	window.updateViewport()
}

func (window *WebGlWindow) SetOnTickCallback(callback func(float32)) {
	window.onTickCallback = callback
}

func (window *WebGlWindow) tick(this js.Value, p []js.Value) any {

	currentTime := time.Now()
	window.deltaTime = float32(currentTime.Sub(window.prevTime).Seconds())
	window.prevTime = currentTime

	window.onTickCallback(window.deltaTime)

	global.Call("requestAnimationFrame", js.FuncOf(window.tick))

	window.mouseX = 0
	window.mouseY = 0

	return nil
}

func (window *WebGlWindow) requestPointerLock(this js.Value, p []js.Value) any {
	window.canvas.Call("requestPointerLock")
	return nil
}

func (window *WebGlWindow) onKeyPress(this js.Value, p []js.Value) any {
	key := p[0].Get("key").String()
	window.pressedKeys[key] = true
	return nil
}

func (window *WebGlWindow) onKeyRelease(this js.Value, p []js.Value) any {
	key := p[0].Get("key").String()
	window.pressedKeys[key] = false
	return nil
}

func (window *WebGlWindow) onMouseMove(this js.Value, p []js.Value) any {
	event := p[0]
	window.mouseX = event.Get("movementX").Int()
	window.mouseY = event.Get("movementY").Int()
	return nil
}

func (window *WebGlWindow) onResize(this js.Value, p []js.Value) any {
	window.updateViewport()
	return nil
}

func (window *WebGlWindow) updateViewport() {
	width := int32(window.canvas.Get("clientWidth").Int())
	height := int32(window.canvas.Get("clientHeight").Int())
	window.canvas.Call("setAttribute", "width", width)
	window.canvas.Call("setAttribute", "height", height)
	window.onResizeCallback(width, height)
}
