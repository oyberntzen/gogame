package ggrenderer

import "github.com/EngoEngine/glm"

var rendererAPI RendererAPI

func init() {
	rendererAPI = NewRendererAPI()
}

func RenderCommandInit() {
	rendererAPI.Init()
}

func RenderCommandSetClearColor(color *glm.Vec4) {
	rendererAPI.SetClearColor(color)
}

func RenderCommandClear() {
	rendererAPI.Clear()
}

func RenderCommandDrawIndexed(vertexArray VertexArray) {
	rendererAPI.DrawIndexed(vertexArray)
}

func RenderCommandSetViewport(x, y, width, height uint32) {
	rendererAPI.SetViewport(x, y, width, height)
}
