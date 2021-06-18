package ggcore

import (
	"unsafe"

	"github.com/oyberntzen/gogame/ggevent"
)

type WindowProps struct {
	Title         string
	Width, Height uint
}

type EventCallbackFunc func(ggevent.Event)

type Window interface {
	OnUpdate()
	Width() uint
	Height() uint
	SetEventCallback(EventCallbackFunc)
	SetVSync(bool)
	VSync() bool
	GetNativeWindow() unsafe.Pointer
	Delete()
}
