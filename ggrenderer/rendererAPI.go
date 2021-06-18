package ggrenderer

import (
	"github.com/EngoEngine/glm"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/oyberntzen/gogame/ggcore"
)

//------------- Abstract -------------

type RendererAPI interface {
	Init()
	SetClearColor(*glm.Vec4)
	Clear()
	DrawIndexed(VertexArray)
	SetViewport(uint32, uint32, uint32, uint32)
}

type API int

const (
	RendererAPINone   API = 0
	RendererAPIOpenGL API = 1
)

var currentAPI API = RendererAPIOpenGL

func NewRendererAPI() RendererAPI {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLRendererAPI()
	}
	ggcore.CoreError("Unknown renderer API")
	return nil
}

func CurrentAPI() API {
	return currentAPI
}

//------------------------------------

//------------- Open GL --------------

type openGLRendererAPI struct {
}

func newOpenGLRendererAPI() *openGLRendererAPI {
	return &openGLRendererAPI{}
}

func (rendererAPI *openGLRendererAPI) Init() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
}

func (rendererAPI *openGLRendererAPI) SetClearColor(color *glm.Vec4) {
	gl.ClearColor(color[0], color[1], color[2], color[3])
}

func (rendererAPI *openGLRendererAPI) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (rendererAPI *openGLRendererAPI) DrawIndexed(vertexArray VertexArray) {
	gl.DrawElements(gl.TRIANGLES, int32(vertexArray.GetIndexBuffer().GetCount()), gl.UNSIGNED_INT, nil)
}

func (rendererAPI *openGLRendererAPI) SetViewport(x, y, width, height uint32) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

//------------------------------------
