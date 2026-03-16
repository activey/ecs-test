package player

type Position struct {
	X, Y float64
}

func NewPosition(x float64, y float64) Position {
	return Position{X: x, Y: y}

}
