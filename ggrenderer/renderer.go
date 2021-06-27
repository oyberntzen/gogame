package ggrenderer

import (
	"github.com/EngoEngine/glm"
	"github.com/oyberntzen/gogame/ggdebug"
)

var viewProjectionMatrix glm.Mat4

func RendererInit() {
	RenderCommandInit()
	Renderer2DInit()
}

func RendererBeginScene(camera *OrthographicCamera) {
	defer ggdebug.Stop(ggdebug.Start())

	viewProjectionMatrix = *camera.GetViewProjectionMatrix()
}

func RendererEndScene() {
	defer ggdebug.Stop(ggdebug.Start())
}

func RendererSubmit(shader Shader, vertexArray VertexArray, transform *glm.Mat4) {
	defer ggdebug.Stop(ggdebug.Start())

	shader.Bind()
	shader.(*OpenGLShader).UploadUniformMat4("u_ViewProjection", &viewProjectionMatrix)
	shader.(*OpenGLShader).UploadUniformMat4("u_Transform", transform)

	vertexArray.Bind()
	RenderCommandDrawIndexed(vertexArray, 0)
}

func RendererOnWindowResize(width, height uint32) {
	defer ggdebug.Stop(ggdebug.Start())

	RenderCommandSetViewport(0, 0, width, height)
}
