package ggrenderer

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/EngoEngine/glm"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/oyberntzen/gogame/ggcore"
)

//------------- Abstract -------------

type Shader interface {
	Delete()
	Bind()
	Unbind()
	GetName() (name string)
}

func NewShaderFromSrc(name, vertexSrc, fragmentSrc string) Shader {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLShaderFromSrc(name, vertexSrc, fragmentSrc)
	}
	ggcore.CoreError("Unknown renderer API")
	return nil
}

func NewShaderFromPath(path string) Shader {
	switch CurrentAPI() {
	case RendererAPINone:
		ggcore.CoreError("RendererAPINone is not supported")
	case RendererAPIOpenGL:
		return newOpenGLShaderFromPath(path)
	}
	ggcore.CoreError("Unknown renderer API")
	return nil
}

type ShaderLibrary struct {
	shaders map[string]Shader
}

func NewShaderLibrary() *ShaderLibrary {
	library := ShaderLibrary{shaders: make(map[string]Shader)}
	return &library
}

func (library *ShaderLibrary) Add(shader Shader) {
	name := shader.GetName()
	if _, ok := library.shaders[name]; ok {
		ggcore.CoreError("Shader %v already exists")
	}
	library.shaders[name] = shader
}

func (library *ShaderLibrary) LoadFromPath(path string) Shader {
	shader := newOpenGLShaderFromPath(path)
	library.Add(shader)
	return shader
}

func (library *ShaderLibrary) Get(name string) Shader {
	if shader, ok := library.shaders[name]; ok {
		return shader
	}
	ggcore.CoreError("Shader %v not found", name)
	return nil
}

//------------------------------------

//------------- Open GL --------------

//should become private
type OpenGLShader struct {
	rendererID uint32
	name       string
}

const typeToken = "#type"

func newOpenGLShaderFromSrc(name, vertexSrc, fragmentSrc string) *OpenGLShader {
	shaderSources := make(map[uint32]string)

	shaderSources[gl.VERTEX_SHADER] = vertexSrc
	shaderSources[gl.FRAGMENT_SHADER] = fragmentSrc

	shader := OpenGLShader{name: name}
	shader.compile(shaderSources)
	return &shader
}

func newOpenGLShaderFromPath(path string) *OpenGLShader {
	file, err := os.Open(path)
	ggcore.CoreCheckError(err)
	scanner := bufio.NewScanner(file)

	shaderSources := make(map[uint32]string)
	var currentType uint32 = 0
	for scanner.Scan() {
		text := scanner.Text()
		newType := false
		if split := strings.Split(text, " "); len(split) == 2 {
			if split[0] == typeToken {
				currentType = shaderTypeFromString(split[1])
				newType = true
			}
		}
		if !newType && currentType != 0 {
			shaderSources[currentType] += text + "\n"
		}
	}
	filename := filepath.Base(path)
	shader := OpenGLShader{name: filename[0 : len(filename)-len(filepath.Ext(filename))]}
	ggcore.CoreInfo("Shader name: %v", shader.name)
	shader.compile(shaderSources)
	return &shader
}

func shaderTypeFromString(shaderType string) uint32 {
	if shaderType == "vertex" {
		return gl.VERTEX_SHADER
	}
	if shaderType == "fragment" || shaderType == "pixel" {
		return gl.FRAGMENT_SHADER
	}
	ggcore.CoreError("Unsupported shader type")
	return 0
}

func (shader *OpenGLShader) Delete() {
	gl.DeleteProgram(shader.rendererID)
}

func (shader *OpenGLShader) Bind() {
	gl.UseProgram(shader.rendererID)
}

func (shader *OpenGLShader) Unbind() {
	gl.UseProgram(0)
}

func (shader *OpenGLShader) GetName() string {
	return shader.name
}

func (shader *OpenGLShader) UploadUniformInt(name string, value int32) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.Uniform1i(location, value)
}

func (shader *OpenGLShader) UploadUniformFloat(name string, value float32) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.Uniform1f(location, value)
}

func (shader *OpenGLShader) UploadUniformFloat2(name string, value *glm.Vec2) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.Uniform2f(location, value[0], value[1])
}

func (shader *OpenGLShader) UploadUniformFloat3(name string, value *glm.Vec3) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.Uniform3f(location, value[0], value[1], value[2])
}

func (shader *OpenGLShader) UploadUniformFloat4(name string, value *glm.Vec4) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.Uniform4f(location, value[0], value[1], value[2], value[3])
}

func (shader *OpenGLShader) UploadUniformMat2(name string, matrix *glm.Mat2) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.UniformMatrix2fv(location, 1, false, &matrix[0])
}

func (shader *OpenGLShader) UploadUniformMat3(name string, matrix *glm.Mat3) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.UniformMatrix3fv(location, 1, false, &matrix[0])
}

func (shader *OpenGLShader) UploadUniformMat4(name string, matrix *glm.Mat4) {
	location := gl.GetUniformLocation(shader.rendererID, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}

func (shader *OpenGLShader) compile(shaderSources map[uint32]string) {
	shader.rendererID = gl.CreateProgram()
	shaderIDs := []uint32{}
	for shaderType, source := range shaderSources {
		currentShader := gl.CreateShader(shaderType)
		csource, free := gl.Strs(source + "\x00")
		gl.ShaderSource(currentShader, 1, csource, nil)
		free()
		gl.CompileShader(currentShader)

		var result int32
		gl.GetShaderiv(currentShader, gl.COMPILE_STATUS, &result)
		if result == gl.FALSE {
			var length int32
			gl.GetShaderiv(currentShader, gl.INFO_LOG_LENGTH, &length)

			message := strings.Repeat("\x00", int(length+1))
			gl.GetShaderInfoLog(currentShader, length, &length, gl.Str(message))
			gl.DeleteShader(currentShader)

			ggcore.CoreError("Shader compilation failure: %v", message)
		}
		gl.AttachShader(shader.rendererID, currentShader)
		shaderIDs = append(shaderIDs, currentShader)
	}

	gl.LinkProgram(shader.rendererID)

	var result int32
	gl.GetProgramiv(shader.rendererID, gl.LINK_STATUS, &result)
	if result == gl.FALSE {
		var length int32
		gl.GetProgramiv(shader.rendererID, gl.INFO_LOG_LENGTH, &length)

		message := strings.Repeat("\x00", int(length+1))
		gl.GetProgramInfoLog(shader.rendererID, length, &length, gl.Str(message))
		gl.DeleteProgram(shader.rendererID)

		for _, currentShader := range shaderIDs {
			gl.DeleteShader(currentShader)
		}
		ggcore.CoreError("Shader program linking failure: %v", message)
	}

	for _, currentShader := range shaderIDs {
		gl.DetachShader(shader.rendererID, currentShader)
	}
}

//------------------------------------
