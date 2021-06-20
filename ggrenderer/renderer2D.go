package ggrenderer

import (
	"unsafe"

	"github.com/EngoEngine/glm"
)

var (
	quadVertexArray VertexArray
	textureShader   Shader
	whiteTexture    Texture
)

type Quad2D struct {
	Position          *glm.Vec2
	Z                 float32
	Size              *glm.Vec2
	Color             *glm.Vec4
	UseColorOnTexture bool
	Texture           Texture
}

func Renderer2DInit() {

	vertexBuffer := NewVertexBuffer([]float32{
		-0.5, -0.5, 0, 0, 0,
		0.5, -0.5, 0, 1, 0,
		0.5, 0.5, 0, 1, 1,
		-0.5, 0.5, 0, 0, 1,
	})
	vertexBuffer.SetLayout(NewBufferLayout([]*BufferElement{
		NewBufferElement(ShaderDataTypeFloat3, "a_Position", false),
		NewBufferElement(ShaderDataTypeFloat2, "a_TexCoord", false),
	}))

	indexBuffer := NewIndexBuffer([]uint32{
		0, 1, 2, 0, 2, 3,
	})

	quadVertexArray = NewVertexArray()
	quadVertexArray.AddVertexBuffer(vertexBuffer)
	quadVertexArray.SetIndexBuffer(indexBuffer)

	whiteTexture = NewTexture2DEmpty(1, 1)
	whiteTexture.SetData(unsafe.Pointer(&[4]uint8{255, 255, 255, 255}))

	vertex := `
	#version 330 core

	layout(location = 0) in vec3 a_Position;
	layout(location = 1) in vec2 a_TexCoord;

	uniform mat4 u_ViewProjection;
	uniform mat4 u_Transform;

	out vec2 v_TexCoord;

	void main() 
	{
		v_TexCoord = a_TexCoord;
		gl_Position = u_ViewProjection * u_Transform * vec4(a_Position, 1.0);
	}
	`

	fragment := `
	#version 330 core

	layout(location = 0) out vec4 color;

	uniform vec4 u_Color;
	uniform sampler2D u_Texture;

	in vec2 v_TexCoord;

	void main() 
	{
		color = texture(u_Texture, v_TexCoord) * u_Color;
	}
	`

	textureShader = NewShaderFromSrc("texture", vertex, fragment)
}

func Renderer2DBeginShutdown() {
	quadVertexArray.Delete()
}

func Renderer2DBeginScene(camera *OrthographicCamera) {
	textureShader.Bind()
	textureShader.(*OpenGLShader).UploadUniformMat4("u_ViewProjection", camera.GetViewProjectionMatrix())
	textureShader.(*OpenGLShader).UploadUniformInt("u_Texture", 0)
}

func Renderer2DEndScene() {

}

func Renderer2DDrawQuad(quad *Quad2D) {
	if quad.Texture == nil || quad.UseColorOnTexture {
		textureShader.(*OpenGLShader).UploadUniformFloat4("u_Color", quad.Color)
	} else {
		textureShader.(*OpenGLShader).UploadUniformFloat4("u_Color", &glm.Vec4{1, 1, 1, 1})
	}

	if quad.Texture == nil {
		whiteTexture.Bind(0)
	} else {
		quad.Texture.Bind(0)
	}

	transform := glm.Translate3D(quad.Position[0], quad.Position[1], quad.Z)
	scale := glm.Scale3D(quad.Size[0], quad.Size[1], 1)
	transform.Mul4With(&scale)
	textureShader.(*OpenGLShader).UploadUniformMat4("u_Transform", &transform)

	RenderCommandDrawIndexed(quadVertexArray)
}
