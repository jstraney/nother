package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// Label is a 2D GUI node that renders text at a position.
type Label struct {
	*node2d.Node2D
	text     string
	fontSize int32
	color    rl.Color
}

// New creates a new Label with the given ID and name.
func New(id string, name string) *Label {
	return &Label{
		Node2D:   node2d.New(id, name),
		text:     "",
		fontSize: 20,
		color:    rl.Black,
	}
}

// Text returns the label's text content.
func (l *Label) Text() string {
	return l.text
}

// SetText updates the label's text content.
func (l *Label) SetText(text string) {
	l.text = text
}

// FontSize returns the label's font size.
func (l *Label) FontSize() int32 {
	return l.fontSize
}

// SetFontSize updates the label's font size.
func (l *Label) SetFontSize(size int32) {
	l.fontSize = size
}

// Color returns the label's text color.
func (l *Label) Color() rl.Color {
	return l.color
}

// SetColor updates the label's text color.
func (l *Label) SetColor(color rl.Color) {
	l.color = color
}

// Render draws the label text at its global position.
func (l *Label) Render() {
	if !l.Enabled() {
		return
	}

	pos := l.GlobalPosition()
	rl.DrawText(l.text, int32(pos.X), int32(pos.Y), l.fontSize, l.color)

	// Render children
	l.Node2D.Render()
}
