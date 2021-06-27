package ggrenderer

import (
	"unsafe"

	"github.com/EngoEngine/glm"
	"github.com/oyberntzen/gogame/ggdebug"
)

const (
	maxQuads        uint32 = 10_000
	maxVertices     uint32 = maxQuads * 4
	maxIndices      uint32 = maxQuads * 6
	maxTextureSlots uint32 = 32
)

var (
	quadVertexArray  VertexArray
	quadVertexBuffer VertexBuffer
	textureShader    Shader
	whiteTexture     Texture

	quadIndexCount uint32 = 0
	quadVertices   []quadVertex

	textureSlots     [maxTextureSlots]Texture
	textureSlotIndex uint32 = 1
)

type Quad2D struct {
	Position     *glm.Vec2
	Z            float32
	Size         *glm.Vec2
	Rotation     float32
	Color        *glm.Vec4
	Texture      Texture
	TilingFactor float32
}

func NewQuad2D() *Quad2D {
	return &Quad2D{
		Position:     &glm.Vec2{0, 0},
		Z:            0,
		Size:         &glm.Vec2{1, 1},
		Rotation:     0,
		Color:        &glm.Vec4{1, 1, 1, 1},
		Texture:      nil,
		TilingFactor: 1,
	}
}

type quadVertex struct {
	Position glm.Vec3
	Color    glm.Vec4
	TexCoord glm.Vec2
	TexIndex float32
}

func Renderer2DInit() {
	defer ggdebug.Stop(ggdebug.Start())

	quadVertices = make([]quadVertex, 0, maxQuads)

	quadVertexBuffer = NewEmptyVertexBuffer(maxVertices * uint32(unsafe.Sizeof(quadVertex{})))
	quadVertexBuffer.SetLayout(NewBufferLayout([]*BufferElement{
		NewBufferElement(ShaderDataTypeFloat3, "a_Position", false),
		NewBufferElement(ShaderDataTypeFloat4, "a_Color", false),
		NewBufferElement(ShaderDataTypeFloat2, "a_TexCoord", false),
		NewBufferElement(ShaderDataTypeFloat, "a_TexIndex", false),
	}))

	indices := make([]uint32, maxIndices)
	var offset uint32 = 0
	for i := 0; i < int(maxIndices); i += 6 {
		indices[i+0] = offset + 0
		indices[i+1] = offset + 1
		indices[i+2] = offset + 2

		indices[i+3] = offset + 2
		indices[i+4] = offset + 3
		indices[i+5] = offset + 0

		offset += 4
	}
	indexBuffer := NewIndexBuffer(indices)

	quadVertexArray = NewVertexArray()
	quadVertexArray.AddVertexBuffer(quadVertexBuffer)
	quadVertexArray.SetIndexBuffer(indexBuffer)

	whiteTexture = NewTexture2DEmpty(1, 1)
	whiteTexture.SetData(unsafe.Pointer(&[4]uint8{255, 255, 255, 255}))
	textureSlots[0] = whiteTexture

	vertex := `
	#version 330 core

	layout(location = 0) in vec3  a_Position;
	layout(location = 1) in vec4  a_Color;
	layout(location = 2) in vec2  a_TexCoord;
	layout(location = 3) in float a_TexIndex;

	uniform mat4 u_ViewProjection;

	out vec4  v_Color;
	out vec2  v_TexCoord;
	out float v_TexIndex;

	void main() 
	{
		v_Color = a_Color;
		v_TexCoord = a_TexCoord;
		v_TexIndex = float(a_TexIndex);
		gl_Position = u_ViewProjection * vec4(a_Position, 1.0);
	}
	`

	fragment := `
	#version 330 core

	layout(location = 0) out vec4 color;

	uniform sampler2D u_Textures[32];

	in vec4  v_Color;
	in vec2  v_TexCoord;
	in float v_TexIndex;

	void main() 
	{
		color = texture(u_Textures[int(v_TexIndex)], v_TexCoord) * v_Color;
		//color = vec4(v_TexIndex, 1, 1, 1);
	}
	`

	textureShader = NewShaderFromSrc("texture", vertex, fragment)

	samplers := make([]int32, maxTextureSlots)
	for i := 0; i < len(samplers); i++ {
		samplers[i] = int32(i)
	}

	textureShader.Bind()
	textureShader.(*OpenGLShader).UploadUniformIntArray("u_Textures", samplers)
}

