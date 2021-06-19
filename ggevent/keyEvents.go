package ggevent

import (
	"fmt"
)

type KeyPressedEvent struct {
	handled     bool
	keyCode     KeyCode
	repeatCount int
}

type KeyReleasedEvent struct {
	handled bool
	keyCode KeyCode
}

type KeyTypedEvent struct {
	handled   bool
	character rune
}

func NewKeyPressedEvent(keyCode KeyCode, repeatCount int) *KeyPressedEvent {
	return &KeyPressedEvent{handled: false, keyCode: keyCode, repeatCount: repeatCount}
}

func (event *KeyPressedEvent) GetEventType() EventType { return EventTypeKeyPressed }
func (event *KeyPressedEvent) GetName() string         { return "KeyPressed" }
func (event *KeyPressedEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryKeyboard
}
func (event *KeyPressedEvent) String() string {
	return fmt.Sprintf("%v: %v (%v)", event.GetName(), event.KeyCode(), event.RepeatCount())
}
func (event *KeyPressedEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *KeyPressedEvent) IsHandled() bool  { return event.handled }
func (event *KeyPressedEvent) SetHandled()      { event.handled = true }
func (event *KeyPressedEvent) KeyCode() KeyCode { return event.keyCode }
func (event *KeyPressedEvent) RepeatCount() int { return event.repeatCount }

func NewKeyReleasedEvent(keyCode KeyCode) *KeyReleasedEvent {
	return &KeyReleasedEvent{handled: false, keyCode: keyCode}
}

func (event *KeyReleasedEvent) GetEventType() EventType { return EventTypeKeyReleased }
func (event *KeyReleasedEvent) GetName() string         { return "KeyReleased" }
func (event *KeyReleasedEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryKeyboard
}
func (event *KeyReleasedEvent) String() string {
	return fmt.Sprintf("%v: %v", event.GetName(), event.KeyCode())
}
func (event *KeyReleasedEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *KeyReleasedEvent) IsHandled() bool  { return event.handled }
func (event *KeyReleasedEvent) SetHandled()      { event.handled = true }
func (event *KeyReleasedEvent) KeyCode() KeyCode { return event.keyCode }

func NewKeyTypedEvent(character rune) *KeyTypedEvent {
	return &KeyTypedEvent{handled: false, character: character}
}

func (event *KeyTypedEvent) GetEventType() EventType { return EventTypeKeyTyped }
func (event *KeyTypedEvent) GetName() string         { return "KeyTyped" }
func (event *KeyTypedEvent) GetCategoryFlags() EventCategory {
	return EventCategoryInput | EventCategoryKeyboard
}
func (event *KeyTypedEvent) String() string {
	return fmt.Sprintf("%v: %v", event.GetName(), event.Character())
}
func (event *KeyTypedEvent) IsInCategory(category EventCategory) bool {
	return event.GetCategoryFlags()&category > 0
}
func (event *KeyTypedEvent) IsHandled() bool { return event.handled }
func (event *KeyTypedEvent) SetHandled()     { event.handled = true }
func (event *KeyTypedEvent) Character() rune { return event.character }
