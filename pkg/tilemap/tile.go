package tilemap

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// CollisionType defines how a tile interacts with collision detection.
type CollisionType int

const (
	// CollisionNone allows full passage through the tile
	CollisionNone CollisionType = iota
	// CollisionFull blocks all passage through the tile
	CollisionFull
	// CollisionPartial allows partial passage (can be extended for edge-based collision)
	CollisionPartial
)

// Tile is a template/configuration that defines how to render a tile from an atlas.
// Multiple grid positions can reference the same Tile.
type Tile struct {
	ID         int
	SourceRect rl.Rectangle
	Collision  CollisionType
	FlipH      bool
	FlipV      bool
	Rotation   float32 // Rotation in degrees (0, 90, 180, 270)
	Properties map[string]interface{}
}

// NewTile creates a new Tile template with the given ID and source rectangle.
func NewTile(id int, sourceRect rl.Rectangle) *Tile {
	return &Tile{
		ID:         id,
		SourceRect: sourceRect,
		Collision:  CollisionNone,
		FlipH:      false,
		FlipV:      false,
		Rotation:   0,
		Properties: make(map[string]interface{}),
	}
}

// IsPassable returns true if the tile allows passage (not fully blocked).
func (t *Tile) IsPassable() bool {
	return t.Collision != CollisionFull
}

// IsBlocked returns true if the tile blocks passage.
func (t *Tile) IsBlocked() bool {
	return t.Collision == CollisionFull
}
