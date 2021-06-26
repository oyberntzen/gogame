package ggrenderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
)

//------------- Abstract -------------

type GraphicsContext interface {
	Init()
	SwapBuffers()
}

//------------------------------------

//------------- Open GL --------------

type OpenGLContext struct {
	window *glfw.Window
}

func NewOpenGLContext(window *glfw.Window) *OpenGLContext {
	return &OpenGLContext{window: window}
}

func (context *OpenGLContext) Init() {
	defer ggdebug.Stop(ggdebug.Start())

	context.window.MakeContextCurrent()
	ggcore.CoreCheckError(gl.Init())

	ggcore.CoreInfo("OpenGL info:")
	ggcore.CoreInfo("	Vendor:   %v", gl.GoStr(gl.GetString(gl.VENDOR)))
	ggcore.CoreInfo("	Renderer: %v", gl.GoStr(gl.GetString(gl.RENDERER)))
	ggcore.CoreInfo("	Version:  %v", gl.GoStr(gl.GetString(gl.VERSION)))
}

func (context *OpenGLContext) SwapBuffers() {
	defer ggdebug.Stop(ggdebug.Start())

	context.window.SwapBuffers()
}

//------------------------------------
