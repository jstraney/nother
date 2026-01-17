package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// Label is a 2D GUI node that renders text at a position.
// Auto-sizes based on text content.
type Label struct {
	*node2d.Node2D
	text     string
	fontSize float32
	spacing  float32
	color    rl.Color
	font     *rl.Font
	width    float32
	height   float32
}

// New creates a new Label with the given ID and name.
func New(id string, name string) *Label {
	return &Label{
		Node2D:   node2d.New(id, name),
		text:     "",
		fontSize: 20.0,
		spacing:  1.0,
		color:    rl.Black,
		font:     nil, // nil means use default font
	}
}

// Text returns the label's text content.
func (l *Label) Text() string {
	return l.text
}

// SetText updates the label's text content and auto-sizes the label.
func (l *Label) SetText(text string) {
	l.text = text
	l.measureAndResize()
}

// measureAndResize measures the text and updates width/height.
func (l *Label) measureAndResize() {
	if l.text == "" {
		l.width = 0
		l.height = l.fontSize
		return
	}

	if l.font != nil {
		// Measure with custom font
		size := rl.MeasureTextEx(*l.font, l.text, l.fontSize, l.spacing)
		l.width = size.X
		l.height = size.Y
	} else {
		// Measure with default font
		textWidth := rl.MeasureText(l.text, int32(l.fontSize))
		l.width = float32(textWidth)
		l.height = l.fontSize
	}
}

// FontSize returns the label's font size.
func (l *Label) FontSize() float32 {
	return l.fontSize
}

// SetFontSize updates the label's font size and re-measures.
func (l *Label) SetFontSize(size float32) {
	l.fontSize = size
	l.measureAndResize()
}

// Spacing returns the character spacing.
func (l *Label) Spacing() float32 {
	return l.spacing
}

// SetSpacing updates the character spacing.
func (l *Label) SetSpacing(spacing float32) {
	l.spacing = spacing
}

// Font returns the current font (nil means default font).
func (l *Label) Font() *rl.Font {
	return l.font
}

// SetFont updates the font and re-measures. Pass nil to use the default font.
func (l *Label) SetFont(font *rl.Font) {
	l.font = font
	l.measureAndResize()
}

// Color returns the label's text color.
func (l *Label) Color() rl.Color {
	return l.color
}

// SetColor updates the label's text color.
func (l *Label) SetColor(color rl.Color) {
	l.color = color
}

// Width returns the measured width of the text.
func (l *Label) Width() float32 {
	return l.width
}

// Height returns the measured height of the text.
func (l *Label) Height() float32 {
	return l.height
}

// ContentWidth returns the content width (same as Width for labels).
func (l *Label) ContentWidth() float32 {
	return l.width
}

// ContentHeight returns the content height (same as Height for labels).
func (l *Label) ContentHeight() float32 {
	return l.height
}

// Render draws the label text at its global position.
func (l *Label) Render() {
	if !l.Enabled() {
		return
	}

	pos := l.GlobalPosition()

	if l.font != nil {
		// Use custom font with DrawTextEx
		rl.DrawTextEx(*l.font, l.text, pos, l.fontSize, l.spacing, l.color)
	} else {
		// Use default font with DrawText
		rl.DrawText(l.text, int32(pos.X), int32(pos.Y), int32(l.fontSize), l.color)
	}

	// Render children
	l.Node2D.Render()
}