func Renderer2DShutdown() {
	defer ggdebug.Stop(ggdebug.Start())

	quadVertexArray.Delete()
}

func Renderer2DBeginScene(camera *OrthographicCamera) {
	defer ggdebug.Stop(ggdebug.Start())

	textureShader.Bind()
	textureShader.(*OpenGLShader).UploadUniformMat4("u_ViewProjection", camera.GetViewProjectionMatrix())
	textureShader.(*OpenGLShader).UploadUniformInt("u_Texture", 0)

	quadVertices = quadVertices[0:0]
	quadIndexCount = 0

	textureSlotIndex = 1
}

func Renderer2DEndScene() {
	defer ggdebug.Stop(ggdebug.Start())

	quadVertexBuffer.SetData(quadVertices, uint32(unsafe.Sizeof(quadVertex{}))*uint32(len(quadVertices)))
	renderer2DFlush()
}

func Renderer2DDrawQuad(quad *Quad2D) {
	defer ggdebug.Stop(ggdebug.Start())

	var textureIndex uint32 = 0
	if quad.Texture != nil {
		for i := 0; i < int(textureSlotIndex); i++ {
			if textureSlots[i] == quad.Texture {
				textureIndex = uint32(i)
				break
			}
		}
		if textureIndex == 0 {
			textureIndex = textureSlotIndex
			textureSlots[textureIndex] = quad.Texture
			textureSlotIndex++
		}
	}

	quadVertices = append(quadVertices, quadVertex{
		Position: quad.Position.Vec3(quad.Z),
		Color:    *quad.Color,
		TexCoord: glm.Vec2{0, 0},
		TexIndex: float32(textureIndex),
	})

	quadVertices = append(quadVertices, quadVertex{
		Position: glm.Vec3{quad.Position.X() + quad.Size.X(), quad.Position.Y(), quad.Z},
		Color:    *quad.Color,
		TexCoord: glm.Vec2{1, 0},
		TexIndex: float32(textureIndex),
	})

	quadVertices = append(quadVertices, quadVertex{
		Position: glm.Vec3{quad.Position.X() + quad.Size.X(), quad.Position.Y() + quad.Size.Y(), quad.Z},
		Color:    *quad.Color,
		TexCoord: glm.Vec2{1, 1},
		TexIndex: float32(textureIndex),
	})

	quadVertices = append(quadVertices, quadVertex{
		Position: glm.Vec3{quad.Position.X(), quad.Position.Y() + quad.Size.Y(), quad.Z},
		Color:    *quad.Color,
		TexCoord: glm.Vec2{0, 1},
		TexIndex: float32(textureIndex),
	})

	quadIndexCount += 6

	/*if quad.Texture == nil || quad.TintTexture {
		textureShader.(*OpenGLShader).UploadUniformFloat4("u_Color", quad.Color)
	} else {
		textureShader.(*OpenGLShader).UploadUniformFloat4("u_Color", &glm.Vec4{1, 1, 1, 1})
	}

	if quad.Texture == nil {
		whiteTexture.Bind(0)
	} else {
		quad.Texture.Bind(0)
	}

	if quad.TilingFactor == 0 {
		textureShader.(*OpenGLShader).UploadUniformFloat("u_TilingFactor", 1)
	} else {
		textureShader.(*OpenGLShader).UploadUniformFloat("u_TilingFactor", quad.TilingFactor)
	}

	transform := glm.Translate3D(quad.Position[0], quad.Position[1], quad.Z)
	if quad.Rotation != 0 {
		rotation := glm.HomogRotate3DZ(glm.DegToRad(quad.Rotation))
		transform.Mul4With(&rotation)
	}
	scale := glm.Scale3D(quad.Size[0], quad.Size[1], 1)
	transform.Mul4With(&scale)

	textureShader.(*OpenGLShader).UploadUniformMat4("u_Transform", &transform)

	RenderCommandDrawIndexed(quadVertexArray)*/
}

func renderer2DFlush() {
	var i uint32
	for i = 0; i < textureSlotIndex; i++ {
		textureSlots[i].Bind(i)
	}

	RenderCommandDrawIndexed(quadVertexArray, quadIndexCount)
}
