package world

import (
	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	math2 "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"image/color"
	"math"
)

type NodeOrientation int

const (
	NodeVertical NodeOrientation = iota
	NodeHorizontal
)

type NavigationNode struct {
	X, Y        int               // Position of the node in tile coordinates
	IsWalkable  bool              // Is this tile walkable?
	Neighbors   []*NavigationNode // Neighboring nodes
	Orientation NodeOrientation
}

func (n *NavigationNode) Vec2() math2.Vec2 {
	return math2.NewVec2(float64(n.X), float64(n.Y))
}

func (n *NavigationNode) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, neighbor := range n.Neighbors {
		if neighbor.IsWalkable {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (n *NavigationNode) PathNeighborCost(to astar.Pather) float64 {
	return 1.0 // Simple movement cost, you can adjust this based on terrain
}

func (n *NavigationNode) PathEstimatedCost(to astar.Pather) float64 {
	neighbor := to.(*NavigationNode)
	return math.Abs(float64(n.X-neighbor.X)) + math.Abs(float64(n.Y-neighbor.Y)) // Manhattan distance
}

func (n *NavigationNode) DrawDebug(screen *ebiten.Image, cameraTransform *transform.TransformData, tileWidth, tileHeight float64) {
	if !n.IsWalkable {
		return
	}

	// Calculate the center position of the node in the world space
	nodeCenterX := float64(n.X)*tileWidth + tileWidth/2
	nodeCenterY := float64(n.Y)*tileHeight + tileHeight/2

	// Apply the camera transformation to world coordinates
	screenX := (nodeCenterX - cameraTransform.LocalPosition.X) * cameraTransform.LocalScale.X
	screenY := (nodeCenterY - cameraTransform.LocalPosition.Y) * cameraTransform.LocalScale.Y

	// Draw a small circle at the node center
	radius := 5.0 // * cameraTransform.LocalScale.X // Scale radius with the camera zoom
	vector.DrawFilledCircle(screen, float32(screenX), float32(screenY), float32(radius), color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
}
