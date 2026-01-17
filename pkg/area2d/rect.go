package area2d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// Rect is a 2D rectangular area node that can optionally render a filled rectangle.
type Rect struct {
	*node2d.Node2D
	width     float32
	height    float32
	visible   bool
	fillColor rl.Color
}

// New creates a new Rect with the given ID and name.
func NewRect(id string, name string) *Rect {
	return &Rect{
		Node2D:    node2d.New(id, name),
		width:     1.0,
		height:    1.0,
		visible:   false,
		fillColor: rl.White,
	}
}

// SetWidth sets the width of the rectangle.
func (r *Rect) SetWidth(width float32) {
	r.width = width
}

// Width returns the width of the rectangle.
func (r *Rect) Width() float32 {
	return r.width
}

// SetHeight sets the height of the rectangle.
func (r *Rect) SetHeight(height float32) {
	r.height = height
}

// Height returns the height of the rectangle.
func (r *Rect) Height() float32 {
	return r.height
}

// SetSize sets both width and height of the rectangle.
func (r *Rect) SetSize(width, height float32) {
	r.width = width
	r.height = height
}

// SetVisible sets whether the rectangle should be rendered.
func (r *Rect) SetVisible(visible bool) {
	r.visible = visible
}

// Visible returns whether the rectangle is currently visible.
func (r *Rect) Visible() bool {
	return r.visible
}

// SetFill sets the fill color of the rectangle.
func (r *Rect) SetFill(color rl.Color) {
	r.fillColor = color
}

// Fill returns the current fill color.
func (r *Rect) Fill() rl.Color {
	return r.fillColor
}

// Render draws the rectangle if it is visible and enabled.
func (r *Rect) Render() {
	if !r.Enabled() || !r.visible {
		r.Node2D.Render()
		return
	}

	pos := r.GlobalPosition()
	// Draw rectangle with top-left corner at position
	rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(r.width), int32(r.height), r.fillColor)
	// Render children
	r.Node2D.Render()
}
