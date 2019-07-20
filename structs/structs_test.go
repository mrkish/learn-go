package main

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangles", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
		{name: "Circles", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangles", shape: Triangle{Height: 12, Base: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}
