package structs_methods_interfaces

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	a, b, c float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
