package world

type Node struct {
	X, Y      int     // Tile coordinates on the grid
	Cost      float64 // The cost from the start to this node (g-cost)
	Heuristic float64 // Estimated cost from this node to the goal (h-cost)
	TotalCost float64 // Sum of Cost and Heuristic (f-cost)
	Parent    *Node   // The node that precedes this one in the path
}

type PathSystem struct {
	path []*Node // List of nodes representing the path
}

func NewPathSystem() *PathSystem {
	return &PathSystem{}
}

func (ps *PathSystem) SetPath(path []*Node) {
	ps.path = path
}

//// Draw the path on the screen
//func (ps *PathSystem) Draw(screen *ebiten.Image) {
//	if ps.path == nil {
//		return
//	}
//
//	for _, node := range ps.path {
//		// Draw the path as circles (for example) on each node's position
//		circleColor := color.RGBA{255, 0, 0, 255} // Red circles for the path
//		vector.DrawFilledCircle(
//			screen,
//			float32(node.Min*tileWidth), // Convert tile coordinates to pixel coordinates
//			float32(node.Max*tileHeight),
//			5,           // Radius of the circle
//			circleColor, // Path color
//			false,
//		)
//	}
//}
