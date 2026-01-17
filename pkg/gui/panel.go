package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/area2d"
)

// Panel is a container GUI element with title bar, background, and border.
type Panel struct {
	*area2d.Rect
	title          string
	titleBar       *Pane
	titleLabel     *Label
	content        *Pane
	titleBarHeight float32
	borderWidth    float32
	borderColor    rl.Color
	titleBgColor   rl.Color
	contentBgColor rl.Color
}

// NewPanel creates a new Panel with the given ID and name.
func NewPanel(id string, name string) *Panel {
	panel := &Panel{
		Rect:           area2d.NewRect(id, name),
		title:          "Panel",
		titleBarHeight: 30,
		borderWidth:    2,
		borderColor:    rl.Black,
		titleBgColor:   rl.DarkGray,
		contentBgColor: rl.LightGray,
	}

	// Initialize title bar
	panel.titleBar = NewPane(id+"_titlebar", name+"_titlebar")
	panel.titleBar.SetVisible(true)
	panel.titleBar.SetFill(panel.titleBgColor)
	panel.titleBar.SetPadding(5)

	// Initialize title label
	panel.titleLabel = New(id+"_title", name+"_title")
	panel.titleLabel.SetText(panel.title)
	panel.titleLabel.SetColor(rl.White)
	panel.titleLabel.SetFontSize(16.0)

	// Initialize content pane
	panel.content = NewPane(id+"_content", name+"_content")
	panel.content.SetVisible(true)
	panel.content.SetFill(panel.contentBgColor)
	panel.content.SetPadding(10)

	// Set up hierarchy
	panel.AddChild(panel.titleBar)
	panel.titleBar.AddChild(panel.titleLabel)
	panel.AddChild(panel.content)

	// Default size
	panel.SetSize(200, 150)
	panel.SetVisible(true)

	return panel
}

// Title returns the panel title.
func (p *Panel) Title() string {
	return p.title
}

// SetTitle sets the panel title.
func (p *Panel) SetTitle(title string) {
	p.title = title
	p.titleLabel.SetText(title)
}

// TitleBar returns the title bar pane.
func (p *Panel) TitleBar() *Pane {
	return p.titleBar
}

// Content returns the content pane where child elements should be added.
func (p *Panel) Content() *Pane {
	return p.content
}

// TitleLabel returns the title label component.
func (p *Panel) TitleLabel() *Label {
	return p.titleLabel
}

// SetTitleBarHeight sets the height of the title bar.
func (p *Panel) SetTitleBarHeight(height float32) {
	p.titleBarHeight = height
	p.updateLayout()
}

// SetBorderWidth sets the border width.
func (p *Panel) SetBorderWidth(width float32) {
	p.borderWidth = width
}

// SetBorderColor sets the border color.
func (p *Panel) SetBorderColor(color rl.Color) {
	p.borderColor = color
}

// SetTitleBgColor sets the title bar background color.
func (p *Panel) SetTitleBgColor(color rl.Color) {
	p.titleBgColor = color
	p.titleBar.SetFill(color)
}

// SetContentBgColor sets the content area background color.
func (p *Panel) SetContentBgColor(color rl.Color) {
	p.contentBgColor = color
	p.content.SetFill(color)
}

// updateLayout recalculates the positions and sizes of title bar and content.
func (p *Panel) updateLayout() {
	width := p.Width()
	height := p.Height()

	// Update title bar
	p.titleBar.SetSize(width, p.titleBarHeight)
	p.titleBar.SetPosition(rl.NewVector2(0, 0))

	// Update content pane
	contentHeight := height - p.titleBarHeight
	p.content.SetSize(width, contentHeight)
	p.content.SetPosition(rl.NewVector2(0, p.titleBarHeight))
}

// SetSize overrides the base SetSize to update child layout.
func (p *Panel) SetSize(width, height float32) {
	p.Rect.SetSize(width, height)
	p.updateLayout()
}

// Render draws the panel with title bar, content, and border.
func (p *Panel) Render() {
	if !p.Enabled() {
		return
	}

	pos := p.GlobalPosition()
	width := p.Width()
	height := p.Height()

	// Draw outer border
	if p.borderWidth > 0 {
		rl.DrawRectangleLinesEx(
			rl.NewRectangle(pos.X, pos.Y, width, height),
			p.borderWidth,
			p.borderColor,
		)
	}

	// Render title bar and content (they will render their backgrounds)
	p.titleBar.Render()
	p.content.Render()
}
