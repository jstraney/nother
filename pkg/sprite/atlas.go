package sprite

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/asset"
)

// SpriteAtlas manages multiple textures and provides frame definitions.
type SpriteAtlas struct {
	Textures   map[string]*asset.Texture
	Frames     []*Frame
	FrameWidth  int32 // For uniform grids
	FrameHeight int32 // For uniform grids
}

// NewSpriteAtlas creates a new sprite atlas.
func NewSpriteAtlas() *SpriteAtlas {
	return &SpriteAtlas{
		Textures:   make(map[string]*asset.Texture),
		Frames:     make([]*Frame, 0),
		FrameWidth:  0,
		FrameHeight: 0,
	}
}

// AddTexture adds a texture to the atlas with the given ID.
func (sa *SpriteAtlas) AddTexture(id string, texture *asset.Texture) {
	sa.Textures[id] = texture
}

// GetTexture retrieves a texture by ID.
func (sa *SpriteAtlas) GetTexture(id string) *asset.Texture {
	return sa.Textures[id]
}

// AddFrame adds a frame to the atlas and returns its index.
func (sa *SpriteAtlas) AddFrame(frame *Frame) int {
	index := len(sa.Frames)
	sa.Frames = append(sa.Frames, frame)
	return index
}

// GetFrame retrieves a frame by index.
func (sa *SpriteAtlas) GetFrame(index int) *Frame {
	if index >= 0 && index < len(sa.Frames) {
		return sa.Frames[index]
	}
	return nil
}

// CreateUniformGrid creates frames from a uniform grid on a texture.
// Returns the starting index of the created frames.
func (sa *SpriteAtlas) CreateUniformGrid(textureID string, frameWidth, frameHeight int32) int {
	texture := sa.GetTexture(textureID)
	if texture == nil {
		return -1
	}

	textureData := texture.Data()
	cols := textureData.Width / frameWidth
	rows := textureData.Height / frameHeight

	startIndex := len(sa.Frames)

	for row := int32(0); row < rows; row++ {
		for col := int32(0); col < cols; col++ {
			sourceRect := rl.NewRectangle(
				float32(col*frameWidth),
				float32(row*frameHeight),
				float32(frameWidth),
				float32(frameHeight),
			)
			frame := NewFrame(textureID, sourceRect)
			sa.AddFrame(frame)
		}
	}

	sa.FrameWidth = frameWidth
	sa.FrameHeight = frameHeight

	return startIndex
}

// CreateFramesFromRects creates frames from an array of custom rectangles.
// Returns the starting index of the created frames.
func (sa *SpriteAtlas) CreateFramesFromRects(textureID string, rects []rl.Rectangle) int {
	startIndex := len(sa.Frames)

	for _, rect := range rects {
		frame := NewFrame(textureID, rect)
		sa.AddFrame(frame)
	}

	return startIndex
}
