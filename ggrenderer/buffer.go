package ggrenderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
)

//------------- Abstract -------------

type ShaderDataType int

const (
	ShaderDataTypeNone   ShaderDataType = 0
	ShaderDataTypeFloat  ShaderDataType = 1
	ShaderDataTypeFloat2 ShaderDataType = 2
	ShaderDataTypeFloat3 ShaderDataType = 3
	ShaderDataTypeFloat4 ShaderDataType = 4
	ShaderDataTypeMat3   ShaderDataType = 5
	ShaderDataTypeMat4   ShaderDataType = 6
	ShaderDataTypeInt    ShaderDataType = 7
	ShaderDataTypeInt2   ShaderDataType = 8
	ShaderDataTypeInt3   ShaderDataType = 9
	ShaderDataTypeInt4   ShaderDataType = 10
	ShaderDataTypeBool   ShaderDataType = 11
)

type BufferElement struct {
	DataType   ShaderDataType
	Name       string
	Normalized bool
	Offset     uint32
	Size       uint32
}

type BufferLayout struct {
	bufferElements []*BufferElement
	stride         uint32
}

type VertexBuffer interface {
	Delete()
	Bind()
	Unbind()
	SetLayout(layout *BufferLayout)
	GetLayout() (layout *BufferLayout)
	SetData(data interface{}, size uint32)
}

type IndexBuffer interface {
	Delete()
	Bind()
	Unbind()
	GetCount() (count uint32)
}

func ShaderDataTypeSize(dataType ShaderDataType) uint32 {
	switch dataType {
	case ShaderDataTypeNone:
		ggcore.CoreError("ShaderDataTypeNone not supported")
	case ShaderDataTypeFloat:
		return 4
	case ShaderDataTypeFloat2:
		return 2 * 4
	case ShaderDataTypeFloat3:
		return 3 * 4
	case ShaderDataTypeFloat4:
		return 4 * 4
	case ShaderDataTypeMat3:
		return 3 * 3 * 4
	case ShaderDataTypeMat4:
		return 4 * 4 * 4
	case ShaderDataTypeInt:
		return 4
	case ShaderDataTypeInt2:
		return 2 * 4
	case ShaderDataTypeInt3:
		return 3 * 4
	case ShaderDataTypeInt4:
		return 4 * 4
	case ShaderDataTypeBool:
		return 1
	}
	ggcore.CoreError("Unknown shader data type")
	return 0
}

func NewBufferElement(dataType ShaderDataType, name string, normalized bool) *BufferElement {
	element := BufferElement{
		DataType:   dataType,
		Name:       name,
		Normalized: normalized,
		Offset:     0,
		Size:       ShaderDataTypeSize(dataType),
	}
	return &element
}

func (element *BufferElement) GetComponentCount() uint32 {
	switch element.DataType {
	case ShaderDataTypeNone:
		ggcore.CoreError("ShaderDataTypeNone not supported")
	case ShaderDataTypeFloat:
		return 1
	case ShaderDataTypeFloat2:
		return 2
	case ShaderDataTypeFloat3:
		return 3
	case ShaderDataTypeFloat4:
		return 4
	case ShaderDataTypeMat3:
		return 3 * 3
	case ShaderDataTypeMat4:
		return 4 * 4
	case ShaderDataTypeInt:
		return 1
	case ShaderDataTypeInt2:
		return 2
	case ShaderDataTypeInt3:
		return 3
	case ShaderDataTypeInt4:
		return 4
	case ShaderDataTypeBool:
		return 1
	}
	ggcore.CoreError("Unknown shader data type")
	return 0
}

func NewBufferLayout(elements []*BufferElement) *BufferLayout {
	layout := BufferLayout{bufferElements: elements}
	layout.calculateOffsetsAndStride()
	return &layout
}

func (layout *BufferLayout) GetElements() *[]*BufferElement {
	return &layout.bufferElements
}

func (layout *BufferLayout) GetStride() uint32 {
	return layout.stride
}

func (layout *BufferLayout) calculateOffsetsAndStride() {
	var offset uint32
	layout.stride = 0
	for _, element := range layout.bufferElements {
		element.Offset = offset
		offset += element.Size
		layout.stride += element.Size
	}
}

func NewVertexBuffer(vertices []float32) VertexBuffer {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLVertexBuffer(vertices)
	}
	ggcore.CoreError("Unknown renderer API")
	return nil
}

func NewEmptyVertexBuffer(size uint32) VertexBuffer {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLEmptyVertexBuffer(size)
	}
	ggcore.CoreError("Unknown renderer API")
	return nil
}

func NewIndexBuffer(indices []uint32) IndexBuffer {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLIndexBuffer(indices)
	}
	ggcore.CoreError("unknown renderer API")
	return nil
}

//------------------------------------

//------------- Open GL --------------

type openGLVertexBuffer struct {
	rendererID uint32
	layout     *BufferLayout
}

func newOpenGLVertexBuffer(vertices []float32) *openGLVertexBuffer {
	defer ggdebug.Stop(ggdebug.Start())

	vertexBuffer := openGLVertexBuffer{}

	gl.CreateBuffers(1, &vertexBuffer.rendererID)
	vertexBuffer.Bind()
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	return &vertexBuffer
}

func newOpenGLEmptyVertexBuffer(size uint32) *openGLVertexBuffer {
	defer ggdebug.Stop(ggdebug.Start())

	vertexBuffer := openGLVertexBuffer{}

	gl.CreateBuffers(1, &vertexBuffer.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, int(size), nil, gl.DYNAMIC_DRAW)

	return &vertexBuffer
}

func (vertexBuffer *openGLVertexBuffer) Delete() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.DeleteBuffers(1, &vertexBuffer.rendererID)
}

func (vertexBuffer *openGLVertexBuffer) Bind() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer.rendererID)
}

func (vertexBuffer *openGLVertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (vertexBuffer *openGLVertexBuffer) SetLayout(layout *BufferLayout) {
	vertexBuffer.layout = layout
}

func (vertexBuffer *openGLVertexBuffer) GetLayout() *BufferLayout {
	return vertexBuffer.layout
}

func (vertexBuffer *openGLVertexBuffer) SetData(data interface{}, size uint32) {
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer.rendererID)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, int(size), gl.Ptr(data))
}

type openGLIndexBuffer struct {
	rendererID uint32
	count      uint32
}

func newOpenGLIndexBuffer(indices []uint32) *openGLIndexBuffer {
	defer ggdebug.Stop(ggdebug.Start())

	indexBuffer := openGLIndexBuffer{}

	gl.CreateBuffers(1, &indexBuffer.rendererID)
	indexBuffer.Bind()
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	indexBuffer.count = uint32(len(indices))

	return &indexBuffer
}

func (indexBuffer *openGLIndexBuffer) Delete() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.DeleteBuffers(1, &indexBuffer.rendererID)
}

func (indexBuffer *openGLIndexBuffer) Bind() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer.rendererID)
}

func (indexBuffer *openGLIndexBuffer) Unbind() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (indexBuffer *openGLIndexBuffer) GetCount() uint32 {
	return indexBuffer.count
}

//------------------------------------
