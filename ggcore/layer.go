package ggcore

import (
	"github.com/oyberntzen/gogame/ggevent"
)

type Layer interface {
	OnAttach()
	OnDetach()
	OnUpdate(Timestep)
	OnImGuiRender()
	OnEvent(ggevent.Event)
	GetName() string
}
