package tilemap

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// TileMap is a node that renders a grid of tiles from an atlas.
type TileMap struct {
	*node2d.Node2D
	Atlas      *Atlas
	Rows       int
	Cols       int
	Grid       [][]int // [row][col] = tile ID
	TileWidth  float32
	TileHeight float32
}

// NewTileMap creates a new TileMap with the given dimensions.
func NewTileMap(id string, name string, atlas *Atlas, rows, cols int) *TileMap {
	// Initialize grid with all zeros (empty tiles)
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	tileWidth := float32(atlas.TileWidth)
	tileHeight := float32(atlas.TileHeight)

	return &TileMap{
		Node2D:     node2d.New(id, name),
		Atlas:      atlas,
		Rows:       rows,
		Cols:       cols,
		Grid:       grid,
		TileWidth:  tileWidth,
		TileHeight: tileHeight,
	}
}

// SetTile sets the tile ID at the given grid position.
func (tm *TileMap) SetTile(row, col, tileID int) {
	if row >= 0 && row < tm.Rows && col >= 0 && col < tm.Cols {
		tm.Grid[row][col] = tileID
	}
}

// GetTile returns the tile ID at the given grid position.
func (tm *TileMap) GetTile(row, col int) int {
	if row >= 0 && row < tm.Rows && col >= 0 && col < tm.Cols {
		return tm.Grid[row][col]
	}
	return 0
}

// GetTileTemplate returns the Tile template at the given grid position.
func (tm *TileMap) GetTileTemplate(row, col int) *Tile {
	tileID := tm.GetTile(row, col)
	return tm.Atlas.GetTile(tileID)
}

// WorldToGrid converts world position to grid coordinates.
func (tm *TileMap) WorldToGrid(worldPos rl.Vector2) (row, col int) {
	mapPos := tm.GlobalPosition()
	localX := worldPos.X - mapPos.X
	localY := worldPos.Y - mapPos.Y

	col = int(localX / tm.TileWidth)
	row = int(localY / tm.TileHeight)
	return row, col
}

// GridToWorld converts grid coordinates to world position (top-left of tile).
func (tm *TileMap) GridToWorld(row, col int) rl.Vector2 {
	mapPos := tm.GlobalPosition()
	x := mapPos.X + float32(col)*tm.TileWidth
	y := mapPos.Y + float32(row)*tm.TileHeight
	return rl.NewVector2(x, y)
}

// IsColliding checks if a world position collides with a blocked tile.
func (tm *TileMap) IsColliding(worldPos rl.Vector2) bool {
	row, col := tm.WorldToGrid(worldPos)
	tile := tm.GetTileTemplate(row, col)
	return tile != nil && tile.IsBlocked()
}

// IsRectColliding checks if a rectangle collides with any blocked tiles.
func (tm *TileMap) IsRectColliding(rect rl.Rectangle) bool {
	// Check all four corners of the rectangle
	topLeft := rl.NewVector2(rect.X, rect.Y)
	topRight := rl.NewVector2(rect.X+rect.Width, rect.Y)
	bottomLeft := rl.NewVector2(rect.X, rect.Y+rect.Height)
	bottomRight := rl.NewVector2(rect.X+rect.Width, rect.Y+rect.Height)

	return tm.IsColliding(topLeft) ||
		tm.IsColliding(topRight) ||
		tm.IsColliding(bottomLeft) ||
		tm.IsColliding(bottomRight)
}

// GetTileAt returns the Tile template at the given world position.
func (tm *TileMap) GetTileAt(worldPos rl.Vector2) *Tile {
	row, col := tm.WorldToGrid(worldPos)
	return tm.GetTileTemplate(row, col)
}

// Render draws all tiles in the map.
func (tm *TileMap) Render() {
	if !tm.Enabled() {
		return
	}

	mapPos := tm.GlobalPosition()

	for row := 0; row < tm.Rows; row++ {
		for col := 0; col < tm.Cols; col++ {
			tileID := tm.Grid[row][col]
			if tileID == 0 {
				continue // Skip empty tiles
			}

			tile := tm.Atlas.GetTile(tileID)
			if tile == nil {
				continue
			}

			// Calculate destination position
			destX := mapPos.X + float32(col)*tm.TileWidth
			destY := mapPos.Y + float32(row)*tm.TileHeight

			destRect := rl.NewRectangle(destX, destY, tm.TileWidth, tm.TileHeight)

			// Apply flip and rotation
			sourceRect := tile.SourceRect

			// Handle horizontal flip
			if tile.FlipH {
				sourceRect.Width = -sourceRect.Width
			}

			// Handle vertical flip
			if tile.FlipV {
				sourceRect.Height = -sourceRect.Height
			}

			// Origin for rotation (center of tile)
			origin := rl.NewVector2(0, 0)
			if tile.Rotation != 0 {
				origin = rl.NewVector2(tm.TileWidth/2, tm.TileHeight/2)
			}

			// Draw the tile
			rl.DrawTexturePro(
				tm.Atlas.Texture,
				sourceRect,
				destRect,
				origin,
				tile.Rotation,
				rl.White,
			)
		}
	}

	// Render children
	tm.Node2D.Render()
}
