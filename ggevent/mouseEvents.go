package ggevent

import (
	"fmt"
)

type MouseMovedEvent struct {
	handled bool
	mouseX  float32
	mouseY  float32
}

type MouseScrolledEvent struct {
	handled bool
	offsetX float32
	offsetY float32
}

type MouseButtonPressedEvent struct {
	handled     bool
	mouseButton MouseButtonCode
}

type MouseButtonReleasedEvent struct {
	handled     bool
	mouseButton MouseButtonCode
}

func NewMouseMovedEvent(mouseX, mouseY float32) *MouseMovedEvent {
	return &MouseMovedEvent{handled: false, mouseX: mouseX, mouseY: mouseY}
}

func (event *MouseMovedEvent) GetEventType() EventType { return EventTypeMouseMoved }
func (event *MouseMovedEvent) GetName() string         { return "MouseMoved" }
func (event *MouseMovedEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryMouse
}
func (event *MouseMovedEvent) String() string {
	return fmt.Sprintf("%v: %v, %v", event.GetName(), event.MouseX(), event.MouseY())
}
func (event *MouseMovedEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *MouseMovedEvent) IsHandled() bool { return event.handled }
func (event *MouseMovedEvent) SetHandled()     { event.handled = true }
func (event *MouseMovedEvent) MouseX() float32 { return event.mouseX }
func (event *MouseMovedEvent) MouseY() float32 { return event.mouseY }

func NewMouseScrolledEvent(offsetX, offsetY float32) *MouseScrolledEvent {
	return &MouseScrolledEvent{handled: false, offsetX: offsetX, offsetY: offsetY}
}

func (event *MouseScrolledEvent) GetEventType() EventType { return EventTypeMouseScrolled }
func (event *MouseScrolledEvent) GetName() string         { return "MouseScrolled" }
func (event *MouseScrolledEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryMouse
}
func (event *MouseScrolledEvent) String() string {
	return fmt.Sprintf("%v: %v, %v", event.GetName(), event.OffsetX(), event.OffsetY())
}
func (event *MouseScrolledEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *MouseScrolledEvent) IsHandled() bool  { return event.handled }
func (event *MouseScrolledEvent) SetHandled()      { event.handled = true }
func (event *MouseScrolledEvent) OffsetX() float32 { return event.offsetX }
func (event *MouseScrolledEvent) OffsetY() float32 { return event.offsetY }

func NewMouseButtonPressedEvent(mouseButton MouseButtonCode) *MouseButtonPressedEvent {
	return &MouseButtonPressedEvent{handled: false, mouseButton: mouseButton}
}

func (event *MouseButtonPressedEvent) GetEventType() EventType { return EventTypeMouseButtonPressed }
func (event *MouseButtonPressedEvent) GetName() string         { return "MouseButtonPressed" }
func (event *MouseButtonPressedEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryMouse | EventCategoryMouseButton
}
func (event *MouseButtonPressedEvent) String() string {
	return fmt.Sprintf("%v: %v", event.GetName(), event.MouseButton())
}
func (event *MouseButtonPressedEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *MouseButtonPressedEvent) IsHandled() bool { return event.handled }
func (event *MouseButtonPressedEvent) SetHandled()     { event.handled = true }
func (event *MouseButtonPressedEvent) MouseButton() MouseButtonCode {
	return event.mouseButton
}

func NewMouseButtonReleasedEvent(mouseButton MouseButtonCode) *MouseButtonReleasedEvent {
	return &MouseButtonReleasedEvent{handled: false, mouseButton: mouseButton}
}

func (event *MouseButtonReleasedEvent) GetEventType() EventType { return EventTypeMouseButtonReleased }
func (event *MouseButtonReleasedEvent) GetName() string         { return "MouseButtonReleased" }
func (event *MouseButtonReleasedEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryMouse | EventCategoryMouseButton
}
func (event *MouseButtonReleasedEvent) String() string {
	return fmt.Sprintf("%v: %v", event.GetName(), event.MouseButton())
}
func (event *MouseButtonReleasedEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *MouseButtonReleasedEvent) IsHandled() bool { return event.handled }
func (event *MouseButtonReleasedEvent) SetHandled()     { event.handled = true }
func (event *MouseButtonReleasedEvent) MouseButton() MouseButtonCode {
	return event.mouseButton
}
