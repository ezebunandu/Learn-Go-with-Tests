package geometry_test

import (
	"geometry"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := geometry.Rectangle{10.0, 10.0}
	got := geometry.Perimeter(rectangle)
	want := 40.0
	if want != got {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   geometry.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: geometry.Rectangle{Length: 12, Breadth: 6}, hasArea: 72.0},
		{name: "Circle", shape: geometry.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: geometry.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if tc.hasArea != got {
				t.Errorf("%#v want %g, got %g", tc.shape, tc.hasArea, got)
			}
		})

	}
}
