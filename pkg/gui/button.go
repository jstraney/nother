package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/control"
	"github.com/jstraney/nother/pkg/node2d"
)

// Button is a 2D GUI button node that responds to click events.
type Button struct {
	*node2d.Node2D
	width      float32
	height     float32
	text       string
	fontSize   int32
	bgColor    rl.Color
	textColor  rl.Color
	hoverColor rl.Color
	isHovered  bool
	onClick    func()
	subscriber *control.Subscriber
}

// New creates a new Button with the given ID and name.
func NewButton(id string, name string) *Button {
	return &Button{
		Node2D:     node2d.New(id, name),
		width:      100,
		height:     40,
		text:       "Button",
		fontSize:   20,
		bgColor:    rl.Blue,
		textColor:  rl.White,
		hoverColor: rl.DarkBlue,
		isHovered:  false,
		onClick:    nil,
		subscriber: nil,
	}
}

// SetSize sets the button dimensions.
func (b *Button) SetSize(width, height float32) {
	b.width = width
	b.height = height
}

// SetText sets the button text.
func (b *Button) SetText(text string) {
	b.text = text
}

// SetFontSize sets the font size for the button text.
func (b *Button) SetFontSize(size int32) {
	b.fontSize = size
}

// SetColors sets the button background and text colors.
func (b *Button) SetColors(bgColor, textColor rl.Color) {
	b.bgColor = bgColor
	b.textColor = textColor
}

// SetHoverColor sets the color when the button is hovered.
func (b *Button) SetHoverColor(color rl.Color) {
	b.hoverColor = color
}

// SetOnClick sets the callback function when the button is clicked.
func (b *Button) SetOnClick(callback func()) {
	b.onClick = callback
}

// BindInputDispatcher binds the button to an input dispatcher for click handling.
func (b *Button) BindInputDispatcher(dispatcher *control.Dispatcher) {
	b.subscriber = dispatcher.NewSubscriber(control.MousePressedEventID(rl.MouseLeftButton))
	b.subscriber.Subscribe(b.handleMouseClick)
}

// handleMouseClick checks if the click is within the button bounds and calls the callback.
func (b *Button) handleMouseClick(event control.Event) {
	if b.onClick == nil {
		return
	}

	mousePos := rl.GetMousePosition()
	pos := b.GlobalPosition()

	// Check if mouse is within button bounds
	if mousePos.X >= pos.X && mousePos.X <= pos.X+b.width &&
		mousePos.Y >= pos.Y && mousePos.Y <= pos.Y+b.height {
		b.onClick()
	}
}

// Render draws the button with its text.
func (b *Button) Render() {
	if !b.Enabled() {
		b.Node2D.Render()
		return
	}

	pos := b.GlobalPosition()
	mousePos := rl.GetMousePosition()

	// Check if hovered
	b.isHovered = mousePos.X >= pos.X && mousePos.X <= pos.X+b.width &&
		mousePos.Y >= pos.Y && mousePos.Y <= pos.Y+b.height

	// Draw button background
	bgColor := b.bgColor
	if b.isHovered {
		bgColor = b.hoverColor
	}
	rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(b.width), int32(b.height), bgColor)

	// Draw border
	rl.DrawRectangleLines(int32(pos.X), int32(pos.Y), int32(b.width), int32(b.height), rl.Black)

	// Draw text centered in button
	textWidth := rl.MeasureText(b.text, b.fontSize)
	textX := int32(pos.X + (b.width-float32(textWidth))/2)
	textY := int32(pos.Y + (b.height-float32(b.fontSize))/2)
	rl.DrawText(b.text, textX, textY, b.fontSize, b.textColor)

	// Render children
	b.Node2D.Render()
}
