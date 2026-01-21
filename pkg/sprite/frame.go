package sprite

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Frame represents a single frame in an animation.
type Frame struct {
	TextureID  string       // ID of the texture in the atlas
	SourceRect rl.Rectangle // Region of the texture for this frame
	Duration   float32      // Duration in seconds (overrides animation FPS if > 0)
	Pivot      rl.Vector2   // Pivot point for rotation (default is center)
}

// NewFrame creates a new Frame with the given texture ID and source rectangle.
func NewFrame(textureID string, sourceRect rl.Rectangle) *Frame {
	return &Frame{
		TextureID:  textureID,
		SourceRect: sourceRect,
		Duration:   0, // 0 means use animation's default FPS
		Pivot:      rl.NewVector2(sourceRect.Width/2, sourceRect.Height/2),
	}
}

// SetDuration sets a custom duration for this frame (overrides animation FPS).
func (f *Frame) SetDuration(duration float32) {
	f.Duration = duration
}

// SetPivot sets the pivot point for rotation.
func (f *Frame) SetPivot(pivot rl.Vector2) {
	f.Pivot = pivot
}
