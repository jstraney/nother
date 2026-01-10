package asset

import (
	"fmt"
	"io/fs"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TextureLoader loads texture assets from the filesystem.
type TextureLoader struct {
	filesystem fs.FS
}

// NewTextureLoader creates a new TextureLoader.
func NewTextureLoader(filesystem fs.FS) *TextureLoader {
	return &TextureLoader{
		filesystem: filesystem,
	}
}

// Load loads a texture from the given path.
func (tl *TextureLoader) Load(path string) (*Texture, error) {
	// Validate path
	if path == "" {
		return nil, fmt.Errorf("texture path cannot be empty")
	}

	// Read the file data from the filesystem
	data, err := fs.ReadFile(tl.filesystem, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read texture file %q: %w", path, err)
	}

	// Determine file type from extension
	ext := filepath.Ext(path)
	if ext == "" {
		return nil, fmt.Errorf("texture file %q has no extension", path)
	}

	// Load texture using raylib
	// raylib-go's LoadImageFromMemory requires file type, data, and data length
	img := rl.LoadImageFromMemory(ext, data, int32(len(data)))

	texture := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)

	if texture.ID == 0 {
		return nil, fmt.Errorf("failed to create GPU texture for %q", path)
	}

	return NewTexture(path, texture), nil
}

// Unload unloads a texture.
func (tl *TextureLoader) Unload(texture *Texture) error {
	if texture == nil {
		return fmt.Errorf("cannot unload nil texture")
	}
	texture.Dispose()
	return nil
}

// SoundLoader loads sound assets from the filesystem.
type SoundLoader struct {
	filesystem fs.FS
}

// NewSoundLoader creates a new SoundLoader.
func NewSoundLoader(filesystem fs.FS) *SoundLoader {
	return &SoundLoader{
		filesystem: filesystem,
	}
}

// Load loads a sound from the given path.
func (sl *SoundLoader) Load(path string) (*Sound, error) {
	// Validate path
	if path == "" {
		return nil, fmt.Errorf("sound path cannot be empty")
	}

	// Convert fs.FS path to absolute path for raylib
	// raylib expects a file path, not bytes, so we need to work with the filesystem appropriately
	// For now, we'll attempt to load directly if filesystem supports it
	fullPath, err := resolveFilesystemPath(sl.filesystem, path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve sound path %q: %w", path, err)
	}

	sound := rl.LoadSound(fullPath)
	if sound.Stream.Buffer == nil {
		return nil, fmt.Errorf("failed to load sound %q", path)
	}

	return NewSound(path, sound), nil
}

// Unload unloads a sound.
func (sl *SoundLoader) Unload(sound *Sound) error {
	if sound == nil {
		return fmt.Errorf("cannot unload nil sound")
	}
	sound.Dispose()
	return nil
}

// ModelLoader loads 3D model assets from the filesystem.
type ModelLoader struct {
	filesystem fs.FS
}

// NewModelLoader creates a new ModelLoader.
func NewModelLoader(filesystem fs.FS) *ModelLoader {
	return &ModelLoader{
		filesystem: filesystem,
	}
}

// Load loads a model from the given path.
func (ml *ModelLoader) Load(path string) (*Model, error) {
	// Validate path
	if path == "" {
		return nil, fmt.Errorf("model path cannot be empty")
	}

	// Convert fs.FS path to absolute path for raylib
	fullPath, err := resolveFilesystemPath(ml.filesystem, path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve model path %q: %w", path, err)
	}

	model := rl.LoadModel(fullPath)
	if model.MeshCount == 0 {
		return nil, fmt.Errorf("failed to load model %q", path)
	}

	return NewModel(path, model), nil
}

// Unload unloads a model.
func (ml *ModelLoader) Unload(model *Model) error {
	if model == nil {
		return fmt.Errorf("cannot unload nil model")
	}
	model.Dispose()
	return nil
}

// resolveFilesystemPath attempts to resolve a logical path within an fs.FS
// to an actual filesystem path that raylib can use.
// This is a helper function to work around raylib's requirement for actual file paths.
func resolveFilesystemPath(fsys fs.FS, logicalPath string) (string, error) {
	// If the filesystem is an os.DirFS, we can extract the actual path
	// Otherwise, we return the logical path and hope it works
	// A proper implementation might cache extracted files to disk
	return logicalPath, nil
}
