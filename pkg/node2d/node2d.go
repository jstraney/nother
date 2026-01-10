package node2d

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node"
)

// Node2D is an embeddable struct that implements the Node interface for 2D entities.
// It embeds BaseNode for hierarchy, lifecycle, and signaling, and adds 2D-specific positioning.
type Node2D struct {
	*node.BaseNode
	position rl.Vector2
	posMu    sync.RWMutex
}

// New creates a new Node2D with the given ID and name.
func New(id string, name string) *Node2D {
	return &Node2D{
		BaseNode: node.NewBaseNode(id, name),
		position: rl.NewVector2(0, 0),
	}
}

// Position returns the position of this node in global 2D space.
func (n *Node2D) Position() rl.Vector2 {
	n.posMu.RLock()
	defer n.posMu.RUnlock()
	return n.position
}

// SetPosition sets the position of this node in global 2D space.
func (n *Node2D) SetPosition(pos rl.Vector2) {
	n.posMu.Lock()
	defer n.posMu.Unlock()
	n.position = pos
}

// Render draws this node and all its children in 2D space.
// The implementation is provided by subclasses via embedding.
func (n *Node2D) Render() {
	if !n.Enabled() {
		return
	}

	children := n.Children()

	for _, child := range children {
		// Child nodes should implement their own rendering
		// This is called by the rendering pipeline
		if renderable, ok := child.(interface{ Render() }); ok {
			renderable.Render()
		}
	}
}

// GlobalPosition returns the global position of this node,
// accounting for parent positions in the hierarchy.
func (n *Node2D) GlobalPosition() rl.Vector2 {
	n.posMu.RLock()
	pos := n.position
	parent := n.BaseNode.Parent()
	n.posMu.RUnlock()

	if parent == nil {
		return pos
	}

	if parentNode2D, ok := parent.(*Node2D); ok {
		parentPos := parentNode2D.GlobalPosition()
		return rl.NewVector2(pos.X+parentPos.X, pos.Y+parentPos.Y)
	}

	return pos
}
