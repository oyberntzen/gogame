package ggevent

import "fmt"

type WindowResizeEvent struct {
	handled bool
	width   int
	height  int
}

type WindowCloseEvent struct {
	handled bool
}

type AppTickEvent struct {
	handled bool
}

type AppUpdateEvent struct {
	handled bool
}

type AppRenderEvent struct {
	handled bool
}

func NewWindowResizeEvent(width, height int) *WindowResizeEvent {
	return &WindowResizeEvent{handled: false, width: width, height: height}
}

func (event *WindowResizeEvent) GetEventType() EventType { return EventTypeWindowResize }
func (event *WindowResizeEvent) GetName() string         { return "WindowResize" }
func (event *WindowResizeEvent) GetCategoryFlags() EventCategory {
	return EventCategoryApplication
}
func (event *WindowResizeEvent) String() string {
	return fmt.Sprintf("%v: %v, %v", event.GetName(), event.Width(), event.Height())
}
func (event *WindowResizeEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *WindowResizeEvent) IsHandled() bool { return event.handled }
func (event *WindowResizeEvent) SetHandled()     { event.handled = true }
func (event *WindowResizeEvent) Width() int      { return event.width }
func (event *WindowResizeEvent) Height() int     { return event.height }

func NewWindowCloseEvent() *WindowCloseEvent {
	return &WindowCloseEvent{handled: false}
}

func (event *WindowCloseEvent) GetEventType() EventType { return EventTypeWindowClose }
func (event *WindowCloseEvent) GetName() string         { return "WindowClose" }
func (event *WindowCloseEvent) GetCategoryFlags() EventCategory {
	return EventCategoryApplication
}
func (event *WindowCloseEvent) String() string {
	return event.GetName()
}
func (event *WindowCloseEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *WindowCloseEvent) IsHandled() bool { return event.handled }
func (event *WindowCloseEvent) SetHandled()     { event.handled = true }

func NewAppTickEvent() *AppTickEvent {
	return &AppTickEvent{handled: false}
}

func (event *AppTickEvent) GetEventType() EventType { return EventTypeAppTick }
func (event *AppTickEvent) GetName() string         { return "AppTick" }
func (event *AppTickEvent) GetCategoryFlags() EventCategory {
	return EventCategoryApplication
}
func (event *AppTickEvent) String() string {
	return event.GetName()
}
func (event *AppTickEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *AppTickEvent) IsHandled() bool { return event.handled }
func (event *AppTickEvent) SetHandled()     { event.handled = true }

func NewAppUpdateEvent() *AppUpdateEvent {
	return &AppUpdateEvent{handled: false}
}

func (event *AppUpdateEvent) GetEventType() EventType { return EventTypeAppUpdate }
func (event *AppUpdateEvent) GetName() string         { return "AppUpdate" }
func (event *AppUpdateEvent) GetCategoryFlags() EventCategory {
	return EventCategoryApplication
}
func (event *AppUpdateEvent) String() string {
	return event.GetName()
}
func (event *AppUpdateEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *AppUpdateEvent) IsHandled() bool { return event.handled }
func (event *AppUpdateEvent) SetHandled()     { event.handled = true }

func NewAppRenderEvent() WindowCloseEvent {
	return WindowCloseEvent{handled: false}
}

func (event *AppRenderEvent) GetEventType() EventType { return EventTypeAppRender }
func (event *AppRenderEvent) GetName() string         { return "AppRender" }
func (event *AppRenderEvent) GetCategoryFlags() EventCategory {
	return EventCategoryApplication
}
func (event *AppRenderEvent) String() string {
	return event.GetName()
}
func (event *AppRenderEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *AppRenderEvent) IsHandled() bool { return event.handled }
func (event *AppRenderEvent) SetHandled()     { event.handled = true }
