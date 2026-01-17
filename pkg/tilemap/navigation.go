package tilemap

import (
	"container/heap"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/node2d"
)

// NavNodeOffset defines where navigation nodes are positioned within tiles.
type NavNodeOffset int

const (
	NavOffsetTopLeft NavNodeOffset = iota
	NavOffsetCenter
	NavOffsetBottomRight
)

// TileNavigation provides A* pathfinding for a TileMap.
// Should be added as a child of TileMap.
type TileNavigation struct {
	*node2d.Node2D
	TileMap    *TileMap
	NodeOffset NavNodeOffset
}

// NewTileNavigation creates a navigation system for the given tilemap.
func NewTileNavigation(id string, name string, tileMap *TileMap) *TileNavigation {
	return &TileNavigation{
		Node2D:     node2d.New(id, name),
		TileMap:    tileMap,
		NodeOffset: NavOffsetCenter,
	}
}

// SetNodeOffset sets where nav nodes are positioned within tiles.
func (tn *TileNavigation) SetNodeOffset(offset NavNodeOffset) {
	tn.NodeOffset = offset
}

// GetNavPosition returns the navigation position for a tile at (row, col).
func (tn *TileNavigation) GetNavPosition(row, col int) rl.Vector2 {
	tilePos := tn.TileMap.GridToWorld(row, col)

	switch tn.NodeOffset {
	case NavOffsetTopLeft:
		return tilePos
	case NavOffsetCenter:
		return rl.NewVector2(
			tilePos.X+tn.TileMap.TileWidth/2,
			tilePos.Y+tn.TileMap.TileHeight/2,
		)
	case NavOffsetBottomRight:
		return rl.NewVector2(
			tilePos.X+tn.TileMap.TileWidth,
			tilePos.Y+tn.TileMap.TileHeight,
		)
	default:
		return tilePos
	}
}

// FindPath finds a path from start world position to goal world position.
// Returns a slice of world positions representing the path, or nil if no path exists.
func (tn *TileNavigation) FindPath(startWorld, goalWorld rl.Vector2) []rl.Vector2 {
	// Convert world positions to grid coordinates
	startRow, startCol := tn.TileMap.WorldToGrid(startWorld)
	goalRow, goalCol := tn.TileMap.WorldToGrid(goalWorld)

	// Check if start or goal are out of bounds or blocked
	if !tn.isPassable(startRow, startCol) || !tn.isPassable(goalRow, goalCol) {
		return nil
	}

	// Run A* pathfinding
	path := tn.astar(startRow, startCol, goalRow, goalCol)
	if path == nil {
		return nil
	}

	// Convert grid path to world positions
	worldPath := make([]rl.Vector2, len(path))
	for i, gridPos := range path {
		worldPath[i] = tn.GetNavPosition(gridPos.row, gridPos.col)
	}

	return worldPath
}

// isPassable checks if a grid position is passable.
func (tn *TileNavigation) isPassable(row, col int) bool {
	// Out of bounds is not passable
	if row < 0 || row >= tn.TileMap.Rows || col < 0 || col >= tn.TileMap.Cols {
		return false
	}

	tile := tn.TileMap.GetTileTemplate(row, col)
	if tile == nil {
		return true // Empty tiles are passable
	}

	return tile.IsPassable()
}

// gridPos represents a grid coordinate.
type gridPos struct {
	row, col int
}

// astarNode represents a node in the A* algorithm.
type astarNode struct {
	pos    gridPos
	g      float32 // Cost from start
	h      float32 // Heuristic cost to goal
	f      float32 // g + h
	parent *astarNode
	index  int     // Index in heap
}

// priorityQueue implements heap.Interface for A* open set.
type priorityQueue []*astarNode

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*astarNode)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

// heuristic calculates the Manhattan distance heuristic.
func (tn *TileNavigation) heuristic(row1, col1, row2, col2 int) float32 {
	dx := float32(abs(col1 - col2))
	dy := float32(abs(row1 - row2))
	return dx + dy
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// astar implements the A* pathfinding algorithm.
func (tn *TileNavigation) astar(startRow, startCol, goalRow, goalCol int) []gridPos {
	start := gridPos{startRow, startCol}
	goal := gridPos{goalRow, goalCol}

	// Open set (priority queue) and closed set
	openSet := &priorityQueue{}
	heap.Init(openSet)

	closedSet := make(map[gridPos]bool)
	allNodes := make(map[gridPos]*astarNode)

	// Add start node
	startNode := &astarNode{
		pos:    start,
		g:      0,
		h:      tn.heuristic(startRow, startCol, goalRow, goalCol),
		parent: nil,
	}
	startNode.f = startNode.g + startNode.h
	heap.Push(openSet, startNode)
	allNodes[start] = startNode

	// Directions: up, down, left, right
	directions := []struct{ dr, dc int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for openSet.Len() > 0 {
		// Get node with lowest f score
		current := heap.Pop(openSet).(*astarNode)

		// Check if we reached the goal
		if current.pos == goal {
			// Reconstruct path
			path := []gridPos{}
			for node := current; node != nil; node = node.parent {
				path = append([]gridPos{node.pos}, path...)
			}
			return path
		}

		closedSet[current.pos] = true

		// Check neighbors
		for _, dir := range directions {
			neighborRow := current.pos.row + dir.dr
			neighborCol := current.pos.col + dir.dc
			neighborPos := gridPos{neighborRow, neighborCol}

			// Skip if not passable or already evaluated
			if !tn.isPassable(neighborRow, neighborCol) || closedSet[neighborPos] {
				continue
			}

			// Calculate tentative g score
			tentativeG := current.g + 1.0

			neighborNode, exists := allNodes[neighborPos]
			if !exists {
				// Create new node
				neighborNode = &astarNode{
					pos:    neighborPos,
					g:      math.MaxFloat32,
					h:      tn.heuristic(neighborRow, neighborCol, goalRow, goalCol),
					parent: nil,
				}
				allNodes[neighborPos] = neighborNode
			}

			// If this path is better
			if tentativeG < neighborNode.g {
				neighborNode.parent = current
				neighborNode.g = tentativeG
				neighborNode.f = neighborNode.g + neighborNode.h

				// Add to open set if not already there
				if neighborNode.index == -1 || neighborNode.index >= openSet.Len() {
					heap.Push(openSet, neighborNode)
				} else {
					heap.Fix(openSet, neighborNode.index)
				}
			}
		}
	}

	// No path found
	return nil
}
