package ggrenderer

import (
	"image"
	"image/draw"
	"os"
	"unsafe"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
)

//------------- Abstract -------------

type Texture interface {
	Delete()
	GetWidth() (width uint32)
	GetHeight() (height uint32)
	SetData(data unsafe.Pointer)
	Bind(slot uint32)
}

func NewTexture2DFromPath(path string) Texture {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLTexture2DFromPath(path)
	}
	ggcore.CoreError("unknown renderer API")
	return nil
}

func NewTexture2DEmpty(width, height uint32) Texture {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLTexture2DEmpty(width, height)
	}
	ggcore.CoreError("unknown renderer API")
	return nil
}

//------------------------------------

//------------- Open GL --------------

type openGLTexture2D struct {
	path          string
	width, height uint32
	rendererID    uint32
}

func newOpenGLTexture2DFromPath(path string) *openGLTexture2D {
	defer ggdebug.Stop(ggdebug.Start())

	texture := openGLTexture2D{path: path}

	reader, err := os.Open(path)
	ggcore.CoreCheckError(err)
	defer reader.Close()

	img, _, err := image.Decode(reader)
	ggcore.CoreCheckError(err)
	bounds := img.Bounds()
	texture.width = uint32(bounds.Size().X)
	texture.height = uint32(bounds.Size().Y)

	imgData := image.NewRGBA(img.Bounds())
	if imgData.Stride != imgData.Rect.Size().X*4 {
		ggcore.CoreError("Unsupported stride")
	}
	draw.Draw(imgData, imgData.Bounds(), imaging.FlipV(img), image.Point{0, 0}, draw.Src)

	gl.CreateTextures(gl.TEXTURE_2D, 1, &texture.rendererID)
	gl.TextureStorage2D(texture.rendererID, 1, gl.RGBA8, int32(texture.width), int32(texture.height))

	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TextureSubImage2D(texture.rendererID, 0, 0, 0, int32(texture.width), int32(texture.height), gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(imgData.Pix))

	return &texture
}

func newOpenGLTexture2DEmpty(width, height uint32) *openGLTexture2D {
	defer ggdebug.Stop(ggdebug.Start())

	texture := openGLTexture2D{
		width:  width,
		height: height,
	}

	gl.CreateTextures(gl.TEXTURE_2D, 1, &texture.rendererID)
	gl.TextureStorage2D(texture.rendererID, 1, gl.RGBA8, int32(texture.width), int32(texture.height))

	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TextureParameteri(texture.rendererID, gl.TEXTURE_WRAP_T, gl.REPEAT)

	return &texture
}

func (texture *openGLTexture2D) Delete() {
	defer ggdebug.Stop(ggdebug.Start())

	gl.DeleteTextures(1, &texture.rendererID)
}

func (texture *openGLTexture2D) GetWidth() uint32 { return texture.width }

func (texture *openGLTexture2D) GetHeight() uint32 { return texture.height }

func (texture *openGLTexture2D) SetData(data unsafe.Pointer) {
	defer ggdebug.Stop(ggdebug.Start())

	gl.TextureSubImage2D(texture.rendererID, 0, 0, 0, int32(texture.width), int32(texture.height), gl.RGBA, gl.UNSIGNED_BYTE, data)
}

func (texture *openGLTexture2D) Bind(slot uint32) {
	defer ggdebug.Stop(ggdebug.Start())

	gl.BindTextureUnit(slot, texture.rendererID)
}

//------------------------------------
