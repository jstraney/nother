package tilemap

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Atlas manages a texture atlas and tile template definitions.
type Atlas struct {
	Texture    rl.Texture2D
	TileWidth  int32
	TileHeight int32
	Tiles      map[int]*Tile
}

// NewAtlas creates a new Atlas with the given texture and tile dimensions.
func NewAtlas(texture rl.Texture2D, tileWidth, tileHeight int32) *Atlas {
	return &Atlas{
		Texture:    texture,
		TileWidth:  tileWidth,
		TileHeight: tileHeight,
		Tiles:      make(map[int]*Tile),
	}
}

// AddTile adds a tile template to the atlas.
func (a *Atlas) AddTile(tile *Tile) {
	a.Tiles[tile.ID] = tile
}

// GetTile retrieves a tile template by ID.
func (a *Atlas) GetTile(id int) *Tile {
	return a.Tiles[id]
}

// CreateTileFromGrid creates a tile template from grid coordinates in the atlas texture.
// gridX and gridY are tile coordinates, not pixel coordinates.
func (a *Atlas) CreateTileFromGrid(id int, gridX, gridY int32) *Tile {
	sourceRect := rl.NewRectangle(
		float32(gridX*a.TileWidth),
		float32(gridY*a.TileHeight),
		float32(a.TileWidth),
		float32(a.TileHeight),
	)
	return NewTile(id, sourceRect)
}

// Unload frees the atlas texture from memory.
func (a *Atlas) Unload() {
	rl.UnloadTexture(a.Texture)
}
