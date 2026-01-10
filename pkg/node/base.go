package node

import (
	"fmt"
	"sync"

	"github.com/jstraney/nother/pkg/signal"
)

// BaseNode implements the Node interface with common functionality
// for all node types. It handles identity, hierarchy, lifecycle, and signaling.
// Specific node types (Node2D, Node3D, etc.) should embed this struct.
type BaseNode struct {
	id       string
	name     string
	parent   Node
	children []Node
	enabled  bool
	mu       sync.RWMutex
}

// NewBaseNode creates a new BaseNode with the given ID and name.
func NewBaseNode(id string, name string) *BaseNode {
	return &BaseNode{
		id:       id,
		name:     name,
		children: make([]Node, 0),
		enabled:  true,
	}
}

// ID returns the unique identifier for this node.
func (n *BaseNode) ID() string {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.id
}

// Name returns the human-readable name of this node.
func (n *BaseNode) Name() string {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.name
}

// SetName updates the name of this node.
func (n *BaseNode) SetName(name string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.name = name
}

// Parent returns the parent node, or nil if this is a root node.
func (n *BaseNode) Parent() Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.parent
}

// SetParent sets the parent node.
func (n *BaseNode) SetParent(parent Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.parent = parent
}

// Children returns a list of child nodes.
func (n *BaseNode) Children() []Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	children := make([]Node, len(n.children))
	copy(children, n.children)
	return children
}

// AddChild adds a child node.
func (n *BaseNode) AddChild(child Node) error {
	if child == nil {
		return fmt.Errorf("cannot add nil child")
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	// Prevent adding duplicate children
	for _, c := range n.children {
		if c == child {
			return fmt.Errorf("child already exists")
		}
	}

	n.children = append(n.children, child)
	child.SetParent(n)
	return nil
}

// RemoveChild removes a child node.
func (n *BaseNode) RemoveChild(child Node) error {
	if child == nil {
		return fmt.Errorf("cannot remove nil child")
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
			child.SetParent(nil)
			return nil
		}
	}

	return fmt.Errorf("child not found")
}

// Update is called once per frame to update the node's state.
// BaseNode updates all child nodes.
func (n *BaseNode) Update(deltaTime float64) {
	if !n.Enabled() {
		return
	}

	n.mu.RLock()
	children := make([]Node, len(n.children))
	copy(children, n.children)
	n.mu.RUnlock()

	for _, child := range children {
		child.Update(deltaTime)
	}
}

// Enable activates the node and its update cycle.
func (n *BaseNode) Enable() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.enabled = true
}

// Disable deactivates the node and its update cycle.
func (n *BaseNode) Disable() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.enabled = false
}

// Enabled returns whether this node is currently enabled.
func (n *BaseNode) Enabled() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.enabled
}

// Destroy is called when the node is being removed from the scene.
func (n *BaseNode) Destroy() {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Destroy all children
	for _, child := range n.children {
		child.Destroy()
	}

	n.children = nil
	n.parent = nil
}

// Emit sends a signal with optional payload using the global signal bus.
func (n *BaseNode) Emit(signalID string, payload any) {
	signal.GlobalBus.Emit(signalID, payload)
}

// On subscribes a callback to a signal using the global signal bus.
func (n *BaseNode) On(signalID string, callback signal.Callback) {
	signal.GlobalBus.On(signalID, callback)
}

// Off unsubscribes all callbacks from a signal using the global signal bus.
func (n *BaseNode) Off(signalID string) {
	signal.GlobalBus.Off(signalID)
}

// RLock acquires a read lock on the node's state.
func (n *BaseNode) RLock() {
	n.mu.RLock()
}

// RUnlock releases a read lock on the node's state.
func (n *BaseNode) RUnlock() {
	n.mu.RUnlock()
}

// Lock acquires a write lock on the node's state.
func (n *BaseNode) Lock() {
	n.mu.Lock()
}

// Unlock releases a write lock on the node's state.
func (n *BaseNode) Unlock() {
	n.mu.Unlock()
}
