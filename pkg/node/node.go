package node

import (
	"github.com/jstraney/nother/pkg/signal"
)

// Node represents a generic entity in the game engine.
// All game entities (Node3D, AudioNode, GUINode, etc.) should implement this interface.
// The interface is intentionally minimal to avoid constraining implementations unnecessarily.
type Node interface {
	signal.Emitter
	signal.Listener
	// ID returns a unique identifier for this node.
	ID() string

	// Name returns the human-readable name of this node.
	Name() string

	// SetName updates the name of this node.
	SetName(name string)

	// Parent returns the parent node, or nil if this is a root node.
	Parent() Node

	// SetParent sets the parent node. Implementations should handle reparenting logic.
	SetParent(parent Node)

	// Children returns a list of child nodes.
	Children() []Node

	// AddChild adds a child node.
	AddChild(child Node) error

	// RemoveChild removes a child node.
	RemoveChild(child Node) error

	// Update is called once per frame to update the node's state.
	// deltaTime is the time elapsed since the last update, in seconds.
	Update(deltaTime float64)

	// Enable activates the node and its update cycle.
	Enable()

	// Disable deactivates the node and its update cycle.
	Disable()

	// Enabled returns whether this node is currently enabled.
	Enabled() bool

	// Destroy is called when the node is being removed from the scene.
	// Implementations should clean up any resources held by this node.
	Destroy()
}
