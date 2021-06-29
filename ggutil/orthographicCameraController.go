package ggutil

import (
	"github.com/EngoEngine/glm"
	"github.com/EngoEngine/math"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
	"github.com/oyberntzen/gogame/ggevent"
	"github.com/oyberntzen/gogame/ggrenderer"
)

type OrthographicCameraBounds struct {
	Left, Right, Bottom, Top float32
}

func (bounds *OrthographicCameraBounds) Width() float32  { return bounds.Right - bounds.Left }
func (bounds *OrthographicCameraBounds) Height() float32 { return bounds.Top - bounds.Bottom }

type OrthographicCameraController struct {
	camera                                                       *ggrenderer.OrthographicCamera
	aspectRatio, zoomLevel                                       float32
	rotation                                                     bool
	cameraPosition                                               glm.Vec3
	cameraRotation                                               float32
	cameraTranslationSpeed, cameraRotationSpeed, cameraZoomSpeed float32

	bounds OrthographicCameraBounds
}

func NewOrthographicCameraController(aspectRatio float32, rotation bool) *OrthographicCameraController {
	defer ggdebug.Stop(ggdebug.Start())

	zoomLevel := float32(1)
	controller := OrthographicCameraController{
		camera:                 ggrenderer.NewOrthographicCamera(-aspectRatio*zoomLevel, aspectRatio*zoomLevel, -zoomLevel, zoomLevel),
		aspectRatio:            aspectRatio,
		zoomLevel:              zoomLevel,
		rotation:               rotation,
		cameraTranslationSpeed: 5,
		cameraRotationSpeed:    180,
		cameraZoomSpeed:        0.25,

		bounds: OrthographicCameraBounds{
			Left:   -aspectRatio * zoomLevel,
			Right:  aspectRatio * zoomLevel,
			Bottom: -zoomLevel,
			Top:    zoomLevel,
		},
	}

	return &controller
}

func (controller *OrthographicCameraController) OnUpdate(ts ggcore.Timestep) {
	defer ggdebug.Stop(ggdebug.Start())

	if ggcore.IsKeyPressed(ggevent.KeyA) {
		controller.cameraPosition[0] -= math.Cos(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
		controller.cameraPosition[1] -= math.Sin(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if ggcore.IsKeyPressed(ggevent.KeyD) {
		controller.cameraPosition[0] += math.Cos(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
		controller.cameraPosition[1] += math.Sin(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if ggcore.IsKeyPressed(ggevent.KeyW) {
		controller.cameraPosition[0] -= math.Sin(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
		controller.cameraPosition[1] += math.Cos(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if ggcore.IsKeyPressed(ggevent.KeyS) {
		controller.cameraPosition[0] += math.Sin(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
		controller.cameraPosition[1] -= math.Cos(glm.DegToRad(controller.cameraRotation)) * controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if controller.rotation {
		if ggcore.IsKeyPressed(ggevent.KeyQ) {
			controller.cameraRotation += controller.cameraRotationSpeed * ts.GetSeconds()
		}
		if ggcore.IsKeyPressed(ggevent.KeyE) {
			controller.cameraRotation -= controller.cameraRotationSpeed * ts.GetSeconds()
		}
		controller.camera.SetRotation(controller.cameraRotation)
	}

	controller.camera.SetPosition(&controller.cameraPosition)
}

func (controller *OrthographicCameraController) OnEvent(event ggevent.Event) {
	defer ggdebug.Stop(ggdebug.Start())

	dispatcher := ggevent.EventDispatcher{Event: event}
	dispatcher.DispatchMouseScrolled(controller.onMouseScrolled)
	dispatcher.DispatchWindowResize(controller.onWindowResize)
}

func (controller *OrthographicCameraController) GetCamera() *ggrenderer.OrthographicCamera {
	return controller.camera
}

func (controller *OrthographicCameraController) onMouseScrolled(event *ggevent.MouseScrolledEvent) bool {
	controller.zoomLevel -= event.OffsetY() * controller.cameraZoomSpeed
	controller.zoomLevel = math.Max(controller.zoomLevel, controller.cameraZoomSpeed)

	controller.bounds = OrthographicCameraBounds{
		Left:   -controller.aspectRatio * controller.zoomLevel,
		Right:  controller.aspectRatio * controller.zoomLevel,
		Bottom: -controller.zoomLevel,
		Top:    controller.zoomLevel,
	}
	controller.camera.SetProjection(controller.bounds.Left, controller.bounds.Right, controller.bounds.Bottom, controller.bounds.Top)
	return false
}

func (controller *OrthographicCameraController) onWindowResize(event *ggevent.WindowResizeEvent) bool {
	controller.aspectRatio = float32(event.Width()) / float32(event.Height())

	controller.bounds = OrthographicCameraBounds{
		Left:   -controller.aspectRatio * controller.zoomLevel,
		Right:  controller.aspectRatio * controller.zoomLevel,
		Bottom: -controller.zoomLevel,
		Top:    controller.zoomLevel,
	}
	controller.camera.SetProjection(controller.bounds.Left, controller.bounds.Right, controller.bounds.Bottom, controller.bounds.Top)
	return false
}

func (controller *OrthographicCameraController) Bounds() *OrthographicCameraBounds {
	return &controller.bounds
}
