package ggimgui

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
	"github.com/oyberntzen/gogame/ggevent"
	"github.com/oyberntzen/gogame/ggimgui/imguiexamples"
)

type ImGuiLayer struct {
	context       *imgui.Context
	imguiRenderer *imguiexamples.OpenGL3
	imguiPlatform *imguiexamples.GLFW
}

func (layer *ImGuiLayer) OnAttach() {
	defer ggdebug.Stop(ggdebug.Start())

	layer.context = imgui.CreateContext(nil)

	io := imgui.CurrentIO()
	io.SetBackendFlags(imgui.BackendFlagsHasMouseCursors | imgui.BackendFlagsHasSetMousePos)
	io.SetConfigFlags(imgui.ConfigFlagsDockingEnable)

	layer.imguiRenderer = imguiexamples.NewOpenGL3(io)

	layer.imguiPlatform = imguiexamples.NewGLFW(io, (*glfw.Window)(ggcore.GetApp().GetWindow().GetNativeWindow()))
}

func (layer *ImGuiLayer) OnDetach() {
	layer.imguiRenderer.Dispose()
	layer.context.Destroy()
}

func (layer *ImGuiLayer) OnUpdate(timestep ggcore.Timestep) {}

func (layer *ImGuiLayer) OnImGuiRender() {}

func (layer *ImGuiLayer) OnEvent(event ggevent.Event) {}

func (layer *ImGuiLayer) GetName() string { return "imgui" }

func (layer *ImGuiLayer) Begin() {
	defer ggdebug.Stop(ggdebug.Start())

	layer.imguiPlatform.NewFrame()
	imgui.NewFrame()
}

func (layer *ImGuiLayer) End() {
	defer ggdebug.Stop(ggdebug.Start())

	imgui.Render()
	layer.imguiRenderer.Render(layer.imguiPlatform.DisplaySize(), layer.imguiPlatform.FramebufferSize(), imgui.RenderedDrawData())
}
