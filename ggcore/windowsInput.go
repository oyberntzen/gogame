// +build windows

package ggcore

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/oyberntzen/gogame/ggevent"
)

func IsKeyPressed(keycode ggevent.KeyCode) bool {
	window := (*glfw.Window)(GetApp().GetWindow().GetNativeWindow())
	state := window.GetKey(glfw.Key(keycode))
	return state == glfw.Press || state == glfw.Repeat
}

func IsMouseButtonPressed(button ggevent.MouseButtonCode) bool {
	window := (*glfw.Window)(GetApp().GetWindow().GetNativeWindow())
	state := window.GetMouseButton(glfw.MouseButton(button))
	return state == glfw.Press
}

func GetMousePos() (float32, float32) {
	window := (*glfw.Window)(GetApp().GetWindow().GetNativeWindow())
	x, y := window.GetCursorPos()
	return float32(x), float32(y)
}

func GetMouseX() float32 {
	x, _ := GetMousePos()
	return x
}

func GetMouseY() float32 {
	_, y := GetMousePos()
	return y
}
