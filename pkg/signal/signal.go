package signal

import "sync"

// Callback is a function type that handles signal emissions.
// It receives the signal payload (can be any type or nil).
type Callback func(payload any)

// Emitter emits signals with optional payloads.
type Emitter interface {
	// Emit sends a signal with the given ID and optional payload to all listeners.
	Emit(signalID string, payload any)
}

// Listener listens for signal emissions and can subscribe/unsubscribe from them.
type Listener interface {
	// On subscribes a callback to a signal. The callback will be called when the signal is emitted.
	On(signalID string, callback Callback)

	// Off unsubscribes all callbacks from a signal.
	Off(signalID string)
}

// Bus manages signal subscriptions and emissions.
type Bus struct {
	listeners map[string][]Callback
	mu        sync.RWMutex
}

// NewBus creates a new signal bus.
func NewBus() *Bus {
	return &Bus{
		listeners: make(map[string][]Callback),
	}
}

// On subscribes a callback to a signal.
func (b *Bus) On(signalID string, callback Callback) {
	if callback == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	b.listeners[signalID] = append(b.listeners[signalID], callback)
}

// Off unsubscribes all callbacks from a signal.
func (b *Bus) Off(signalID string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.listeners, signalID)
}

// Emit sends a signal with optional payload to all listeners.
func (b *Bus) Emit(signalID string, payload any) {
	b.mu.RLock()
	callbacks := make([]Callback, len(b.listeners[signalID]))
	copy(callbacks, b.listeners[signalID])
	b.mu.RUnlock()

	for _, callback := range callbacks {
		callback(payload)
	}
}

// GlobalBus is a package-level singleton for signal management.
var GlobalBus = NewBus()