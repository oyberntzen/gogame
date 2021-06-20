package ggrenderer

import (
	"github.com/EngoEngine/glm"
)

var (
	quadVertexArray VertexArray
	flatColorShader Shader
)

func Renderer2DInit() {

	vertexBuffer := NewVertexBuffer([]float32{
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		0.5, 0.5, 0,
		-0.5, 0.5, 0,
	})
	vertexBuffer.SetLayout(NewBufferLayout([]*BufferElement{
		NewBufferElement(ShaderDataTypeFloat3, "a_Position", false),
	}))

	indexBuffer := NewIndexBuffer([]uint32{
		0, 1, 2, 0, 2, 3,
	})

	quadVertexArray = NewVertexArray()
	quadVertexArray.AddVertexBuffer(vertexBuffer)
	quadVertexArray.SetIndexBuffer(indexBuffer)

	vertex := `
	#version 330 core

	layout(location = 0) in vec3 a_Position;

	uniform mat4 u_ViewProjection;
	uniform mat4 u_Transform;

	void main() 
	{
		gl_Position = u_ViewProjection * u_Transform * vec4(a_Position, 1.0);
	}
	`

	fragment := `
	#version 330 core

	layout(location = 0) out vec4 color;

	uniform vec4 u_Color;

	void main() 
	{
		color = vec4(u_Color);
	}
	`

	flatColorShader = NewShaderFromSrc("flatColor", vertex, fragment)
}

func Renderer2DBeginShutdown() {
	quadVertexArray.Delete()
	flatColorShader.Delete()
}

func Renderer2DBeginScene(camera *OrthographicCamera) {
	flatColorShader.Bind()
	flatColorShader.(*OpenGLShader).UploadUniformMat4("u_ViewProjection", camera.GetViewProjectionMatrix())

}

func Renderer2DEndScene() {

}

func Renderer2DDrawQuad(position *glm.Vec2, size *glm.Vec2, color *glm.Vec4) {
	Renderer2DDrawQuadWithZ(&glm.Vec3{position[0], position[1], 0}, size, color)
}

func Renderer2DDrawQuadWithZ(position *glm.Vec3, size *glm.Vec2, color *glm.Vec4) {
	flatColorShader.Bind()

	transform := glm.Translate3D(position[0], position[1], position[2])
	scale := glm.Scale3D(size[0], size[1], 1)
	transform.Mul4With(&scale)
	flatColorShader.(*OpenGLShader).UploadUniformMat4("u_Transform", &transform)
	flatColorShader.(*OpenGLShader).UploadUniformFloat4("u_Color", color)

	RenderCommandDrawIndexed(quadVertexArray)
}
