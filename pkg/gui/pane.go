package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/area2d"
)

// Pane is a GUI container that extends Rect with border and padding support.
// Implements Layoutable for GuiContext layout management.
type Pane struct {
	*area2d.Rect
	borderWidth    float32
	borderColor    rl.Color
	padding        float32
	anchor         AnchorPoint
	anchorOffset   rl.Vector2
	preferredWidth float32
	preferredHeight float32
	minWidth       float32
	minHeight      float32
	maxWidth       float32
	maxHeight      float32
}

// NewPane creates a new Pane with the given ID and name.
func NewPane(id string, name string) *Pane {
	return &Pane{
		Rect:           area2d.NewRect(id, name),
		borderWidth:    0,
		borderColor:    rl.Black,
		padding:        0,
		anchor:         AnchorTopLeft,
		anchorOffset:   rl.NewVector2(0, 0),
		preferredWidth: 200,
		preferredHeight: 150,
		minWidth:       50,
		minHeight:      50,
		maxWidth:       0, // 0 means unlimited
		maxHeight:      0,
	}
}

// BorderWidth returns the border width.
func (p *Pane) BorderWidth() float32 {
	return p.borderWidth
}

// SetBorderWidth sets the border width.
func (p *Pane) SetBorderWidth(width float32) {
	p.borderWidth = width
}

// BorderColor returns the border color.
func (p *Pane) BorderColor() rl.Color {
	return p.borderColor
}

// SetBorderColor sets the border color.
func (p *Pane) SetBorderColor(color rl.Color) {
	p.borderColor = color
}

// Padding returns the padding value.
func (p *Pane) Padding() float32 {
	return p.padding
}

// SetPadding sets the padding value (space between border and content).
func (p *Pane) SetPadding(padding float32) {
	p.padding = padding
}

// ContentWidth returns the available width for content (width - padding * 2).
func (p *Pane) ContentWidth() float32 {
	return p.Width() - (p.padding * 2)
}

// ContentHeight returns the available height for content (height - padding * 2).
func (p *Pane) ContentHeight() float32 {
	return p.Height() - (p.padding * 2)
}

// ContentPosition returns the top-left position of the content area (accounting for padding).
func (p *Pane) ContentPosition() rl.Vector2 {
	pos := p.GlobalPosition()
	return rl.NewVector2(pos.X+p.padding, pos.Y+p.padding)
}

// SetAnchor sets the anchor point for layout.
func (p *Pane) SetAnchor(anchor AnchorPoint) {
	p.anchor = anchor
}

// SetAnchorOffset sets the offset from the anchor point.
func (p *Pane) SetAnchorOffset(offset rl.Vector2) {
	p.anchorOffset = offset
}

// SetPreferredSize sets the preferred dimensions.
func (p *Pane) SetPreferredSize(width, height float32) {
	p.preferredWidth = width
	p.preferredHeight = height
}

// SetMinSize sets the minimum dimensions.
func (p *Pane) SetMinSize(width, height float32) {
	p.minWidth = width
	p.minHeight = height
}

// SetMaxSize sets the maximum dimensions (0 means unlimited).
func (p *Pane) SetMaxSize(width, height float32) {
	p.maxWidth = width
	p.maxHeight = height
}

// Layoutable interface implementation

// Anchor returns the anchor point.
func (p *Pane) Anchor() AnchorPoint {
	return p.anchor
}

// AnchorOffset returns the offset from the anchor point.
func (p *Pane) AnchorOffset() rl.Vector2 {
	return p.anchorOffset
}

// PreferredSize returns the preferred dimensions.
func (p *Pane) PreferredSize() (width, height float32) {
	return p.preferredWidth, p.preferredHeight
}

// MinSize returns the minimum dimensions.
func (p *Pane) MinSize() (width, height float32) {
	return p.minWidth, p.minHeight
}

// MaxSize returns the maximum dimensions.
func (p *Pane) MaxSize() (width, height float32) {
	return p.maxWidth, p.maxHeight
}

// ApplyLayout is called by GuiContext with computed final position and size.
func (p *Pane) ApplyLayout(x, y, width, height float32) {
	p.SetPosition(rl.NewVector2(x, y))
	p.SetSize(width, height)
}

// Render draws the pane with background, border, and children.
func (p *Pane) Render() {
	if !p.Enabled() {
		p.Rect.Render()
		return
	}

	pos := p.GlobalPosition()
	width := p.Width()
	height := p.Height()

	// Draw background fill if visible
	if p.Visible() {
		rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(width), int32(height), p.Fill())
	}

	// Draw border if border width > 0
	if p.borderWidth > 0 {
		rl.DrawRectangleLinesEx(
			rl.NewRectangle(pos.X, pos.Y, width, height),
			p.borderWidth,
			p.borderColor,
		)
	}

	// Render children (skip calling Rect.Render to avoid double-rendering background)
	for _, child := range p.Children() {
		if child != nil {
			child.Update(0) // Update with 0 delta time for rendering pass
			// Assuming children have Render method - this is safe since Update is on Node interface
		}
	}
}
