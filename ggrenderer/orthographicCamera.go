package ggrenderer

import (
	"github.com/EngoEngine/glm"
)

type OrthographicCamera struct {
	projectionMatrix     glm.Mat4
	viewMatrix           glm.Mat4
	viewProjectionMatrix glm.Mat4

	position glm.Vec3
	rotation float32
}

func NewOrthographicCamera(left, right, bottom, top float32) *OrthographicCamera {
	camera := OrthographicCamera{
		projectionMatrix: glm.Ortho2D(left, right, bottom, top),
		viewMatrix:       glm.Ident4(),
	}
	camera.viewProjectionMatrix = camera.projectionMatrix.Mul4(&camera.viewMatrix)
	return &camera
}

func (camera *OrthographicCamera) SetProjection(left, right, bottom, top float32) {
	camera.projectionMatrix = glm.Ortho2D(left, right, bottom, top)
	camera.viewProjectionMatrix = camera.projectionMatrix.Mul4(&camera.viewMatrix)
}

func (camera *OrthographicCamera) SetPosition(position *glm.Vec3) {
	camera.position = *position
	camera.recalculateViewMatrix()
}
func (camera *OrthographicCamera) GetPosition() *glm.Vec3 { return &camera.position }

func (camera *OrthographicCamera) SetRotation(rotation float32) {
	camera.rotation = glm.DegToRad(rotation)
	camera.recalculateViewMatrix()
}
func (camera *OrthographicCamera) GetRotation() float32 { return glm.RadToDeg(camera.rotation) }

func (camera *OrthographicCamera) GetProjectionMatrix() *glm.Mat4 { return &camera.projectionMatrix }
func (camera *OrthographicCamera) GetViewMatrix() *glm.Mat4       { return &camera.viewMatrix }
func (camera *OrthographicCamera) GetViewProjectionMatrix() *glm.Mat4 {
	return &camera.viewProjectionMatrix
}

func (camera *OrthographicCamera) recalculateViewMatrix() {
	position := glm.Translate3D(camera.position.X(), camera.position.Y(), camera.position.Z())
	rotation := glm.HomogRotate3DZ(camera.rotation)
	transform := position.Mul4(&rotation)
	camera.viewMatrix = transform.Inverse()
	camera.viewProjectionMatrix = camera.projectionMatrix.Mul4(&camera.viewMatrix)
}
