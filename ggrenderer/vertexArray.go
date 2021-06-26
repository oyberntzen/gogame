package ggrenderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
)

//------------- Abstract -------------

type VertexArray interface {
	Delete()
	Bind()
	Unbind()
	AddVertexBuffer(vertexBuffer VertexBuffer)
	SetIndexBuffer(indexBuffer IndexBuffer)
	GetVertexBuffers() (vertexBuffers []VertexBuffer)
	GetIndexBuffer() (indexBuffer IndexBuffer)
}

func NewVertexArray() VertexArray {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLVertexArray()
	}
	ggcore.CoreError("unknown renderer API")
	return nil
}

//------------------------------------

//------------- Open GL --------------

type openGLVertexArray struct {
	rendererID    uint32
	vertexBuffers []VertexBuffer
	indexBuffer   IndexBuffer
}

func ShaderDataTypeToOpenGLType(dataType ShaderDataType) uint32 {
	switch dataType {
	case ShaderDataTypeNone:
		panic("ShaderDataTypeNone not supported")
	case ShaderDataTypeFloat:
		return gl.FLOAT
	case ShaderDataTypeFloat2:
		return gl.FLOAT
	case ShaderDataTypeFloat3:
		return gl.FLOAT
	case ShaderDataTypeFloat4:
		return gl.FLOAT
	case ShaderDataTypeMat3:
		return gl.FLOAT
	case ShaderDataTypeMat4:
		return gl.FLOAT
	case ShaderDataTypeInt:
		return gl.INT
	case ShaderDataTypeInt2:
		return gl.INT
	case ShaderDataTypeInt3:
		return gl.INT
	case ShaderDataTypeInt4:
		return gl.INT
	case ShaderDataTypeBool:
		return gl.BOOL
	}
	panic("unknown shader data type")
}

func newOpenGLVertexArray() *openGLVertexArray {
	defer ggdebug.Stop(ggdebug.Start())

	vertexArray := openGLVertexArray{}
	gl.CreateVertexArrays(1, &vertexArray.rendererID)
	return &vertexArray
}

func (vertexArray *openGLVertexArray) Delete() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.DeleteVertexArrays(1, &vertexArray.rendererID)
}

func (vertexArray *openGLVertexArray) Bind() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.BindVertexArray(vertexArray.rendererID)
}

func (vertexArray *openGLVertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (vertexArray *openGLVertexArray) AddVertexBuffer(vertexBuffer VertexBuffer) {
	defer ggdebug.Stop(ggdebug.Start())

	layout := vertexBuffer.GetLayout()
	if len(*layout.GetElements()) == 0 {
		ggcore.CoreError("Vertex buffer layout is empty")
	}

	vertexArray.Bind()
	vertexBuffer.Bind()

	for index, element := range *layout.GetElements() {
		gl.EnableVertexAttribArray(uint32(index))
		gl.VertexAttribPointerWithOffset(
			uint32(index),
			int32(element.GetComponentCount()),
			ShaderDataTypeToOpenGLType(element.DataType),
			element.Normalized,
			int32(layout.GetStride()),
			uintptr(element.Offset),
		)
	}

	vertexArray.vertexBuffers = append(vertexArray.vertexBuffers, vertexBuffer)
}

func (vertexArray *openGLVertexArray) SetIndexBuffer(indexBuffer IndexBuffer) {
	vertexArray.Bind()
	indexBuffer.Bind()

	vertexArray.indexBuffer = indexBuffer
}

func (vertexArray *openGLVertexArray) GetVertexBuffers() []VertexBuffer {
	return vertexArray.vertexBuffers
}

func (vertexArray *openGLVertexArray) GetIndexBuffer() IndexBuffer {
	return vertexArray.indexBuffer
}

//------------------------------------
