package scene

import (
	"github.com/jstraney/nother/pkg/node"
	"github.com/jstraney/nother/pkg/node2d"
	"github.com/jstraney/nother/pkg/signal"
)

// BaseScene provides a basic implementation of the Scene interface.
// It can be embedded in concrete scene types to avoid boilerplate.
type BaseScene struct {
	id       string
	paused   bool
	rootNode *node2d.Node2D
}

// NewBaseScene creates a new BaseScene with the given ID.
func NewBaseScene(id string) *BaseScene {
	return &BaseScene{
		id:       id,
		paused:   false,
		rootNode: node2d.New(id+"-root", "Root"),
	}
}

// ID returns the scene ID.
func (s *BaseScene) ID() string {
	return s.id
}

// Name returns the scene name (same as ID for base implementation).
func (s *BaseScene) Name() string {
	return s.id
}

// SetName is a no-op for the base implementation.
func (s *BaseScene) SetName(name string) {
	// Not implemented in base
}

// Root returns the root node.
func (s *BaseScene) Root() node.Node {
	return s.rootNode
}

// Pause pauses the scene.
func (s *BaseScene) Pause() {
	s.paused = true
	s.rootNode.Disable()
}

// Resume resumes the scene.
func (s *BaseScene) Resume() {
	s.paused = false
	s.rootNode.Enable()
}

// Paused returns whether the scene is paused.
func (s *BaseScene) Paused() bool {
	return s.paused
}

// Update updates all nodes in the scene.
func (s *BaseScene) Update(deltaTime float64) {
	if s.paused {
		return
	}
	s.rootNode.Update(deltaTime)
}

// Load is a no-op in the base implementation.
func (s *BaseScene) Load() error {
	return nil
}

// Unload is a no-op in the base implementation.
func (s *BaseScene) Unload() error {
	return nil
}

// Emit sends a signal using the global signal bus.
func (s *BaseScene) Emit(signalID string, payload any) {
	signal.GlobalBus.Emit(signalID, payload)
}

// On subscribes a callback to a signal using the global signal bus.
func (s *BaseScene) On(signalID string, callback signal.Callback) {
	signal.GlobalBus.On(signalID, callback)
}

// Off unsubscribes all callbacks from a signal using the global signal bus.
func (s *BaseScene) Off(signalID string) {
	signal.GlobalBus.Off(signalID)
}
