package scene

import (
	"encoding/json"

	"github.com/jstraney/nother/pkg/node"
	"github.com/jstraney/nother/pkg/signal"
)

// BaseScene provides a basic implementation of the Scene interface.
// It can be embedded in concrete scene types to avoid boilerplate.
type BaseScene struct {
	id       string
	paused   bool
	rootNode node.Node
}

// NewBaseScene creates a new BaseScene with the given ID.
func NewBaseScene(id string) *BaseScene {
	return &BaseScene{
		id:       id,
		paused:   false,
		rootNode: nil,
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

// Set the scene root node
func (s *BaseScene) SetRoot(rootNode node.Node) {
	s.rootNode = rootNode
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

type BaseSceneSerial struct {
	ID       string    `json:"id"`
	Paused   bool      `json:"paused"`
	RootNode node.Node `json:"rootNode"`
}

// UnmarshalJSON unmarshals JSON data into a BaseScene using BaseSceneSerial as an intermediate.
func (s *BaseScene) UnmarshalJSON(data []byte) error {
	var serial BaseSceneSerial
	if err := json.Unmarshal(data, &serial); err != nil {
		return err
	}

	s.id = serial.ID
	s.paused = serial.Paused
	s.rootNode = serial.RootNode

	return nil
}
