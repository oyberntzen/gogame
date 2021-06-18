package gogame

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggevent"
	"github.com/oyberntzen/gogame/ggimgui"
	"github.com/oyberntzen/gogame/ggrenderer"
	"github.com/oyberntzen/gogame/ggwindow"
)

type ClientApplication interface {
	Init(*CoreApplication)
}

type CoreApplication struct {
	clientApp     ClientApplication
	window        ggcore.Window
	running       bool
	layerStack    ggcore.LayerStack
	imguiLayer    *ggimgui.ImGuiLayer
	lastFrameTime float32
	minimized     bool
}

func newCoreApplication(clientApp ClientApplication) *CoreApplication {
	props := ggcore.WindowProps{Title: "GoGame", Width: 1280, Height: 720}
	app := CoreApplication{clientApp: clientApp, window: ggwindow.NewWindow(props), running: true}
	ggcore.SetApp(&app)

	app.window.SetEventCallback(app.onEvent)

	ggrenderer.RendererInit()

	app.imguiLayer = &ggimgui.ImGuiLayer{}
	app.PushOverlay(app.imguiLayer)

	return &app
}

func (app *CoreApplication) PushLayer(layer ggcore.Layer) {
	app.layerStack.PushLayer(layer)
}

func (app *CoreApplication) PushOverlay(layer ggcore.Layer) {
	app.layerStack.PushOverlay(layer)
}

func (app *CoreApplication) GetWindow() ggcore.Window {
	return app.window
}

func (app *CoreApplication) run() {
	app.clientApp.Init(app)

	for app.running {
		time := float32(glfw.GetTime())
		timestep := ggcore.Timestep(time - app.lastFrameTime)
		app.lastFrameTime = float32(time)

		if !app.minimized {
			for _, layer := range app.layerStack.Layers {
				layer.OnUpdate(timestep)
			}

			app.imguiLayer.Begin()
			for _, layer := range app.layerStack.Layers {
				layer.OnImGuiRender()
			}
			app.imguiLayer.End()
		}

		app.window.OnUpdate()
	}

	ggcore.CoreInfo("Stopping application")
}

func (app *CoreApplication) onEvent(event ggevent.Event) {
	dispatcher := ggevent.EventDispatcher{Event: event}

	dispatcher.DispatchWindowClose(app.onWindowClose)
	dispatcher.DispatchWindowResize(app.onWindowResize)

	for i := len(app.layerStack.Layers) - 1; i >= 0; i-- {
		app.layerStack.Layers[i].OnEvent(event)
		if event.IsHandled() {
			break
		}
	}
}

func (app *CoreApplication) onWindowClose(event *ggevent.WindowCloseEvent) bool {
	app.running = false
	return true
}

func (app *CoreApplication) onWindowResize(event *ggevent.WindowResizeEvent) bool {
	if event.Width() == 0 || event.Height() == 0 {
		app.minimized = true
		return false
	}
	app.minimized = false
	ggrenderer.RendererOnWindowResize(uint32(event.Width()), uint32(event.Height()))
	return false
}
