package sprite

// Animation represents a named sequence of frame indices.
type Animation struct {
	Name          string
	FrameIndices  []int   // Indices into SpriteAtlas.Frames
	FPS           float32 // Default frames per second
	Loop          bool    // Whether the animation loops
	PingPong      bool    // If true, animation plays forward then backward
}

// NewAnimation creates a new animation with the given name and frame indices.
func NewAnimation(name string, frameIndices []int, fps float32) *Animation {
	return &Animation{
		Name:         name,
		FrameIndices: frameIndices,
		FPS:          fps,
		Loop:         true,
		PingPong:     false,
	}
}

// SetLoop sets whether the animation loops.
func (a *Animation) SetLoop(loop bool) {
	a.Loop = loop
}

// SetPingPong sets whether the animation ping-pongs (plays forward then backward).
func (a *Animation) SetPingPong(pingPong bool) {
	a.PingPong = pingPong
}

// FrameCount returns the number of frames in the animation.
func (a *Animation) FrameCount() int {
	return len(a.FrameIndices)
}

// GetFrameIndex returns the frame index at the given position in the sequence.
func (a *Animation) GetFrameIndex(position int) int {
	if position >= 0 && position < len(a.FrameIndices) {
		return a.FrameIndices[position]
	}
	return -1
}
