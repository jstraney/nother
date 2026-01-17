package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// AnchorPoint defines where a component is anchored in its container.
type AnchorPoint int

const (
	AnchorTopLeft AnchorPoint = iota
	AnchorTopCenter
	AnchorTopRight
	AnchorCenterLeft
	AnchorCenter
	AnchorCenterRight
	AnchorBottomLeft
	AnchorBottomCenter
	AnchorBottomRight
)

// Layoutable defines components that participate in GuiContext layout management.
type Layoutable interface {
	// Anchor returns the anchor point for this component
	Anchor() AnchorPoint

	// AnchorOffset returns the offset from the anchor point
	AnchorOffset() rl.Vector2

	// PreferredSize returns the component's preferred dimensions
	PreferredSize() (width, height float32)

	// MinSize returns the minimum allowed dimensions
	MinSize() (width, height float32)

	// MaxSize returns the maximum allowed dimensions (0 means unlimited)
	MaxSize() (width, height float32)

	// ApplyLayout is called by GuiContext with computed final position and size
	ApplyLayout(x, y, width, height float32)
}

// LayoutConstraints holds size constraints for layout calculation.
type LayoutConstraints struct {
	Anchor         AnchorPoint
	AnchorOffset   rl.Vector2
	PreferredWidth float32
	PreferredHeight float32
	MinWidth       float32
	MinHeight      float32
	MaxWidth       float32
	MaxHeight      float32
}

// ComputedLayout represents the final computed position and size.
type ComputedLayout struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}
