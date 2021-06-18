package ggrenderer

import "github.com/EngoEngine/glm"

type sceneData struct {
	viewProjectionMatrix glm.Mat4
}

var data sceneData

func RendererInit() {
	RenderCommandInit()
}

func RendererBeginScene(camera *OrthographicCamera) {
	data.viewProjectionMatrix = *camera.GetViewProjectionMatrix()
}

func RendererEndScene() {

}

func RendererSubmit(shader Shader, vertexArray VertexArray, transform *glm.Mat4) {
	shader.Bind()
	shader.(*OpenGLShader).UploadUniformMat4("u_ViewProjection", &data.viewProjectionMatrix)
	shader.(*OpenGLShader).UploadUniformMat4("u_Transform", transform)

	vertexArray.Bind()
	RenderCommandDrawIndexed(vertexArray)
}

func RendererOnWindowResize(width, height uint32) {
	RenderCommandSetViewport(0, 0, width, height)
}
