// +build !windows

package ggcore

func init() {
	CoreWarn("GoGame does not support other platforms than windows")
}

func IsKeyPressed(keycode KeyCode) bool                { return false }
func IsMouseButtonPressed(button MouseButtonCode) bool { return false }
func GetMousePos() (float32, float32)                  { return 0, 0 }
func GetMouseX() float32                               { return 0 }
func GetMouseY() float32                               { return 0 }
