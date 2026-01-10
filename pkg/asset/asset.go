package asset

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Loader is the interface that loaders must implement.
// Each loader is responsible for loading a specific type of asset from the filesystem.
type Loader interface {
	// Load loads an asset from the given path within the filesystem.
	Load(path string) error
}

// Texture is a 2D image asset wrapping raylib's Texture2D.
type Texture struct {
	path    string
	texture rl.Texture2D
}

// NewTexture creates a new Texture wrapper.
func NewTexture(path string, texture rl.Texture2D) *Texture {
	return &Texture{
		path:    path,
		texture: texture,
	}
}

// Path returns the filesystem path this texture was loaded from.
func (t *Texture) Path() string {
	return t.path
}

// Data returns the underlying raylib Texture2D.
func (t *Texture) Data() rl.Texture2D {
	return t.texture
}

// Dispose unloads the texture from GPU memory.
func (t *Texture) Dispose() {
	rl.UnloadTexture(t.texture)
}

// Sound is an audio asset wrapping raylib's Sound.
type Sound struct {
	path  string
	sound rl.Sound
}

// NewSound creates a new Sound wrapper.
func NewSound(path string, sound rl.Sound) *Sound {
	return &Sound{
		path:  path,
		sound: sound,
	}
}

// Path returns the filesystem path this sound was loaded from.
func (s *Sound) Path() string {
	return s.path
}

// Data returns the underlying raylib Sound.
func (s *Sound) Data() rl.Sound {
	return s.sound
}

// Dispose unloads the sound from memory.
func (s *Sound) Dispose() {
	rl.UnloadSound(s.sound)
}

// Model is a 3D model asset wrapping raylib's Model.
type Model struct {
	path  string
	model rl.Model
}

// NewModel creates a new Model wrapper.
func NewModel(path string, model rl.Model) *Model {
	return &Model{
		path:  path,
		model: model,
	}
}

// Path returns the filesystem path this model was loaded from.
func (m *Model) Path() string {
	return m.path
}

// Data returns the underlying raylib Model.
func (m *Model) Data() rl.Model {
	return m.model
}

// Dispose unloads the model from memory.
func (m *Model) Dispose() {
	rl.UnloadModel(m.model)
}
