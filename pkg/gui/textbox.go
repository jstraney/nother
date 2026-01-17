package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/area2d"
)

// TextBox is an editable text input GUI element.
type TextBox struct {
	*area2d.Rect
	text          string
	placeholder   string
	maxLength     int
	fontSize      float32
	textColor     rl.Color
	bgColor       rl.Color
	borderColor   rl.Color
	focusedColor  rl.Color
	borderWidth   float32
	padding       float32
	isFocused     bool
	cursorPos     int
	cursorBlink   float64
	blinkInterval float64
	showCursor    bool
}

// NewTextBox creates a new TextBox with the given ID and name.
func NewTextBox(id string, name string) *TextBox {
	textBox := &TextBox{
		Rect:          area2d.NewRect(id, name),
		text:          "",
		placeholder:   "Enter text...",
		maxLength:     50,
		fontSize:      16.0,
		textColor:     rl.Black,
		bgColor:       rl.White,
		borderColor:   rl.Gray,
		focusedColor:  rl.Blue,
		borderWidth:   2,
		padding:       8,
		isFocused:     false,
		cursorPos:     0,
		cursorBlink:   0,
		blinkInterval: 0.5,
		showCursor:    true,
	}

	textBox.SetSize(200, 30)
	textBox.SetVisible(true)

	return textBox
}

// Text returns the current text content.
func (t *TextBox) Text() string {
	return t.text
}

// SetText sets the text content.
func (t *TextBox) SetText(text string) {
	if len(text) > t.maxLength {
		t.text = text[:t.maxLength]
	} else {
		t.text = text
	}
	t.cursorPos = len(t.text)
}

// Placeholder returns the placeholder text.
func (t *TextBox) Placeholder() string {
	return t.placeholder
}

// SetPlaceholder sets the placeholder text shown when empty.
func (t *TextBox) SetPlaceholder(placeholder string) {
	t.placeholder = placeholder
}

// MaxLength returns the maximum text length.
func (t *TextBox) MaxLength() int {
	return t.maxLength
}

// SetMaxLength sets the maximum allowed text length.
func (t *TextBox) SetMaxLength(length int) {
	t.maxLength = length
	if len(t.text) > length {
		t.text = t.text[:length]
		if t.cursorPos > length {
			t.cursorPos = length
		}
	}
}

// SetFontSize sets the font size for the text.
func (t *TextBox) SetFontSize(size float32) {
	t.fontSize = size
}

// SetTextColor sets the text color.
func (t *TextBox) SetTextColor(color rl.Color) {
	t.textColor = color
}

// SetBgColor sets the background color.
func (t *TextBox) SetBgColor(color rl.Color) {
	t.bgColor = color
	t.SetFill(color)
}

// SetBorderColor sets the border color.
func (t *TextBox) SetBorderColor(color rl.Color) {
	t.borderColor = color
}

// SetFocusedColor sets the border color when focused.
func (t *TextBox) SetFocusedColor(color rl.Color) {
	t.focusedColor = color
}

// IsFocused returns whether the textbox is currently focused.
func (t *TextBox) IsFocused() bool {
	return t.isFocused
}

// SetFocused sets the focus state.
func (t *TextBox) SetFocused(focused bool) {
	t.isFocused = focused
	if focused {
		t.cursorPos = len(t.text)
	}
}

// Update handles text input and cursor blinking.
func (t *TextBox) Update(dt float64) {
	if !t.Enabled() {
		return
	}

	mousePos := rl.GetMousePosition()
	pos := t.GlobalPosition()

	// Check for mouse click to focus/unfocus
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		wasInside := mousePos.X >= pos.X &&
			mousePos.X <= pos.X+t.Width() &&
			mousePos.Y >= pos.Y &&
			mousePos.Y <= pos.Y+t.Height()

		t.isFocused = wasInside
		if wasInside {
			t.cursorBlink = 0
			t.showCursor = true
		}
	}

	// Handle text input when focused
	if t.isFocused {
		// Get character input
		char := rl.GetCharPressed()
		for char > 0 {
			// Only accept printable characters and within max length
			if char >= 32 && char < 127 && len(t.text) < t.maxLength {
				t.text = t.text[:t.cursorPos] + string(rune(char)) + t.text[t.cursorPos:]
				t.cursorPos++
			}
			char = rl.GetCharPressed()
		}

		// Handle backspace
		if rl.IsKeyPressed(rl.KeyBackspace) && t.cursorPos > 0 {
			t.text = t.text[:t.cursorPos-1] + t.text[t.cursorPos:]
			t.cursorPos--
		}

		// Handle delete
		if rl.IsKeyPressed(rl.KeyDelete) && t.cursorPos < len(t.text) {
			t.text = t.text[:t.cursorPos] + t.text[t.cursorPos+1:]
		}

		// Handle left arrow
		if rl.IsKeyPressed(rl.KeyLeft) && t.cursorPos > 0 {
			t.cursorPos--
		}

		// Handle right arrow
		if rl.IsKeyPressed(rl.KeyRight) && t.cursorPos < len(t.text) {
			t.cursorPos++
		}

		// Handle home key
		if rl.IsKeyPressed(rl.KeyHome) {
			t.cursorPos = 0
		}

		// Handle end key
		if rl.IsKeyPressed(rl.KeyEnd) {
			t.cursorPos = len(t.text)
		}

		// Update cursor blink
		t.cursorBlink += dt
		if t.cursorBlink >= t.blinkInterval {
			t.showCursor = !t.showCursor
			t.cursorBlink = 0
		}
	}

	// Call base Update
	t.Rect.Update(dt)
}

// Render draws the textbox with text and cursor.
func (t *TextBox) Render() {
	if !t.Enabled() {
		return
	}

	pos := t.GlobalPosition()
	width := t.Width()
	height := t.Height()

	// Draw background
	rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(width), int32(height), t.bgColor)

	// Draw border with focus color if focused
	borderColor := t.borderColor
	if t.isFocused {
		borderColor = t.focusedColor
	}
	rl.DrawRectangleLinesEx(
		rl.NewRectangle(pos.X, pos.Y, width, height),
		t.borderWidth,
		borderColor,
	)

	// Draw text or placeholder
	textToDraw := t.text
	textColor := t.textColor
	if len(t.text) == 0 && !t.isFocused {
		textToDraw = t.placeholder
		textColor = rl.Gray
	}

	textX := pos.X + t.padding
	textY := pos.Y + (height-t.fontSize)/2

	rl.DrawText(textToDraw, int32(textX), int32(textY), int32(t.fontSize), textColor)

	// Draw cursor if focused and visible
	if t.isFocused && t.showCursor {
		// Calculate cursor position
		textBeforeCursor := t.text[:t.cursorPos]
		cursorX := textX + float32(rl.MeasureText(textBeforeCursor, int32(t.fontSize)))
		cursorY := textY

		rl.DrawRectangle(
			int32(cursorX),
			int32(cursorY),
			2,
			int32(t.fontSize),
			t.textColor,
		)
	}

	// Render children
	t.Rect.Render()
}
