package area2d

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// Circle is a 2D circular area node that can optionally render a filled circle.
type Circle struct {
	*node2d.Node2D
	radius   float32
	visible  bool
	fillColor rl.Color
}

// New creates a new Circle with the given ID and name.
func NewCircle(id string, name string) *Circle {
	return &Circle{
		Node2D:    node2d.New(id, name),
		radius:    1.0,
		visible:   false,
		fillColor: rl.White,
	}
}

// SetRadius sets the radius of the circle.
func (c *Circle) SetRadius(radius float32) {
	c.radius = radius
}

// Radius returns the radius of the circle.
func (c *Circle) Radius() float32 {
	return c.radius
}

// SetVisible sets whether the circle should be rendered.
func (c *Circle) SetVisible(visible bool) {
	c.visible = visible
}

// Visible returns whether the circle is currently visible.
func (c *Circle) Visible() bool {
	return c.visible
}

// SetFill sets the fill color of the circle.
func (c *Circle) SetFill(color rl.Color) {
	c.fillColor = color
}

// Fill returns the current fill color.
func (c *Circle) Fill() rl.Color {
	return c.fillColor
}

// Render draws the circle if it is visible and enabled.
func (c *Circle) Render() {
	if !c.Enabled() || !c.visible {
		c.Node2D.Render()
		return
	}

	pos := c.GlobalPosition()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), c.radius, c.fillColor)

	// Render children
	c.Node2D.Render()
}
