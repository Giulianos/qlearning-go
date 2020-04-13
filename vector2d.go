package qlearning

import (
	"fmt"
	"math"
)

type Vec2D struct {
	X int
	Y int
}

func (v Vec2D) RotateCCW() Vec2D {
	return Vec2D{X: v.Y, Y: -v.X}
}

func (v Vec2D) RotateCW() Vec2D {
	return Vec2D{X: -v.Y, Y: v.X}
}

func (v Vec2D) Add(u Vec2D) Vec2D {
	return Vec2D{X: v.X + u.X, Y: v.Y + u.Y}
}

func (v Vec2D) Subtract(u Vec2D) Vec2D {
	return Vec2D{X: v.X - u.X, Y: v.Y - u.Y}
}

func (v Vec2D) Norm() float64 {
	x, y := float64(v.X), float64(v.Y)
	return math.Sqrt(x*x + y*y)
}

func (v Vec2D) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}
