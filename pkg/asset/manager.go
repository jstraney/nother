package asset

import (
	"fmt"
	"io/fs"
	"sync"
)

// AssetManager loads and caches assets, handling lifecycle management.
// It uses loaders to support multiple asset types (Texture, Sound, Model).
type AssetManager struct {
	filesystem    fs.FS
	textureLoader *TextureLoader
	soundLoader   *SoundLoader
	modelLoader   *ModelLoader

	// Caches for loaded assets
	textures map[string]*Texture
	sounds   map[string]*Sound
	models   map[string]*Model

	mu sync.RWMutex
}

// New creates a new AssetManager with the given filesystem.
// The filesystem is used to locate and load assets.
func New(filesystem fs.FS) *AssetManager {
	return &AssetManager{
		filesystem:    filesystem,
		textureLoader: NewTextureLoader(filesystem),
		soundLoader:   NewSoundLoader(filesystem),
		modelLoader:   NewModelLoader(filesystem),
		textures:      make(map[string]*Texture),
		sounds:        make(map[string]*Sound),
		models:        make(map[string]*Model),
	}
}

// LoadTexture loads a texture from the given path.
// If the texture is already loaded, it returns the cached asset.
func (am *AssetManager) LoadTexture(path string) (*Texture, error) {
	if path == "" {
		return nil, fmt.Errorf("texture path cannot be empty")
	}

	am.mu.RLock()
	if texture, ok := am.textures[path]; ok {
		am.mu.RUnlock()
		return texture, nil
	}
	am.mu.RUnlock()

	// Load the texture
	texture, err := am.textureLoader.Load(path)
	if err != nil {
		return nil, err
	}

	// Cache it
	am.mu.Lock()
	am.textures[path] = texture
	am.mu.Unlock()

	return texture, nil
}

// LoadSound loads a sound from the given path.
// If the sound is already loaded, it returns the cached asset.
func (am *AssetManager) LoadSound(path string) (*Sound, error) {
	if path == "" {
		return nil, fmt.Errorf("sound path cannot be empty")
	}

	am.mu.RLock()
	if sound, ok := am.sounds[path]; ok {
		am.mu.RUnlock()
		return sound, nil
	}
	am.mu.RUnlock()

	// Load the sound
	sound, err := am.soundLoader.Load(path)
	if err != nil {
		return nil, err
	}

	// Cache it
	am.mu.Lock()
	am.sounds[path] = sound
	am.mu.Unlock()

	return sound, nil
}

// LoadModel loads a model from the given path.
// If the model is already loaded, it returns the cached asset.
func (am *AssetManager) LoadModel(path string) (*Model, error) {
	if path == "" {
		return nil, fmt.Errorf("model path cannot be empty")
	}

	am.mu.RLock()
	if model, ok := am.models[path]; ok {
		am.mu.RUnlock()
		return model, nil
	}
	am.mu.RUnlock()

	// Load the model
	model, err := am.modelLoader.Load(path)
	if err != nil {
		return nil, err
	}

	// Cache it
	am.mu.Lock()
	am.models[path] = model
	am.mu.Unlock()

	return model, nil
}

// UnloadTexture unloads a texture from the cache.
func (am *AssetManager) UnloadTexture(path string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	texture, ok := am.textures[path]
	if !ok {
		return fmt.Errorf("texture %q not found in cache", path)
	}

	err := am.textureLoader.Unload(texture)
	if err != nil {
		return err
	}

	delete(am.textures, path)
	return nil
}

// UnloadSound unloads a sound from the cache.
func (am *AssetManager) UnloadSound(path string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	sound, ok := am.sounds[path]
	if !ok {
		return fmt.Errorf("sound %q not found in cache", path)
	}

	err := am.soundLoader.Unload(sound)
	if err != nil {
		return err
	}

	delete(am.sounds, path)
	return nil
}

// UnloadModel unloads a model from the cache.
func (am *AssetManager) UnloadModel(path string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	model, ok := am.models[path]
	if !ok {
		return fmt.Errorf("model %q not found in cache", path)
	}

	err := am.modelLoader.Unload(model)
	if err != nil {
		return err
	}

	delete(am.models, path)
	return nil
}

// UnloadAll unloads all cached assets.
func (am *AssetManager) UnloadAll() error {
	am.mu.Lock()
	defer am.mu.Unlock()

	var lastErr error

	// Unload all textures
	for path, texture := range am.textures {
		if err := am.textureLoader.Unload(texture); err != nil {
			lastErr = fmt.Errorf("failed to unload texture %q: %w", path, err)
		}
	}
	am.textures = make(map[string]*Texture)

	// Unload all sounds
	for path, sound := range am.sounds {
		if err := am.soundLoader.Unload(sound); err != nil {
			lastErr = fmt.Errorf("failed to unload sound %q: %w", path, err)
		}
	}
	am.sounds = make(map[string]*Sound)

	// Unload all models
	for path, model := range am.models {
		if err := am.modelLoader.Unload(model); err != nil {
			lastErr = fmt.Errorf("failed to unload model %q: %w", path, err)
		}
	}
	am.models = make(map[string]*Model)

	return lastErr
}

// TextureCount returns the number of cached textures.
func (am *AssetManager) TextureCount() int {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return len(am.textures)
}

// SoundCount returns the number of cached sounds.
func (am *AssetManager) SoundCount() int {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return len(am.sounds)
}

// ModelCount returns the number of cached models.
func (am *AssetManager) ModelCount() int {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return len(am.models)
}
