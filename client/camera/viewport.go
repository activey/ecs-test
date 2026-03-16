package camera

import "github.com/yohamta/donburi/features/math"

type Viewport struct {
	Min math.Vec2
	Max math.Vec2
}

func (v Viewport) Dy() float64 {
	return v.Max.Y - v.Min.Y
}

func (v Viewport) Dx() float64 {
	return v.Max.X - v.Min.X
}
