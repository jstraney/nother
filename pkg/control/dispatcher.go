package control

import (
	"fmt"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// EventType represents the type of input event.
type EventType int

const (
	KeyPressed EventType = iota
	KeyDown
	KeyReleased
	MousePressed
	MouseDown
	MouseReleased
	MouseWheelMove
)

// Event represents an input event with optional data.
type Event struct {
	Type EventType
	Data interface{} // Could be key code, mouse button, or scroll delta
}

// Handler is a function that handles an input event.
type Handler func(event Event)

// Dispatcher manages input event polling and dispatching to subscribers.
type Dispatcher struct {
	subscribers        map[string]*Subscriber
	watchedKeys        map[int32]bool
	mouseEventsEnabled bool
	mu                 sync.RWMutex
}

// New creates a new Dispatcher.
func New() *Dispatcher {
	return &Dispatcher{
		subscribers:        make(map[string]*Subscriber),
		watchedKeys:        make(map[int32]bool),
		mouseEventsEnabled: false,
	}
}

// WatchKeys registers keys to be polled for KeyDown events.
// By default, no keys are watched, so no KeyDown polling occurs.
// This should be called before the main game loop starts.
func (d *Dispatcher) WatchKeys(keys ...int32) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, key := range keys {
		d.watchedKeys[key] = true
	}
}

// UnwatchKey stops watching a specific key.
func (d *Dispatcher) UnwatchKey(key int32) {
	d.mu.Lock()
	defer d.mu.Unlock()

	delete(d.watchedKeys, key)
}

// UnwatchAllKeys stops watching all keys.
func (d *Dispatcher) UnwatchAllKeys() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.watchedKeys = make(map[int32]bool)
}

// WatchMouseEvents enables polling for mouse input events.
// By default, mouse events are not polled.
func (d *Dispatcher) WatchMouseEvents() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.mouseEventsEnabled = true
}

// UnwatchMouseEvents disables polling for mouse input events.
func (d *Dispatcher) UnwatchMouseEvents() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.mouseEventsEnabled = false
}

// NewSubscriber creates a new Subscriber for the given event ID.
func (d *Dispatcher) NewSubscriber(eventID string) *Subscriber {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.subscribers[eventID]; !exists {
		d.subscribers[eventID] = &Subscriber{
			eventID:  eventID,
			handlers: make([]Handler, 0),
		}
	}

	return d.subscribers[eventID]
}

// dispatch sends an event to all subscribers of a specific event ID.
func (d *Dispatcher) dispatch(eventID string, event Event) {
	d.mu.RLock()
	subscriber, exists := d.subscribers[eventID]
	d.mu.RUnlock()

	if !exists {
		return
	}

	subscriber.dispatch(event)
}

// Poll checks raylib input state and dispatches events accordingly.
// This should be called once per frame in the main game loop.
func (d *Dispatcher) Poll() {
	if len(d.watchedKeys) > 0 {
		// Poll keyboard events
		d.pollKeyboardEvents()
	}

	// Poll mouse events only if enabled
	d.mu.RLock()
	mouseEnabled := d.mouseEventsEnabled
	d.mu.RUnlock()

	if mouseEnabled {
		d.pollMouseInputEvents()
	}
}

// pollKeyboardEvents checks for keyboard input and dispatches events.
func (d *Dispatcher) pollKeyboardEvents() {
	// Check for any pressed key
	key := rl.GetKeyPressed()
	if key != 0 {
		eventID := fmt.Sprintf("key:pressed:%d", key)
		d.dispatch(eventID, Event{Type: KeyPressed, Data: key})
	}

	// Get the watched keys
	d.mu.RLock()
	watchedKeys := make([]int32, 0, len(d.watchedKeys))
	for key := range d.watchedKeys {
		watchedKeys = append(watchedKeys, key)
	}
	d.mu.RUnlock()

	// Poll only watched keys for KeyDown events
	for _, key := range watchedKeys {
		if rl.IsKeyDown(key) {
			eventID := fmt.Sprintf("key:down:%d", key)
			d.dispatch(eventID, Event{Type: KeyDown, Data: key})
		}
	}
}

// pollMouseInputEvents checks for mouse input and dispatches events.
func (d *Dispatcher) pollMouseInputEvents() {
	mouseButtons := []rl.MouseButton{
		rl.MouseLeftButton,
		rl.MouseRightButton,
		rl.MouseMiddleButton,
	}

	for _, button := range mouseButtons {
		if rl.IsMouseButtonPressed(button) {
			eventID := fmt.Sprintf("mouse:pressed:%d", button)
			d.dispatch(eventID, Event{Type: MousePressed, Data: button})
		}

		if rl.IsMouseButtonDown(button) {
			eventID := fmt.Sprintf("mouse:down:%d", button)
			d.dispatch(eventID, Event{Type: MouseDown, Data: button})
		}

		if rl.IsMouseButtonReleased(button) {
			eventID := fmt.Sprintf("mouse:released:%d", button)
			d.dispatch(eventID, Event{Type: MouseReleased, Data: button})
		}
	}

	// Check mouse wheel
	wheelMove := rl.GetMouseWheelMove()
	if wheelMove != 0 {
		eventID := "mouse:wheel"
		d.dispatch(eventID, Event{Type: MouseWheelMove, Data: wheelMove})
	}
}

// Helper functions for convenient event ID generation

// KeyPressedEventID returns the event ID for a key press.
func KeyPressedEventID(key int32) string {
	return fmt.Sprintf("key:pressed:%d", key)
}

// KeyDownEventID returns the event ID for a key being held down.
func KeyDownEventID(key int32) string {
	return fmt.Sprintf("key:down:%d", key)
}

// KeyReleasedEventID returns the event ID for a key release.
func KeyReleasedEventID(key int32) string {
	return fmt.Sprintf("key:released:%d", key)
}

// MousePressedEventID returns the event ID for a mouse button press.
func MousePressedEventID(button rl.MouseButton) string {
	return fmt.Sprintf("mouse:pressed:%d", button)
}

// MouseDownEventID returns the event ID for a mouse button being held.
func MouseDownEventID(button int32) string {
	return fmt.Sprintf("mouse:down:%d", button)
}

// MouseReleasedEventID returns the event ID for a mouse button release.
func MouseReleasedEventID(button int32) string {
	return fmt.Sprintf("mouse:released:%d", button)
}

// MouseWheelEventID returns the event ID for mouse wheel movement.
func MouseWheelEventID() string {
	return "mouse:wheel"
}
