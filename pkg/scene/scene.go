package scene

import (
	"github.com/jstraney/nother/pkg/node"
	"github.com/jstraney/nother/pkg/signal"
)

// Scene represents a collection of nodes that form a logical game scene.
// Scenes can be paused and resumed, allowing the engine to manage multiple scenes
// (e.g., menus, levels, overlays) independently.
type Scene interface {
	signal.Emitter
	signal.Listener
	// ID returns a unique identifier for this scene.
	ID() string

	// Name returns the human-readable name of this scene.
	Name() string

	// SetName updates the name of this scene.
	SetName(name string)

	// Root returns the root node of this scene's node tree.
	// All nodes in the scene should be descendants of this root.
	Root() node.Node

	// Pause pauses the scene, stopping updates for all nodes.
	Pause()

	// Resume resumes the scene, resuming updates for all nodes.
	Resume()

	// Paused returns whether this scene is currently paused.
	Paused() bool

	// Update updates all nodes in the scene.
	// deltaTime is the time elapsed since the last update, in seconds.
	// Update should only proceed if the scene is not paused.
	Update(deltaTime float64)

	// Load is called when the scene is first loaded.
	// Implementations should initialize scene resources here.
	Load() error

	// Unload is called when the scene is being removed.
	// Implementations should clean up scene resources here.
	Unload() error
}
