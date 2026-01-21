package sprite

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// AnimatedSprite is a sprite that renders animated frames.
// The AnimationManager should be added as a child node.
type AnimatedSprite struct {
	*node2d.Node2D
	manager  *AnimationManager
	scale    float32
	rotation float32
	tint     rl.Color
}

// NewAnimatedSprite creates a new animated sprite with the given animation manager.
func NewAnimatedSprite(id string, name string, manager *AnimationManager) *AnimatedSprite {
	sprite := &AnimatedSprite{
		Node2D:   node2d.New(id, name),
		manager:  manager,
		scale:    1.0,
		rotation: 0.0,
		tint:     rl.White,
	}

	// Add manager as child if not already added
	if manager.Parent() == nil {
		sprite.AddChild(manager)
	}

	return sprite
}

// Manager returns the animation manager.
func (as *AnimatedSprite) Manager() *AnimationManager {
	return as.manager
}

// SetScale sets the scale factor for the sprite.
func (as *AnimatedSprite) SetScale(scale float32) {
	as.scale = scale
}

// Scale returns the current scale factor.
func (as *AnimatedSprite) Scale() float32 {
	return as.scale
}

// SetRotation sets the rotation angle in degrees.
func (as *AnimatedSprite) SetRotation(rotation float32) {
	as.rotation = rotation
}

// Rotation returns the current rotation angle in degrees.
func (as *AnimatedSprite) Rotation() float32 {
	return as.rotation
}

// SetTint sets the color tint for the sprite.
func (as *AnimatedSprite) SetTint(color rl.Color) {
	as.tint = color
}

// Tint returns the current color tint.
func (as *AnimatedSprite) Tint() rl.Color {
	return as.tint
}

// Render draws the current animation frame at the sprite's global position.
func (as *AnimatedSprite) Render() {
	if !as.Enabled() {
		return
	}

	// Get current frame from animation manager
	frame := as.manager.CurrentFrame()
	if frame == nil {
		as.Node2D.Render()
		return
	}

	// Get texture from atlas
	texture := as.manager.Atlas.GetTexture(frame.TextureID)
	if texture == nil {
		as.Node2D.Render()
		return
	}

	pos := as.GlobalPosition()
	textureData := texture.Data()

	// Calculate destination rectangle with scale
	destRect := rl.NewRectangle(
		pos.X,
		pos.Y,
		frame.SourceRect.Width*as.scale,
		frame.SourceRect.Height*as.scale,
	)

	// Use frame's pivot point as origin
	origin := rl.NewVector2(
		frame.Pivot.X*as.scale,
		frame.Pivot.Y*as.scale,
	)

	// Draw the frame
	rl.DrawTexturePro(
		textureData,
		frame.SourceRect,
		destRect,
		origin,
		as.rotation,
		as.tint,
	)

	// Render children
	as.Node2D.Render()
}
