package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
	"github.com/jstraney/nother/pkg/signal"
)

// WindowResizeEvent is the event data for window resize signals.
type WindowResizeEvent struct {
	Width  int
	Height int
}

// GuiContext is the root GUI container that manages layout and theming.
type GuiContext struct {
	*node2d.Node2D
	theme        *Theme
	windowWidth  int
	windowHeight int
	enabled      bool
	subscriber   *signal.Subscriber
}

// NewGuiContext creates a new GUI context with the given theme.
func NewGuiContext(id string, name string, theme *Theme) *GuiContext {
	if theme == nil {
		theme = NewDefaultTheme()
	}

	ctx := &GuiContext{
		Node2D:       node2d.New(id, name),
		theme:        theme,
		windowWidth:  800,
		windowHeight: 600,
		enabled:      true,
		subscriber:   nil,
	}

	// Subscribe to window resize events
	ctx.subscriber = signal.GlobalBus.NewSubscriber("window.resize")
	ctx.subscriber.Subscribe(ctx.handleWindowResize)

	return ctx
}

// Theme returns the current theme.
func (c *GuiContext) Theme() *Theme {
	return c.theme
}

// SetTheme updates the theme.
func (c *GuiContext) SetTheme(theme *Theme) {
	c.theme = theme
}

// WindowWidth returns the current window width.
func (c *GuiContext) WindowWidth() int {
	return c.windowWidth
}

// WindowHeight returns the current window height.
func (c *GuiContext) WindowHeight() int {
	return c.windowHeight
}

// SetWindowSize manually sets the window dimensions (useful for initialization).
func (c *GuiContext) SetWindowSize(width, height int) {
	c.windowWidth = width
	c.windowHeight = height
	c.performLayout()
}

// IsEnabled returns whether the GUI is enabled.
func (c *GuiContext) IsEnabled() bool {
	return c.enabled
}

// SetEnabled enables or disables the entire GUI tree.
func (c *GuiContext) SetEnabled(enabled bool) {
	c.enabled = enabled
	if enabled {
		c.Enable()
	} else {
		c.Disable()
	}
}

// handleWindowResize processes window resize signals.
func (c *GuiContext) handleWindowResize(event signal.Event) {
	if resizeEvent, ok := event.Data.(WindowResizeEvent); ok {
		c.windowWidth = resizeEvent.Width
		c.windowHeight = resizeEvent.Height
		c.performLayout()
	}
}

// performLayout recalculates layout for all Layoutable children.
func (c *GuiContext) performLayout() {
	children := c.Children()
	if len(children) == 0 {
		return
	}

	// Collect layout constraints from all Layoutable children
	layoutables := make([]Layoutable, 0)
	for _, child := range children {
		if layoutable, ok := child.(Layoutable); ok {
			layoutables = append(layoutables, layoutable)
		}
	}

	if len(layoutables) == 0 {
		return
	}

	// Compute layout for each Layoutable component
	for _, layoutable := range layoutables {
		layout := c.computeLayout(layoutable)
		layoutable.ApplyLayout(layout.X, layout.Y, layout.Width, layout.Height)
	}

	// Emit layout complete signal
	signal.GlobalBus.Emit("gui.layout.complete", nil)
}

// computeLayout calculates the final position and size for a Layoutable component.
func (c *GuiContext) computeLayout(layoutable Layoutable) ComputedLayout {
	anchor := layoutable.Anchor()
	offset := layoutable.AnchorOffset()
	prefWidth, prefHeight := layoutable.PreferredSize()
	minWidth, minHeight := layoutable.MinSize()
	maxWidth, maxHeight := layoutable.MaxSize()

	// Clamp preferred size to min/max constraints
	width := prefWidth
	height := prefHeight

	if width < minWidth {
		width = minWidth
	}
	if minWidth > 0 && width < minWidth {
		width = minWidth
	}
	if maxWidth > 0 && width > maxWidth {
		width = maxWidth
	}

	if height < minHeight {
		height = minHeight
	}
	if minHeight > 0 && height < minHeight {
		height = minHeight
	}
	if maxHeight > 0 && height > maxHeight {
		height = maxHeight
	}

	// Compute base anchor position
	var anchorX, anchorY float32

	switch anchor {
	case AnchorTopLeft:
		anchorX, anchorY = 0, 0
	case AnchorTopCenter:
		anchorX, anchorY = float32(c.windowWidth)/2, 0
	case AnchorTopRight:
		anchorX, anchorY = float32(c.windowWidth), 0
	case AnchorCenterLeft:
		anchorX, anchorY = 0, float32(c.windowHeight)/2
	case AnchorCenter:
		anchorX, anchorY = float32(c.windowWidth)/2, float32(c.windowHeight)/2
	case AnchorCenterRight:
		anchorX, anchorY = float32(c.windowWidth), float32(c.windowHeight)/2
	case AnchorBottomLeft:
		anchorX, anchorY = 0, float32(c.windowHeight)
	case AnchorBottomCenter:
		anchorX, anchorY = float32(c.windowWidth)/2, float32(c.windowHeight)
	case AnchorBottomRight:
		anchorX, anchorY = float32(c.windowWidth), float32(c.windowHeight)
	}

	// Apply offset from anchor point
	x := anchorX + offset.X
	y := anchorY + offset.Y

	// Adjust position based on anchor alignment
	// (e.g., TopRight means the component's top-right corner is at the anchor point)
	switch anchor {
	case AnchorTopCenter, AnchorCenter, AnchorBottomCenter:
		x -= width / 2
	case AnchorTopRight, AnchorCenterRight, AnchorBottomRight:
		x -= width
	}

	switch anchor {
	case AnchorCenterLeft, AnchorCenter, AnchorCenterRight:
		y -= height / 2
	case AnchorBottomLeft, AnchorBottomCenter, AnchorBottomRight:
		y -= height
	}

	// Clamp to window bounds
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if x+width > float32(c.windowWidth) {
		x = float32(c.windowWidth) - width
		if x < 0 {
			x = 0
			width = float32(c.windowWidth)
		}
	}
	if y+height > float32(c.windowHeight) {
		y = float32(c.windowHeight) - height
		if y < 0 {
			y = 0
			height = float32(c.windowHeight)
		}
	}

	return ComputedLayout{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

// Update is called each frame to update the GUI tree.
func (c *GuiContext) Update(dt float32) {
	if !c.enabled {
		return
	}
	c.Node2D.Update(dt)
}

// Render draws the GUI tree.
func (c *GuiContext) Render() {
	if !c.enabled {
		return
	}
	c.Node2D.Render()
}

// Destroy cleans up the GuiContext.
func (c *GuiContext) Destroy() {
	if c.subscriber != nil {
		signal.GlobalBus.Unsubscribe("window.resize", c.subscriber)
	}
	c.Node2D.Destroy()
}
