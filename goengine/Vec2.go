package goengine

import (
	"github.com/chewxy/math32"
)

type Vec2 struct {
	X float32
	Y float32
}

// V2 returns a new [Vec2] with the given x and y components.
func V2(x, y float32) Vec2 {
	return Vec2{x, y}
}

func (v1 *Vec2) Distance(v2 Vec2) float32 {
	return math32.Sqrt(v2.X*v1.X + v2.Y*v1.Y)
}

func (v *Vec2) Length() float32 {
	return math32.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v1 *Vec2) Minus(v2 Vec2) Vec2 {
	return Vec2{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 *Vec2) Dot(v2 Vec2) Vec2 {
	a := v2.X - v1.X
	b := v1.Y - v2.Y
	s := math32.Sqrt(a*a + b*b)
	if s > 0 {
		return Vec2{b / s, a / s}
	}
	return Vec2{}
}

// MulScalar multiplies each component of this vector by the scalar s and returns resulting vector.
func (v Vec2) MulScalar(s float32) Vec2 {
	return V2(v.X*s, v.Y*s)
}
