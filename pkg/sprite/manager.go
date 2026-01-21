package sprite

import (
	"github.com/jstraney/nother/pkg/node2d"
	"github.com/jstraney/nother/pkg/signal"
)

// AnimationState represents the playback state of the animation.
type AnimationState int

const (
	StateStopped AnimationState = iota
	StatePlaying
	StatePaused
)

// AnimationManager manages a collection of animations and controls playback.
// Should be a child node of AnimatedSprite.
type AnimationManager struct {
	*node2d.Node2D
	Atlas            *SpriteAtlas
	Animations       map[string]*Animation
	currentAnimation *Animation
	defaultAnimation string
	currentFrame     int
	frameTime        float64
	state            AnimationState
	direction        int // 1 for forward, -1 for backward (for ping-pong)
	switchOnFinish   string // Name of animation to switch to when current finishes
}

// NewAnimationManager creates a new animation manager.
func NewAnimationManager(id string, name string, atlas *SpriteAtlas) *AnimationManager {
	return &AnimationManager{
		Node2D:           node2d.New(id, name),
		Atlas:            atlas,
		Animations:       make(map[string]*Animation),
		currentAnimation: nil,
		defaultAnimation: "",
		currentFrame:     0,
		frameTime:        0,
		state:            StateStopped,
		direction:        1,
		switchOnFinish:   "",
	}
}

// AddAnimation adds an animation to the manager.
func (am *AnimationManager) AddAnimation(animation *Animation) {
	am.Animations[animation.Name] = animation
}

// GetAnimation retrieves an animation by name.
func (am *AnimationManager) GetAnimation(name string) *Animation {
	return am.Animations[name]
}

// SetDefaultAnimation sets the default animation to play.
func (am *AnimationManager) SetDefaultAnimation(name string) {
	am.defaultAnimation = name
	if am.currentAnimation == nil && name != "" {
		am.Play(name)
	}
}

// Play starts playing the specified animation.
func (am *AnimationManager) Play(name string) {
	animation := am.Animations[name]
	if animation == nil {
		return
	}

	// Reset if switching to a different animation
	if am.currentAnimation != animation {
		am.currentAnimation = animation
		am.currentFrame = 0
		am.frameTime = 0
		am.direction = 1
	}

	am.state = StatePlaying
}

// Stop stops the current animation and resets to the first frame.
func (am *AnimationManager) Stop() {
	am.state = StateStopped
	am.currentFrame = 0
	am.frameTime = 0
	am.direction = 1
}

// Pause pauses the current animation at the current frame.
func (am *AnimationManager) Pause() {
	am.state = StatePaused
}

// Resume resumes a paused animation.
func (am *AnimationManager) Resume() {
	if am.state == StatePaused {
		am.state = StatePlaying
	}
}

// IsPlaying returns true if an animation is currently playing.
func (am *AnimationManager) IsPlaying() bool {
	return am.state == StatePlaying
}

// IsPaused returns true if the animation is paused.
func (am *AnimationManager) IsPaused() bool {
	return am.state == StatePaused
}

// IsStopped returns true if no animation is playing.
func (am *AnimationManager) IsStopped() bool {
	return am.state == StateStopped
}

// CurrentAnimation returns the currently playing animation name, or empty string.
func (am *AnimationManager) CurrentAnimation() string {
	if am.currentAnimation != nil {
		return am.currentAnimation.Name
	}
	return ""
}

// CurrentFrame returns the current frame from the atlas, or nil.
func (am *AnimationManager) CurrentFrame() *Frame {
	if am.currentAnimation == nil {
		return nil
	}

	frameIndex := am.currentAnimation.GetFrameIndex(am.currentFrame)
	return am.Atlas.GetFrame(frameIndex)
}

// SwitchOnFinish sets an animation to switch to when the current one finishes.
func (am *AnimationManager) SwitchOnFinish(animationName string) {
	am.switchOnFinish = animationName
}

// Update advances the animation based on delta time.
func (am *AnimationManager) Update(dt float64) {
	if am.state != StatePlaying || am.currentAnimation == nil {
		am.Node2D.Update(dt)
		return
	}

	// Get current frame from animation
	frameIndex := am.currentAnimation.GetFrameIndex(am.currentFrame)
	frame := am.Atlas.GetFrame(frameIndex)
	if frame == nil {
		am.Node2D.Update(dt)
		return
	}

	// Determine frame duration
	frameDuration := float64(1.0) / float64(am.currentAnimation.FPS)
	if frame.Duration > 0 {
		frameDuration = float64(frame.Duration)
	}

	// Advance time
	am.frameTime += dt

	// Check if it's time to advance to the next frame
	if am.frameTime >= frameDuration {
		am.frameTime -= frameDuration
		am.advanceFrame()
	}

	am.Node2D.Update(dt)
}

// advanceFrame moves to the next frame in the animation.
func (am *AnimationManager) advanceFrame() {
	if am.currentAnimation == nil {
		return
	}

	am.currentFrame += am.direction

	// Handle animation completion
	if am.direction > 0 && am.currentFrame >= am.currentAnimation.FrameCount() {
		am.onAnimationComplete()
	} else if am.direction < 0 && am.currentFrame < 0 {
		am.onAnimationComplete()
	}
}

// onAnimationComplete handles animation completion logic.
func (am *AnimationManager) onAnimationComplete() {
	if am.currentAnimation == nil {
		return
	}

	// Emit animation finished signal
	signal.GlobalBus.Emit("animation.finished", map[string]interface{}{
		"animation": am.currentAnimation.Name,
	})

	// Handle looping
	if am.currentAnimation.Loop {
		if am.currentAnimation.PingPong {
			// Reverse direction for ping-pong
			am.direction = -am.direction
			if am.direction > 0 {
				am.currentFrame = 0
			} else {
				am.currentFrame = am.currentAnimation.FrameCount() - 1
			}
		} else {
			// Loop back to start
			am.currentFrame = 0
		}
	} else {
		// Non-looping animation finished
		am.currentFrame = am.currentAnimation.FrameCount() - 1
		am.state = StateStopped

		// Switch to next animation if specified
		if am.switchOnFinish != "" {
			nextAnim := am.switchOnFinish
			am.switchOnFinish = ""
			am.Play(nextAnim)
		}
	}
}
