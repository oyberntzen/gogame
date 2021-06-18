package ggevent

type EventType int

const (
	EventTypeNone                EventType = 0
	EventTypeWindowClose         EventType = 1
	EventTypeWindowResize        EventType = 2
	EventTypeWindowFocus         EventType = 3
	EventTypeWindowLostFocus     EventType = 4
	EventTypeWindowMoved         EventType = 5
	EventTypeAppTick             EventType = 6
	EventTypeAppUpdate           EventType = 7
	EventTypeAppRender           EventType = 8
	EventTypeKeyPressed          EventType = 9
	EventTypeKeyReleased         EventType = 10
	EventTypeKeyTyped            EventType = 11
	EventTypeMouseButtonPressed  EventType = 12
	EventTypeMouseButtonReleased EventType = 13
	EventTypeMouseMoved          EventType = 14
	EventTypeMouseScrolled       EventType = 15
)

type EventCategory int

const (
	EventCategoryNone        = 0
	EventCategoryApplication = 1 << 0
	EventCategoryInput       = 1 << 1
	EventCategoryKeyboard    = 1 << 2
	EventCategoryMouse       = 1 << 3
	EventCategoryMouseButton = 1 << 4
)

type Event interface {
	GetEventType() EventType
	GetName() string
	GetCategoryFlags() EventCategory
	String() string
	IsInCategory(EventCategory) bool
	IsHandled() bool
	SetHandled()
}

type EventDispatcher struct {
	Event Event
}

func (dispatcher *EventDispatcher) Dispatch(eventType EventType, function func(Event) bool) bool {
	if dispatcher.Event.GetEventType() == eventType {
		if function(dispatcher.Event) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchWindowClose(function func(*WindowCloseEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeWindowClose {
		if function(dispatcher.Event.(*WindowCloseEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchWindowResize(function func(*WindowResizeEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeWindowResize {
		if function(dispatcher.Event.(*WindowResizeEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchAppTick(function func(*AppTickEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeAppTick {
		if function(dispatcher.Event.(*AppTickEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchAppUpdate(function func(*AppUpdateEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeAppUpdate {
		if function(dispatcher.Event.(*AppUpdateEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchAppRender(function func(*AppRenderEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeAppRender {
		if function(dispatcher.Event.(*AppRenderEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchKeyPressed(function func(*KeyPressedEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeKeyPressed {
		if function(dispatcher.Event.(*KeyPressedEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchKeyReleased(function func(*KeyReleasedEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeKeyReleased {
		if function(dispatcher.Event.(*KeyReleasedEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchKeyTyped(function func(*KeyTypedEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeKeyTyped {
		if function(dispatcher.Event.(*KeyTypedEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchMouseButtonPressed(function func(*MouseButtonPressedEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeMouseButtonPressed {
		if function(dispatcher.Event.(*MouseButtonPressedEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchMouseButtonReleased(function func(*MouseButtonReleasedEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeMouseButtonReleased {
		if function(dispatcher.Event.(*MouseButtonReleasedEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchMouseMoved(function func(*MouseMovedEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeMouseMoved {
		if function(dispatcher.Event.(*MouseMovedEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}

func (dispatcher *EventDispatcher) DispatchMouseScrolled(function func(*MouseScrolledEvent) bool) bool {
	if dispatcher.Event.GetEventType() == EventTypeMouseScrolled {
		if function(dispatcher.Event.(*MouseScrolledEvent)) {
			dispatcher.Event.SetHandled()
		}
		return true
	}
	return false
}
