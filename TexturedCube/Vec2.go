package main

import (
	"github.com/chewxy/math32"
)

type vec2 struct {
	x float32
	y float32
}

func (v1 *vec2) Distance(v2 vec2) float32 {
	return math32.Sqrt(v2.x*v1.x + v2.y*v1.y)
}

func (v *vec2) Length() float32 {
	return math32.Sqrt(v.x*v.x + v.y*v.y)
}

func (v1 *vec2) Minus(v2 vec2) vec2 {
	return vec2{v1.x - v2.x, v1.y - v2.y}
}

func (v1 *vec2) Dot(v2 vec2) vec2 {
	a := v2.x - v1.x
	b := v1.y - v2.y
	s := math32.Sqrt(a*a + b*b)
	if s > 0 {
		return vec2{b / s, a / s}
	}
	return vec2{}
}
