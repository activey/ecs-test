package world

import (
	"github.com/yohamta/donburi/features/transform"
)

type NavigationGrid struct {
	Nodes map[[2]int]*NavigationNode // Map of visible nodes, key is [x, y] coordinates
}

func NewNavigationGrid() *NavigationGrid {
	return &NavigationGrid{
		Nodes: make(map[[2]int]*NavigationNode),
	}
}

func (g *NavigationGrid) AddNode(x, y int, isWalkable bool) {
	node := &NavigationNode{X: x, Y: y, IsWalkable: isWalkable}
	g.Nodes[[2]int{x, y}] = node
}

func (g *NavigationGrid) BuildNeighbors() {
	directions := [][2]int{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0}, // up, down, left, right
	}
	for _, node := range g.Nodes {
		for _, dir := range directions {
			neighborPos := [2]int{node.X + dir[0], node.Y + dir[1]}
			if neighbor, exists := g.Nodes[neighborPos]; exists {
				node.Neighbors = append(node.Neighbors, neighbor)
			}
		}
	}
}

func (g *NavigationGrid) GetNodeFromPixel(x, y float64, tileWidth, tileHeight int, cameraTransform *transform.TransformData) *NavigationNode {
	worldX := (x / cameraTransform.LocalScale.X) + cameraTransform.LocalPosition.X
	worldY := (y / cameraTransform.LocalScale.Y) + cameraTransform.LocalPosition.Y

	tileX := int(worldX) / tileWidth
	tileY := int(worldY) / tileHeight

	return g.Nodes[[2]int{tileX, tileY}]
}
