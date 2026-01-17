package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/area2d"
)

// Checkbox is a toggleable checkbox GUI element with optional label.
type Checkbox struct {
	*area2d.Rect
	checked        bool
	label          *Label
	boxSize        float32
	borderWidth    float32
	borderColor    rl.Color
	bgColor        rl.Color
	checkColor     rl.Color
	hoverColor     rl.Color
	isHovered      bool
	onChangeHandler func(bool)
}

// NewCheckbox creates a new Checkbox with the given ID and name.
func NewCheckbox(id string, name string) *Checkbox {
	checkbox := &Checkbox{
		Rect:            area2d.NewRect(id, name),
		checked:         false,
		label:           New(id+"_label", name+"_label"),
		boxSize:         20,
		borderWidth:     2,
		borderColor:     rl.Black,
		bgColor:         rl.White,
		checkColor:      rl.DarkGreen,
		hoverColor:      rl.LightGray,
		isHovered:       false,
		onChangeHandler: nil,
	}

	// Configure label
	checkbox.label.SetText("")
	checkbox.label.SetColor(rl.Black)
	checkbox.label.SetFontSize(16.0)

	// Add label as child
	checkbox.AddChild(checkbox.label)

	// Set default size (box + label space)
	checkbox.SetSize(150, checkbox.boxSize)
	checkbox.SetVisible(true)

	return checkbox
}

// Checked returns whether the checkbox is checked.
func (c *Checkbox) Checked() bool {
	return c.checked
}

// SetChecked sets the checked state.
func (c *Checkbox) SetChecked(checked bool) {
	if c.checked != checked {
		c.checked = checked
		if c.onChangeHandler != nil {
			c.onChangeHandler(checked)
		}
	}
}

// Toggle toggles the checked state.
func (c *Checkbox) Toggle() {
	c.SetChecked(!c.checked)
}

// Label returns the label component.
func (c *Checkbox) Label() *Label {
	return c.label
}

// SetLabelText sets the label text.
func (c *Checkbox) SetLabelText(text string) {
	c.label.SetText(text)
}

// OnChange sets the callback function called when the checkbox state changes.
func (c *Checkbox) OnChange(handler func(bool)) {
	c.onChangeHandler = handler
}

// SetBoxSize sets the size of the checkbox box.
func (c *Checkbox) SetBoxSize(size float32) {
	c.boxSize = size
}

// SetBorderColor sets the border color.
func (c *Checkbox) SetBorderColor(color rl.Color) {
	c.borderColor = color
}

// SetBgColor sets the background color.
func (c *Checkbox) SetBgColor(color rl.Color) {
	c.bgColor = color
}

// SetCheckColor sets the check mark color.
func (c *Checkbox) SetCheckColor(color rl.Color) {
	c.checkColor = color
}

// SetHoverColor sets the hover background color.
func (c *Checkbox) SetHoverColor(color rl.Color) {
	c.hoverColor = color
}

// Update handles checkbox interaction.
func (c *Checkbox) Update(dt float32) {
	if !c.Enabled() {
		return
	}

	mousePos := rl.GetMousePosition()
	pos := c.GlobalPosition()

	// Check if mouse is over the checkbox box
	c.isHovered = mousePos.X >= pos.X &&
		mousePos.X <= pos.X+c.boxSize &&
		mousePos.Y >= pos.Y &&
		mousePos.Y <= pos.Y+c.boxSize

	// Handle click to toggle
	if c.isHovered && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		c.Toggle()
	}

	// Position label to the right of the box
	labelX := c.boxSize + 8 // 8px spacing
	labelY := (c.boxSize - c.label.FontSize()) / 2
	c.label.SetPosition(rl.NewVector2(labelX, labelY))

	// Call base Update
	c.Rect.Update(dt)
}

// Render draws the checkbox box, check mark, and label.
func (c *Checkbox) Render() {
	if !c.Enabled() {
		return
	}

	pos := c.GlobalPosition()

	// Determine background color
	bgColor := c.bgColor
	if c.isHovered {
		bgColor = c.hoverColor
	}

	// Draw checkbox box
	rl.DrawRectangle(
		int32(pos.X),
		int32(pos.Y),
		int32(c.boxSize),
		int32(c.boxSize),
		bgColor,
	)

	// Draw border
	rl.DrawRectangleLinesEx(
		rl.NewRectangle(pos.X, pos.Y, c.boxSize, c.boxSize),
		c.borderWidth,
		c.borderColor,
	)

	// Draw check mark if checked
	if c.checked {
		padding := c.boxSize * 0.2
		rl.DrawRectangle(
			int32(pos.X+padding),
			int32(pos.Y+padding),
			int32(c.boxSize-padding*2),
			int32(c.boxSize-padding*2),
			c.checkColor,
		)
	}

	// Render label
	c.label.Render()
}
