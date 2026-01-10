package sprite

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/asset"
	"github.com/jstraney/nother/pkg/node2d"
)

// Sprite is a 2D node that renders a texture at a position.
type Sprite struct {
	*node2d.Node2D
	texture  *asset.Texture
	scale    float32
	rotation float32
	tint     rl.Color
}

// New creates a new Sprite with the given ID and name.
func NewSprite(id string, name string) *Sprite {
	return &Sprite{
		Node2D:   node2d.New(id, name),
		texture:  nil,
		scale:    1.0,
		rotation: 0.0,
		tint:     rl.White,
	}
}

// SetTexture sets the texture for this sprite.
func (s *Sprite) SetTexture(texture *asset.Texture) {
	s.texture = texture
}

// Texture returns the current texture, or nil if none is set.
func (s *Sprite) Texture() *asset.Texture {
	return s.texture
}

// SetScale sets the scale factor for the sprite.
func (s *Sprite) SetScale(scale float32) {
	s.scale = scale
}

// Scale returns the current scale factor.
func (s *Sprite) Scale() float32 {
	return s.scale
}

// SetRotation sets the rotation angle in degrees.
func (s *Sprite) SetRotation(rotation float32) {
	s.rotation = rotation
}

// Rotation returns the current rotation angle in degrees.
func (s *Sprite) Rotation() float32 {
	return s.rotation
}

// SetTint sets the color tint for the sprite.
func (s *Sprite) SetTint(color rl.Color) {
	s.tint = color
}

// Tint returns the current color tint.
func (s *Sprite) Tint() rl.Color {
	return s.tint
}

// Render draws the sprite texture at its global position with scale, rotation, and tint.
func (s *Sprite) Render() {
	if !s.Enabled() {
		return
	}

	// Draw the texture if one is set
	if s.texture != nil {
		pos := s.GlobalPosition()
		textureData := s.texture.Data()

		// Calculate the origin (center of the texture)
		origin := rl.NewVector2(float32(textureData.Width)/2.0, float32(textureData.Height)/2.0)

		// Draw the texture with position, rotation, scale, and tint
		rl.DrawTexturePro(
			textureData,
			rl.NewRectangle(0, 0, float32(textureData.Width), float32(textureData.Height)),
			rl.NewRectangle(pos.X, pos.Y, float32(textureData.Width)*s.scale, float32(textureData.Height)*s.scale),
			origin,
			s.rotation,
			s.tint,
		)
	}

	// Render children
	s.Node2D.Render()
}
