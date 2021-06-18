// +build windows

package ggwindow

import (
	"runtime"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/oyberntzen/gogame/ggconstants"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggevent"
	"github.com/oyberntzen/gogame/ggrenderer"
)

func init() {
	runtime.LockOSThread()
}

var (
	glfwInitialized bool = false
	windowDatas     map[int]*WindowData
	windowDataID    int
	dataIDs         []int = []int{0}
)

type WindowData struct {
	Title             string
	Width, Height     uint
	VSync             bool
	EventCallbackFunc ggcore.EventCallbackFunc
}

type WindowsWindow struct {
	window  *glfw.Window
	context ggrenderer.GraphicsContext
	data    WindowData
	dataID  int
}

func init() {
	windowDatas = make(map[int]*WindowData)
}

func addWindowData(data *WindowData) int {
	windowDataID++
	windowDatas[windowDataID] = data
	dataIDs = append(dataIDs, windowDataID)
	return windowDataID
}

func getWindowData(id int) *WindowData {
	return windowDatas[id]
}

func NewWindow(props ggcore.WindowProps) ggcore.Window {
	window := WindowsWindow{}
	window.init(props)
	return &window
}

func (window *WindowsWindow) OnUpdate() {
	glfw.PollEvents()
	window.context.SwapBuffers()

	//id := window.dataID
	//window.window.SetUserPointer(unsafe.Pointer(&id))
}

func (window *WindowsWindow) Width() uint  { return window.data.Width }
func (window *WindowsWindow) Height() uint { return window.data.Height }

func (window *WindowsWindow) SetEventCallback(function ggcore.EventCallbackFunc) {
	window.data.EventCallbackFunc = function
}

func (window *WindowsWindow) SetVSync(enabled bool) {
	if enabled {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}
	window.data.VSync = enabled
}

func (window *WindowsWindow) VSync() bool {
	return window.data.VSync
}

func (window *WindowsWindow) GetNativeWindow() unsafe.Pointer {
	return unsafe.Pointer(window.window)
}

func (window *WindowsWindow) init(props ggcore.WindowProps) {
	window.data = WindowData{Title: props.Title, Width: props.Width, Height: props.Height}
	window.SetEventCallback(func(event ggevent.Event) {
		ggcore.CoreWarn("Event callback function not defined")
	})

	ggcore.CoreInfo("Creating window: %v (%v, %v)", window.data.Title, window.data.Width, window.data.Height)

	if !glfwInitialized {
		ggcore.CoreCheckError(glfw.Init())
		glfwInitialized = true
	}

	var err error
	window.window, err = glfw.CreateWindow(int(window.data.Width), int(window.data.Height), window.data.Title, nil, nil)
	ggcore.CoreCheckError(err)

	window.context = ggrenderer.NewOpenGLContext(window.window)

	window.context.Init()

	window.dataID = addWindowData(&window.data)
	window.window.SetUserPointer(unsafe.Pointer(&dataIDs[windowDataID]))
	window.SetVSync(false)

	window.window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		data.Width = uint(width)
		data.Height = uint(height)
		event := ggevent.NewWindowResizeEvent(width, height)
		data.EventCallbackFunc(event)
	})

	window.window.SetCloseCallback(func(w *glfw.Window) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		event := ggevent.NewWindowCloseEvent()
		data.EventCallbackFunc(event)
	})

	window.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		switch action {
		case glfw.Press:
			{
				event := ggevent.NewKeyPressedEvent(ggconstants.KeyCode(key), 0)
				data.EventCallbackFunc(event)
			}
		case glfw.Release:
			{
				event := ggevent.NewKeyReleasedEvent(ggconstants.KeyCode(key))
				data.EventCallbackFunc(event)
			}
		case glfw.Repeat:
			{
				event := ggevent.NewKeyPressedEvent(ggconstants.KeyCode(key), 1)
				data.EventCallbackFunc(event)
			}
		}
	})

	window.window.SetCharCallback(func(w *glfw.Window, char rune) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		event := ggevent.NewKeyTypedEvent(char)
		data.EventCallbackFunc(event)
	})

	window.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		switch action {
		case glfw.Press:
			{
				event := ggevent.NewMouseButtonPressedEvent(ggconstants.MouseButtonCode(button))
				data.EventCallbackFunc(event)
			}
		case glfw.Release:
			{
				event := ggevent.NewMouseButtonReleasedEvent(ggconstants.MouseButtonCode(button))
				data.EventCallbackFunc(event)
			}
		}
	})

	window.window.SetScrollCallback(func(w *glfw.Window, xpos, ypos float64) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		event := ggevent.NewMouseScrolledEvent(float32(xpos), float32(ypos))
		data.EventCallbackFunc(event)
	})

	window.window.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
		id := (*int)(w.GetUserPointer())
		data := getWindowData(*id)
		event := ggevent.NewMouseMovedEvent(float32(xpos), float32(ypos))
		data.EventCallbackFunc(event)
	})
}

func (window *WindowsWindow) Delete() {
	window.window.Destroy()
}
