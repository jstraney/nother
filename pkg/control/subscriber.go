package control

import "sync"

// Subscriber listens for specific input events and dispatches them to registered handlers.
// Nodes receive a Subscriber and call its methods to bind callbacks.
type Subscriber struct {
	eventID  string
	handlers []Handler
	mu       sync.RWMutex
}

// Subscribe registers a handler for this event.
func (s *Subscriber) Subscribe(handler Handler) {
	if handler == nil {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers = append(s.handlers, handler)
}

// Unsubscribe removes all handlers.
func (s *Subscriber) Unsubscribe() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handlers = s.handlers[:0]
}

// dispatch sends an event to all handlers.
func (s *Subscriber) dispatch(event Event) {
	s.mu.RLock()
	handlers := make([]Handler, len(s.handlers))
	copy(handlers, s.handlers)
	s.mu.RUnlock()

	for _, handler := range handlers {
		handler(event)
	}
}

// EventID returns the event ID this subscriber is listening to.
func (s *Subscriber) EventID() string {
	return s.eventID
}
