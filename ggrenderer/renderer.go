package ggrenderer

import "github.com/EngoEngine/glm"

var viewProjectionMatrix glm.Mat4

func RendererInit() {
	RenderCommandInit()
	Renderer2DInit()
}

func RendererBeginScene(camera *OrthographicCamera) {
	viewProjectionMatrix = *camera.GetViewProjectionMatrix()
}

func RendererEndScene() {

}

func RendererSubmit(shader Shader, vertexArray VertexArray, transform *glm.Mat4) {
	shader.Bind()
	shader.(*OpenGLShader).UploadUniformMat4("u_ViewProjection", &viewProjectionMatrix)
	shader.(*OpenGLShader).UploadUniformMat4("u_Transform", transform)

	vertexArray.Bind()
	RenderCommandDrawIndexed(vertexArray)
}

func RendererOnWindowResize(width, height uint32) {
	RenderCommandSetViewport(0, 0, width, height)
}
