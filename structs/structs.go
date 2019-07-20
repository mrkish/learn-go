package main

import "math"

func main() {}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return 2 * (r.Height + r.Width)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2

}
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
