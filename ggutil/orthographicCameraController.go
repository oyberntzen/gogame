package ggutil

import (
	"github.com/EngoEngine/glm"
	"github.com/EngoEngine/math"
	"github.com/oyberntzen/gogame/ggconstants"
	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggevent"
	"github.com/oyberntzen/gogame/ggrenderer"
)

type OrthographicCameraController struct {
	camera                                                       *ggrenderer.OrthographicCamera
	aspectRatio, zoomLevel                                       float32
	rotation                                                     bool
	cameraPosition                                               glm.Vec3
	cameraRotation                                               float32
	cameraTranslationSpeed, cameraRotationSpeed, cameraZoomSpeed float32
}

func NewOrthographicCameraController(aspectRatio float32, rotation bool) *OrthographicCameraController {
	zoomLevel := float32(1)
	controller := OrthographicCameraController{
		camera:                 ggrenderer.NewOrthographicCamera(-aspectRatio*zoomLevel, aspectRatio*zoomLevel, -zoomLevel, zoomLevel),
		aspectRatio:            aspectRatio,
		zoomLevel:              zoomLevel,
		rotation:               rotation,
		cameraTranslationSpeed: 5,
		cameraRotationSpeed:    180,
		cameraZoomSpeed:        0.25,
	}

	return &controller
}

func (controller *OrthographicCameraController) OnUpdate(ts ggcore.Timestep) {
	if ggcore.IsKeyPressed(ggconstants.KeyA) {
		controller.cameraPosition[0] -= controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if ggcore.IsKeyPressed(ggconstants.KeyD) {
		controller.cameraPosition[0] += controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if ggcore.IsKeyPressed(ggconstants.KeyW) {
		controller.cameraPosition[1] += controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if ggcore.IsKeyPressed(ggconstants.KeyS) {
		controller.cameraPosition[1] -= controller.cameraTranslationSpeed * ts.GetSeconds()
	}
	if controller.rotation {
		if ggcore.IsKeyPressed(ggconstants.KeyQ) {
			controller.cameraRotation += controller.cameraRotationSpeed * ts.GetSeconds()
		}
		if ggcore.IsKeyPressed(ggconstants.KeyE) {
			controller.cameraRotation -= controller.cameraRotationSpeed * ts.GetSeconds()
		}
		controller.camera.SetRotation(controller.cameraRotation)
	}

	controller.camera.SetPosition(&controller.cameraPosition)

	controller.cameraTranslationSpeed = controller.zoomLevel
}

func (controller *OrthographicCameraController) OnEvent(event ggevent.Event) {
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
	controller.camera.SetProjection(-controller.aspectRatio*controller.zoomLevel, controller.aspectRatio*controller.zoomLevel, -controller.zoomLevel, controller.zoomLevel)
	return false
}

func (controller *OrthographicCameraController) onWindowResize(event *ggevent.WindowResizeEvent) bool {
	controller.aspectRatio = float32(event.Width()) / float32(event.Height())
	controller.camera.SetProjection(-controller.aspectRatio*controller.zoomLevel, controller.aspectRatio*controller.zoomLevel, -controller.zoomLevel, controller.zoomLevel)
	return false
}